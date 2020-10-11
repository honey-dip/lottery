package service

import (
	"github.com/honey-dip/lottery/src/app/domain"
)

type LotteryService struct {
	ImageRepository ImageRepository
}

func (service *LotteryService) Draw(candidates []string, numberOfWiners int) (winners []string, image string, err error) {
	lottery := domain.NewLottery(candidates, numberOfWiners)
	lottery.Draw()
	winners = lottery.GetWinners()
	image, err = service.ImageRepository.Create(winners)
	return
}
