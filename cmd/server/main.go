package main

import (
	"awesomeProject1/api/proto"
	"awesomeProject1/pkg/roulette"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &roulette.GRPCServer{}
	proto.RegisterRouletteServer(s, srv)

	/*
		db, err := repository.NewPostgresDB(repository.Config{

			Host:     "localhost",
			Port:     "5436",
			Username: "postgres",
			Password: "1111",
			DBName:   "postgres",
			SSLMode:  "disable",
		})
		if err != nil {
			log.Fatal("Failed to initialize db: ", err.Error())
		}

		repos := repository.NewRepository(db)
		services :=
	*/

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
