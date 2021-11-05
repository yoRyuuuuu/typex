package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yoRyuuuuu/typex/proto"
	"github.com/yoRyuuuuu/typex/server"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "8743", "The port to listen")
	flag.Parse()
	log.Printf("listening on port %s", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	game := server.NewGame()
	game.Start()

	s := grpc.NewServer()
	server := server.NewGameServer(game)
	proto.RegisterGameServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
