package main

import (
	"awesomeProject1/pkg/api"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Not enough arguments!")
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewRouletteClient(conn)
	res, err := c.Spin(context.Background(), &api.NewBet{Bet: flag.Arg(0)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
