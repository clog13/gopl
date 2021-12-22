package ch1

import (
	"image"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func Server1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lisajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lisajous(w http.ResponseWriter) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), whiteIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
