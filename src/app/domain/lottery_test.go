package domain

import (
	"testing"
	"time"
)

func TestDraw(t *testing.T) {
	candidates := []string{"a", "b", "c", "d", "e"}
	numberOfWiner := 2
	lottery := NewLottery(candidates, numberOfWiner)
	lottery.Draw(int(time.Now().Unix()))
	winners := lottery.GetWinners()
	if len(winners) != numberOfWiner {
		t.Fatalf("Expected winner length: 2, received: %#v", len(winners))
	}
	for _, winner := range winners {
		included := false
		for _, candidate := range candidates {
			if winner == candidate {
				included = true
			}
		}
		if !included {
			t.Fatal("winner doesn't exist in candidates")
		}
	}
}
