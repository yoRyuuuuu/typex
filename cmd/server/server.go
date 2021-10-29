package main

import (
	"log"
	"net"

	"github.com/yoRyuuuuu/typex/typex-server/proto"
	"github.com/yoRyuuuuu/typex/typex-server/server"
	"google.golang.org/grpc"
)

var port = ":8080"

func main() {
	log.Printf("listening on port %s", port)
	lis, err := net.Listen("tcp", port)
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
