package local

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type ImageRepository struct{}

func NewRepository() *ImageRepository {
	return &ImageRepository{}
}

func (repo *ImageRepository) Create(texts []string) (path string, err error) {
	width := 600
	height := 315
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	path = "/tmp/" + RandomString(20) + ".png"
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}

	fontBytes, err := ioutil.ReadFile("../font.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	background := color.RGBA{7, 104, 159, 0xff}
	textcolor := color.RGBA{255, 201, 69, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, background)
		}
	}

	size := 24.0
	spacing := 1.5

	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(textcolor))

	pt := freetype.Pt(10, 10+int(c.PointToFixed(size)>>6))
	for _, s := range texts {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(size * spacing)
	}

	defer file.Close()
	png.Encode(file, img)
	return
}

func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
