package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc/metadata"
)

var PlayerCount = 5
var clientTimeout = 15.0

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
	s.game.PlayerCount--
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
	headers, _ := metadata.FromIncomingContext(ctx)
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
			case *proto.Request_Attack:
				s.handleAttackRequest(req, clt)
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
	if len(s.clients) >= PlayerCount {
		return nil, errors.New("The server is full")
	}

	log.Printf("connect this server [Name: %v]", req.GetName())
	s.mu.Lock()
	id := uuid.New()

	// ???????????????????????????????????????????????????
	s.game.PlayerID = append(s.game.PlayerID, id)
	s.game.Problem[id] = NewDatasetIterator()
	playerInfo := &PlayerInfo{
		Health: InitialHealth,
		Name:   req.GetName(),
	}
	s.game.PlayerInfo[id] = playerInfo

	players := []*proto.Player{}
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		// ????????????????????????????????????????????????????????????
		player := &proto.Player{
			Id:     clt.id.String(),
			Name:   clt.name,
			Health: InitialHealth,
		}
		players = append(players, player)

		// ????????????????????????????????????????????????
		resp := &proto.Response{
			Event: &proto.Response_Join{
				Join: &proto.Join{
					Player: &proto.Player{
						Id:     id.String(),
						Name:   req.GetName(),
						Health: InitialHealth,
					},
				},
			},
		}

		if err := clt.streamServer.Send(resp); err != nil {
			log.Printf("failed to send finish event %v: %v", clt.name, err)
		}
	}

	player := &proto.Player{
		Id:     id.String(),
		Name:   req.GetName(),
		Health: InitialHealth,
	}
	players = append(players, player)

	s.clients[id] = &client{
		id:          id,
		done:        make(chan error),
		lastMessage: time.Now(),
		name:        req.GetName(),
	}

	s.game.PlayerCount++
	log.Printf("[Player: %v]", s.game.PlayerCount)
	s.mu.Unlock()

	fmt.Println(players)
	return &proto.ConnectResponse{
		Id:     id.String(),
		Player: players,
	}, nil
}

func (s *GameServer) handleAttackRequest(req *proto.Request, clt *client) {
	s.game.ActionChannel <- AttackAction{
		ID:     clt.id,
		Text:   req.GetAttack().GetText(),
		Target: req.GetAttack().GetTargetId(),
	}
}

// backend????????????????????????????????????
func (s *GameServer) watchEvent() {
	for {
		event := <-s.game.EventChannel
		switch event := event.(type) {
		case StartEvent:
			s.handleStartEvent()
		case QuestionEvent:
			s.handleQuestionEvent(event)
		case FinishEvent:
			s.handleFinishEvent(event)
		case DamageEvent:
			s.handleDamageEvent(event)
		}
	}
}

func (s *GameServer) handleDamageEvent(event DamageEvent) {
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		resp := &proto.Response{
			Event: &proto.Response_Damage{
				Damage: &proto.Damage{
					Id:     event.ID,
					Health: int64(event.Damage),
				},
			},
		}

		if err := clt.streamServer.Send(resp); err != nil {
			log.Printf("failed to send finish event %v: %v", clt.name, err)
		}
	}
}

func (s *GameServer) handleFinishEvent(event FinishEvent) {
	// ??????????????????????????????
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		res := &proto.Response{
			Event: &proto.Response_Finish{
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
	// ????????????????????????????????????
	for _, clt := range s.clients {
		if clt.streamServer == nil {
			continue
		}

		res := &proto.Response{
			Event: &proto.Response_Start{},
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
		Event: &proto.Response_Question{
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
			if time.Since(client.lastMessage).Minutes() > clientTimeout {
				client.done <- errors.New("you have been timed out")
				return
			}
		}
		<-timeoutTicker.C
	}
}
