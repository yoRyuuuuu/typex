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
	Stream       proto.Game_StreamClient
	EventChannel chan Event
}

func NewGameClient() *GameClient {
	return &GameClient{
		Stream:       nil,
		EventChannel: make(chan Event),
	}
}

func (c *GameClient) Start() {
	go c.watchEvent()
}

// サーバからのEvent通知を捌く
func (c *GameClient) watchEvent() {
	for {
		res, err := c.Stream.Recv()
		if err != nil {
			log.Printf("can not receive %v\n", err)
			return
		}
		switch res.GetEvent().(type) {
		case *proto.Response_Question: // お題通知
			c.EventChannel <- QuestionEvent{
				Text: res.GetQuestion().GetText(),
			}
		case *proto.Response_Start: // ゲーム開始通知
			c.EventChannel <- StartEvent{}
		case *proto.Response_Finish: // ゲーム終了通知
			c.EventChannel <- FinishEvent{
				Winner: res.GetFinish().GetWinner(),
			}
		case *proto.Response_Join: // 参加通知
			c.EventChannel <- JoinEvent{
				ID:     res.GetJoin().GetPlayer().Id,
				Name:   res.GetJoin().GetPlayer().Name,
				Health: int(res.GetJoin().GetPlayer().Health),
			}
		case *proto.Response_Damage: // ダメージ通知
			c.EventChannel <- DamageEvent{
				ID:     res.GetDamage().GetId(),
				Damage: int(res.GetDamage().GetHealth()),
			}
		}
	}
}

func (c *GameClient) handleAttackAction(text string, target string) {
	req := &proto.Request{
		Action: &proto.Request_Attack{
			Attack: &proto.Attack{Text: text, TargetId: target}},
	}
	if err := c.Stream.Send(req); err != nil {
		log.Printf("can not send %v\n", err)
		return
	}
}

// サーバ接続処理
func (g *Game) connect(grpcClient proto.GameClient, name string) (*proto.ConnectResponse, error) {
	req := proto.ConnectRequest{Name: name}
	resp, err := grpcClient.Connect(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	header := metadata.New(map[string]string{"authorization": resp.Id})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := grpcClient.Stream(ctx)
	if err != nil {
		return nil, err
	}
	g.Stream = stream
	return resp, nil
}
