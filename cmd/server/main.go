package main

import (
	"awesomeProject1/pkg/api"
	"awesomeProject1/pkg/roulette"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &roulette.GRPCServer{}
	api.RegisterRouletteServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
