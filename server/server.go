package server

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yoRyuuuuu/typex/typex-server/proto"
	"google.golang.org/grpc/metadata"
)

var clientTimeout = 15.0
var maxClients = 2

type client struct {
	streamServer proto.Game_StreamServer
	done         chan error
	lastMessage  time.Time
	id           uuid.UUID
	name         string
}

type GameServer struct {
	proto.UnimplementedGameServer
	clients map[uuid.UUID]*client
	mu      sync.RWMutex
	game    *Game
}

func (s *GameServer) removeClient(id uuid.UUID) {
	s.mu.Lock()
	delete(s.clients, id)
	s.game.Player--
	s.mu.Unlock()
}

func NewGameServer(game *Game) *GameServer {
	server := &GameServer{
		clients: make(map[uuid.UUID]*client),
		game:    game,
	}
	go server.watchEvent()
	go server.watchTimeout()
	return server
}

func (s *GameServer) getClientFromContext(ctx context.Context) (*client, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	tokenRaw := headers["authorization"]
	if len(tokenRaw) == 0 {
		return nil, errors.New("no token provided")
	}
	token, err := uuid.Parse(tokenRaw[0])
	if err != nil {
		return nil, errors.New("cannot parse token")
	}
	s.mu.RLock()
	clt, ok := s.clients[token]
	s.mu.RUnlock()
	if !ok {
		return nil, errors.New("token no recognized")
	}
	return clt, nil
}

func (s *GameServer) Stream(srv proto.Game_StreamServer) error {
	ctx := srv.Context()
	clt, err := s.getClientFromContext(ctx)
	if err != nil {
		return err
	}
	if clt.streamServer != nil {
		return errors.New("stream already active")
	}
	clt.streamServer = srv
	log.Println("start new server")

	// Wait for stream requests.
	go func() {
		for {
			req, err := srv.Recv()
			if err != nil {
				log.Printf("receive error %v", err)
				clt.done <- errors.New("failed to receive request")
				return
			}

			log.Printf("got message [ID %v, Request %+v]", clt.id, req)
			clt.lastMessage = time.Now()

			switch req.GetAction().(type) {
			case *proto.Request_Answer:
				s.handleAnswerRequest(req, clt)
			}
		}
	}()

	var doneError error
	select {
	case <-ctx.Done():
		doneError = ctx.Err()
	case doneError = <-clt.done:
	}

	log.Printf("stream done with error %v", doneError)
	log.Printf("%s - removing client", clt.id)
	s.removeClient(clt.id)

	return doneError
}

func (s *GameServer) Connect(ctx context.Context, req *proto.ConnectRequest) (*proto.ConnectResponse, error) {
	if len(s.clients) >= maxClients {
		return nil, errors.New("The server is full")
	}

	log.Printf("connect this server [Name: %v]", req.GetName())
	s.mu.Lock()
	token := uuid.New()
	s.clients[token] = &client{
		id:          token,
		done:        make(chan error),
		lastMessage: time.Now(),
		name:        req.GetName(),
	}

	s.game.Problems[token] = NewStubProblems()
	s.game.Name[token] = req.GetName()
	s.game.Player++
	log.Printf("[Player: %v]", s.game.Player)
	s.mu.Unlock()

	return &proto.ConnectResponse{
		Token: token.String(),
	}, nil
}

func (s *GameServer) handleAnswerRequest(req *proto.Request, clt *client) {
	s.game.ActionChannel <- AnswerAction{
		ID: clt.id,
	}
}

// backendから通知される変更の処理
func (s *GameServer) watchEvent() {
	for {
		event := <-s.game.EventChannel
		switch event.(type) {
		case StartEvent:
			s.handleStartEvent()
		case QuestionEvent:
			event := event.(QuestionEvent)
			s.handleQuestionEvent(event)
		case FinishEvent:
			event := event.(FinishEvent)
			s.handleFinishEvent(event)
		}
	}
}

func (s *GameServer) handleFinishEvent(event FinishEvent) {
	// ゲーム終了を通知する
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		res := &proto.Response{
			Action: &proto.Response_Finish{
				Finish: &proto.Finish{
					Winner: event.Winner,
				},
			},
		}

		if err := clt.streamServer.Send(res); err != nil {
			log.Printf("failed to send finish event %v: %v", clt.name, err)
		}
	}
}

func (s *GameServer) handleStartEvent() {
	// プレイヤー全員へ通知する
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		res := &proto.Response{
			Action: &proto.Response_Start{},
		}

		if err := clt.streamServer.Send(res); err != nil {
			log.Printf("failed to send start event %v: %v", clt.name, err)
		}
	}
}

func (s *GameServer) handleQuestionEvent(event QuestionEvent) {
	id := event.ID
	text := event.Text

	clt := s.clients[id]
	if clt == nil {
		return
	}

	res := &proto.Response{
		Action: &proto.Response_Question{
			Question: &proto.Question{
				Text: text,
			},
		},
	}

	if err := clt.streamServer.Send(res); err != nil {
		log.Printf("failed to send question event to %v: %v", clt.name, err)
		return
	}

	log.Printf("send %v to %v\n", text, clt.name)
}

func (s *GameServer) watchTimeout() {
	timeoutTicker := time.NewTicker(1 * time.Minute)
	for {
		for _, client := range s.clients {
			if time.Now().Sub(client.lastMessage).Minutes() > clientTimeout {
				client.done <- errors.New("you have been timed out")
				return
			}
		}
		<-timeoutTicker.C
	}
}
