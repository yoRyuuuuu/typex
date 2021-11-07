package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yoRyuuuuu/typex/client"
	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("addr", "localhost", "The address to listen on.")
	port := flag.Int("port", 8743, "The port to listen on.")
	name := flag.String("name", "Hoge", "Player name")
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", *address, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can Not connect with server %v", err)
	}

	game := client.NewGame()
	clt := client.NewGameClient(game)
	grpcClient := proto.NewGameClient(conn)
	err = clt.Connect(grpcClient, *name)

	view := client.NewView(game)

	if err != nil {
		log.Fatalf("connect request failed %v", err)
	}

	clt.Start()
	view.Start()
}
