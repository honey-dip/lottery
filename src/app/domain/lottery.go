package domain

import (
	"math/rand"
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

func (lottery *Lottery) Draw(unixTime int) {
	lottery.Winners = []string{}
	rand.Seed(int64(unixTime))
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
