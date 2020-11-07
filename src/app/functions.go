package p

import (
	"github.com/honey-dip/lottery/src/app/infrastructure/gcp"
	"github.com/honey-dip/lottery/src/app/interface/controller"
	"log"
	"net/http"
)

func Lottery(w http.ResponseWriter, r *http.Request) {
	repo := gcp.NewRepository()
	c := controller.NewController(repo)
	err := c.Get(w, r)
	if err != nil {
		log.Fatalf("error : %v", err)
	}
	return
}
