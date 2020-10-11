package domain

import (
	"math/rand"
	"time"
)

type Lottery struct {
	Candidates      []string
	NumberOfWinners int
	Winners         []string
}

func NewLottery(candidates []string, numberOfWinners int) *Lottery {
	return &Lottery{
		Candidates:      candidates,
		NumberOfWinners: numberOfWinners,
	}
}

func (lottery *Lottery) Draw() {
	lottery.Winners = []string{}
	rand.Seed(time.Now().UTC().UnixNano())
	targets := lottery.Candidates
	rand.Shuffle(len(targets), func(i, j int) {
		targets[i], targets[j] = targets[j], targets[i]
	})
	for i := 0; i < lottery.NumberOfWinners; i++ {
		lottery.Winners = append(lottery.Winners, targets[i])
	}
}

func (lottery Lottery) GetWinners() []string {
	return lottery.Winners
}
