package main

import (
	"github.com/van0828/GoPL_PracticeAnswer/util"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{color.White, color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = iota // first color in palette
	blackIndex
	redIndex
	greenIndex
	blueIndex
)

type GifConfig struct {
	ColorIndex int `url:"color_index"`
	Cycles     int
	Res        float64
	Size       int
	Nframes    int
	Delay      int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var input GifConfig
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := util.UrlParameterUnMarshal(r.URL.Query(), &input); err != nil {
			log.Fatal(err)
		}
		if input.ColorIndex <= 0 || input.ColorIndex >= 4 {
			input.ColorIndex = blackIndex
		}
		if input.Cycles <= 0 {
			input.Cycles = 5
		}
		if input.Res <= 0.0 {
			input.Res = 0.001
		}
		if input.Size <= 0 {
			input.Size = 100
		}
		if input.Nframes <= 0 {
			input.Nframes = 64
		}
		if input.Delay <= 0 {
			input.Delay = 8
		}
		lissajous(w, input)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func lissajous(out io.Writer, conf GifConfig) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: conf.Nframes}
	phase := 0.0 // phase difference
	for i := 0; i < conf.Nframes; i++ {
		rect := image.Rect(0, 0, 2*conf.Size+1, 2*conf.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(conf.Cycles)*2*math.Pi; t += conf.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(conf.Size+int(x*float64(conf.Size)+0.5), conf.Size+int(y*float64(conf.Size)+0.5),
				uint8(conf.ColorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, conf.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
