package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yoRyuuuuu/typex/typex-server/client"
	"github.com/yoRyuuuuu/typex/typex-server/proto"
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

	grpcClient := proto.NewGameClient(conn)
	client := client.NewGameClient()
	err = client.Connect(grpcClient, *name)
	if err != nil {
		log.Fatalf("connect request failed %v", err)
	}

	client.Start()
}
