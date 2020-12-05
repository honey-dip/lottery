package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	BucketName = "lottery-image"
)

type ImageRepository struct{}

func NewRepository() *ImageRepository {
	return &ImageRepository{}
}

func (repo *ImageRepository) Create(texts []string, fontpath string) (path string, err error) {
	width := 600
	height := 315
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	filename := RandomString(20) + ".png"
	path = "/tmp/" + filename
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}

	fontBytes, err := ioutil.ReadFile(fontpath)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return
	}

	background := color.RGBA{7, 4, 69, 0xff}
	textcolor := color.RGBA{245, 248, 187, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, background)
		}
	}

	size := 45.0
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

	png.Encode(file, img)
	file.Close()

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(BucketName).Object(filename).NewWriter(ctx)
	t, _ := os.Open(path)
	if _, err = io.Copy(wc, t); err != nil {
		return "", err
	}
	defer wc.Close()
	return "https://storage.googleapis.com/" + BucketName + "/" + filename, nil
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
