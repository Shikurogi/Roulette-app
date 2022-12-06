package roulette

import (
	"awesomeProject1/api/proto"
	"golang.org/x/net/context"
	"math/rand"
	"strconv"
	"time"
)

type GRPCServer struct{}

func (s *GRPCServer) Spin(ctx context.Context, req *proto.NewBet) (*proto.BetResult, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 36
	ans := "Lose "
	n := rand.Intn(max-min) + min
	x := int32(n)
	if req.GetBet() == x {
		ans = "Win "
	}
	ans += strconv.Itoa(n)
	return &proto.BetResult{Result: ans}, nil
}
