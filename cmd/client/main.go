package main

import (
	"awesomeProject1/api/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
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

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	resX := int32(x)

	c := proto.NewRouletteClient(conn)
	res, err := c.Spin(context.Background(), &proto.NewBet{Bet: resX})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
