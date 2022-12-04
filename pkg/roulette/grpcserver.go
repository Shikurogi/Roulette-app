package roulette

import (
	"awesomeProject1/pkg/api"
	"golang.org/x/net/context"
	"math/rand"
	"strconv"
	"time"
)

type GRPCServer struct{}

func (s *GRPCServer) Spin(ctx context.Context, req *api.NewBet) (*api.BetResult, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 36
	ans := "Lose "
	n := strconv.Itoa(rand.Intn(max-min) + min)
	if req.GetBet() == n {
		ans = "Win "
	}
	ans += n
	return &api.BetResult{Result: ans}, nil
}
