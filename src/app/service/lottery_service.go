package service

import (
	"github.com/honey-dip/lottery/src/app/domain"
	"log"
)

type LotteryService struct {
	ImageRepository ImageRepository
}

func NewService(repo ImageRepository) *LotteryService {
	return &LotteryService{
		ImageRepository: repo,
	}
}

func (service *LotteryService) Draw(candidates []string, numberOfWiners int, unixTime int, fontpath string) (winners []string, image string, err error) {
	lottery := domain.NewLottery(candidates, numberOfWiners)
	log.Print("draw")
	lottery.Draw(unixTime)
	winners = lottery.GetWinners()
	log.Print("create image")
	image, err = service.ImageRepository.Create(winners, fontpath)
	return
}
