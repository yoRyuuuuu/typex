package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yoRyuuuuu/typex/typex-server/proto"
	"google.golang.org/grpc/metadata"
)

type GameClient struct {
	Stream proto.Game_StreamClient
}

func NewGameClient() *GameClient {
	return &GameClient{}
}

func (c *GameClient) Connect(grpcClient proto.GameClient, playerName string) error {
	req := proto.ConnectRequest{
		Name: playerName,
	}

	res, err := grpcClient.Connect(context.Background(), &req)
	if err != nil {
		return err
	}

	log.Printf("Token: %v\n", res.Token)
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
	// Handle stream messages.
	go func() {
		for {
			res, err := c.Stream.Recv()
			if err != nil {
				log.Printf("can not receive %v", err)
				return
			}

			switch res.GetAction().(type) {
			case *proto.Response_Question:
				c.handleQuestionResponse(res)
			case *proto.Response_Start:
				c.handleStartResponse(res)
			}
		}
	}()

	// Handle local game engine changes.
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	for {
		if sc.Scan() {
			req := &proto.Request{
				Action: &proto.Request_Answer{},
			}

			err := c.Stream.Send(req)
			if err != nil {
				log.Printf("failed to send message %v", err)
				return
			}

		} else {
			log.Printf("input scanner failure %v", sc.Err())
			return
		}
	}
}

func (c *GameClient) handleQuestionResponse(res *proto.Response) {
	fmt.Printf("%v\n", res.GetQuestion().GetText())
}

func (c *GameClient) handleStartResponse(res *proto.Response) {
	limit := 5 * time.Second
	count := 0
	output := []string{"4", "3", "2", "1", "start!!"}
	for begin := time.Now(); time.Since(begin) < limit; {
		fmt.Println(output[count])
		count += 1
		time.Sleep(1 * time.Second)
	}
}
