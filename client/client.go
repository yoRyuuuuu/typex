package client

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/yoRyuuuuu/typex/client/backend"
	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc/metadata"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	// 自身のPlayerInfoを追加
	playerInfo := &backend.PlayerInfo{
		ID:   c.game.ID,
		Name: playerName,
	}
	c.game.PlayerInfo[c.game.ID] = playerInfo

	// Playerが参加したことを通知
	for _, player := range res.GetPlayer() {
		playerInfo := &backend.PlayerInfo{
			ID:   player.Id,
			Name: player.Name,
		}
		c.game.PlayerInfo[player.Id] = playerInfo
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

			switch res.GetEvent().(type) {
			case *proto.Response_Question:
				c.game.EventChannel <- backend.QuestionEvent{
					Text: res.GetQuestion().GetText(),
				}
			case *proto.Response_Start:
				c.game.EventChannel <- backend.StartEvent{}
			case *proto.Response_Finish:
				c.game.EventChannel <- backend.FinishEvent{
					Winner: res.GetFinish().GetWinner(),
				}
			case *proto.Response_Join:
				c.game.EventChannel <- backend.JoinEvent{
					ID:   res.GetJoin().GetPlayer().Id,
					Name: res.GetJoin().GetPlayer().Name,
				}
			case *proto.Response_Damage:
				// 体力を更新する処理
				c.game.EventChannel <- backend.DamageEvent{
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
			case backend.Attack:
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
