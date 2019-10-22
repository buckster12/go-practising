package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x9f, 0xff, 0x33, 0xFF},
	color.RGBA{0xf9, 0xff, 0x33, 0xFF},
	color.RGBA{0xa0, 0x78, 0x44, 0xFF},
	color.RGBA{0x0a, 0xc9, 0xff, 0xFF},
	color.RGBA{0x9b, 0x9c, 0x76, 0xFF},
	color.RGBA{0xb9, 0x2f, 0xa3, 0xFF},
	color.RGBA{0xa3, 0xf2, 0xf9, 0xFF},
	color.RGBA{0x3a, 0x0a, 0xc2, 0xFF},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			rand.Seed(time.Now().Unix())
			randIndex := rand.Intn(len(palette))
			colorIndex := uint8(randIndex)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
