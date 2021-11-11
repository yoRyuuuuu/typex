package client

import (
	"context"
	"log"

	"github.com/yoRyuuuuu/typex/client/backend"
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
		c.game.PlayerID = append(c.game.PlayerID, player.Id)
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
				c.game.EventReceiver <- backend.QuestionEvent{
					Text: res.GetQuestion().GetText(),
				}
			case *proto.Response_Start:
				c.game.EventReceiver <- backend.StartEvent{}
			case *proto.Response_Finish:
				c.game.EventReceiver <- backend.FinishEvent{
					Winner: res.GetFinish().GetWinner(),
				}
			case *proto.Response_Join:
				c.game.EventReceiver <- backend.JoinEvent{
					ID:   res.GetJoin().GetPlayer().Id,
					Name: res.GetJoin().GetPlayer().Name,
				}
			case *proto.Response_Attack:
				c.game.EventReceiver <- backend.DamageEvent{
					ID:     res.GetAttack().GetId(),
					Damage: int(res.GetAttack().GetHealth()),
				}
			}
		}
	}()

	go func() {
		for {
			event := <-c.game.EventSender

			switch event.(type) {
			case backend.DamageEvent:
				req := &proto.Request{
					Action: &proto.Request_Answer{Answer: &proto.Answer{}},
				}

				if err := c.Stream.Send(req); err != nil {
					log.Printf("%v", err)
				}
			}
		}
	}()
}
