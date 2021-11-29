package client

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc/metadata"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type GameClient struct {
	Stream proto.Game_StreamClient
	game   *Game
}

func NewGameClient(game *Game) *GameClient {
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
	for _, player := range res.GetPlayer() {
		playerInfo := &PlayerInfo{
			ID:     player.Id,
			Name:   player.Name,
			Health: int(player.Health),
		}
		c.game.PlayerInfo[player.Id] = playerInfo
		if player.Id != c.game.ID {
			c.game.PlayerID = append(c.game.PlayerID, player.Id)
		}
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

			switch res.GetEvent().(type) {
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
					ID:     res.GetJoin().GetPlayer().Id,
					Name:   res.GetJoin().GetPlayer().Name,
					Health: int(res.GetJoin().GetPlayer().Health),
				}
			case *proto.Response_Damage:
				// 体力を更新する処理
				c.game.EventChannel <- DamageEvent{
					ID:     res.GetDamage().GetId(),
					Damage: int(res.GetDamage().GetHealth()),
				}
			}
		}
	}()

	go func() {
		for {
			action := <-c.game.ActionSender

			switch action := action.(type) {
			case Attack:
				req := &proto.Request{
					Action: &proto.Request_Attack{
						Attack: &proto.Attack{
							Text: action.Text,
							Id:   action.ID,
						},
					},
				}

				if err := c.Stream.Send(req); err != nil {
					log.Printf("%v", err)
				}
			}
		}
	}()
}
