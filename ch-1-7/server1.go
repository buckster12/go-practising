package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
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

var mu sync.Mutex
var count int

func main() {
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	// mu.Lock()
	// count++
	// mu.Unlock()
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// func counter(w http.ResponseWriter, r *http.Request) {
// mu.Lock()
// fmt.Fprintf(w, "Count: %d\n", count)
// mu.Unlock()
// }

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
