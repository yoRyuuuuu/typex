package client

import (
	"context"
	"log"

	"github.com/yoRyuuuuu/typex/client/backend"
	. "github.com/yoRyuuuuu/typex/common"
	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc/metadata"
)

type GameClient struct {
	Stream proto.Game_StreamClient
	game   *backend.Game
}

func NewGameClient(game *backend.Game) *GameClient {
	return &GameClient{
		game: game,
	}
}

func (c *GameClient) Connect(grpcClient proto.GameClient, playerName string) error {
	req := proto.ConnectRequest{
		Name: playerName,
	}

	res, err := grpcClient.Connect(context.Background(), &req)
	if err != nil {
		return err
	}

	c.game.ID = res.GetToken()

	// Playerが参加したことを通知
	for _, player := range res.GetPlayer() {
		p := backend.Player{
			ID:   player.Id,
			Name: player.Name,
		}
		c.game.Players[player.Id] = p
	}

	header := metadata.New(map[string]string{"authorization": res.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := grpcClient.Stream(ctx)
	if err != nil {
		return err
	}

	c.Stream = stream

	return nil
}

func (c *GameClient) Start() {
	go func() {
		for {
			res, err := c.Stream.Recv()
			if err != nil {
				log.Printf("can not receive %v", err)
				return
			}

			switch res.GetAction().(type) {
			case *proto.Response_Question:
				c.game.EventChannel <- QuestionEvent{
					Text: res.GetQuestion().GetText(),
				}
			case *proto.Response_Start:
				c.game.EventChannel <- StartEvent{}
			case *proto.Response_Finish:
				c.game.EventChannel <- FinishEvent{
					Winner: res.GetFinish().GetWinner(),
				}
			case *proto.Response_Join:
				c.game.EventChannel <- JoinEvent{
					ID:   res.GetJoin().GetPlayer().Id,
					Name: res.GetJoin().GetPlayer().Name,
				}
			case *proto.Response_Attack:
				c.game.EventChannel <- AttackEvent{
					ID:     res.GetAttack().GetId(),
					Health: int(res.GetAttack().GetHealth()),
				}
			}
		}
	}()

	go func() {
		for {
			action := <-c.game.ActionChannel

			switch action.(type) {
			case backend.Answer:
				req := &proto.Request{
					Action: &proto.Request_Answer{
						Answer: &proto.Answer{},
					},
				}
				if err := c.Stream.Send(req); err != nil {
					log.Printf("%v", err)
				}
			}
		}
	}()
}
