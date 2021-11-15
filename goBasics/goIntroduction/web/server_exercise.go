package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var cycles = 5

const (
	res     = 0.01
	size    = 100
	nframes = 64
	delay   = 8
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm: %v\n", err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			var err error
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintf(w, "atoi: %s\n", err)
			}
		}
	}
	lissajous(w)
}

func lissajous(out io.Writer) { // 不太明白
	var palette = []color.Color{color.White, color.Black, color.RGBA{0x22, 0x66, 0x22, 50}}

	const (
		whiteIndex = 0 // first color in palette
		backIndex  = 1 // second color in palette
		userIndex  = 2
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), userIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous %v\n", err)
	}
}
