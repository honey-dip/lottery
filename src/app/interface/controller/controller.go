package controller

import (
	"github.com/honey-dip/lottery/src/app/service"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Controller struct {
	repo ImageRepository
}

func NewController(repo ImageRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

type Params struct {
	ImageUrl string
	Winners  string
}

func (controller *Controller) Get(w io.Writer, r *http.Request) error {
	q := r.URL.Query()
	candidatesQuery := "NoCandidate"
	numberQuery := "1"
	unixtimeQuery := strconv.FormatInt(time.Now().Unix(), 10)
	if q.Get("candidates") != "" {
		candidatesQuery = q.Get("candidates")
	}
	if q.Get("number") != "" {
		numberQuery = q.Get("number")
	}
	if q.Get("time") != "" {
		unixtimeQuery = q.Get("time")
	}
	unixtime, _ := strconv.Atoi(unixtimeQuery)
	candidates := strings.Split(candidatesQuery, ",")
	number, _ := strconv.Atoi(numberQuery)
	service := service.NewService(controller.repo)
	fontpath, _ := filepath.Abs("./serverless_function_source_code/font.ttf")
	winners, image, err := service.Draw(candidates, number, unixtime, fontpath)
	if err != nil {
		return err
	}
	params := Params{
		ImageUrl: image,
		Winners:  strings.Join(winners, ","),
	}
	tplpath, _ := filepath.Abs("./serverless_function_source_code/template.html")
	tpl := template.Must(template.ParseFiles(tplpath))
	err = tpl.Execute(w, params)
	if err != nil {
		return err
	}
	return nil
}
