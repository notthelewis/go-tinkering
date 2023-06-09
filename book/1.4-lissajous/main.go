// go run . >> test.gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func addRandomColoursToPallete(nFrames int) {
    for i := 0; i < nFrames; i++ {
        pallete = append(pallete, color.RGBA{
            uint8(rand.Intn(255)),
            uint8(rand.Intn(255)),
            uint8(rand.Intn(255)),
            uint8(rand.Intn(255)),
        })
    }
}


func lissajous(file *os.File) {
	const (
		cycles  = 5     // Number of complete X oscillator revolutions
		res     = 0.001 // Angular resolution
		size    = 100   // image canvas covers [-size..+suze]
		nframes = 64    // number of animation frames
		delay   = 8
	)

	freq := rand.Float64() * 3.0 // Relative frequency of Y oscillator
	anim := gif.GIF{LoopCount: nframes} 
	phase := 0.0

    addRandomColoursToPallete(nframes)

    // Apparently, sinusoid is another word for sine wave... the more you know 

	for i := 0; i < nframes; i++ {
        // set all pixels to zro
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(file, &anim)
}
