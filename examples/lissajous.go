package examples

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

type Lissajous struct {
	palette []color.Color
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func NewLissajous(palette []color.Color, cycles int, res float64, size int, nframes int, delay int) (*Lissajous, error) {
	fnName := "NewLissajous"
	if err := validateGreaterThanEqualOne(cycles, fnName, "cycles"); err != nil {
		return nil, err
	}
	if err := validateGreaterThanEqualOne(nframes, fnName, "nframes"); err != nil {
		return nil, err
	}
	if err := validateGreaterThanEqualOne(delay, fnName, "delay"); err != nil {
		return nil, err
	}

	return &Lissajous{
		palette: palette,
		cycles:  cycles,
		res:     res,
		size:    size,
		nframes: nframes,
		delay:   delay}, nil

}

func validateGreaterThanEqualOne(n int, fn, name string) error {
	if n < 1 {
		return fmt.Errorf("%s: provided %s invalid (%d < 1)", fn, name, n)
	}

	return nil
}

func (l *Lissajous) Animate(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: l.nframes}
	phase := 0.0

	randomPaletteIndex := func() uint8 {
		return uint8(1 + rand.Int()%len(l.palette) - 1)
	}

	for i := 0; i < l.nframes; i++ {
		rect := image.Rect(0, 0, 2*l.size+1, 2*l.size+1)
		img := image.NewPaletted(rect, l.palette)
		for t := 0.0; t < float64(l.cycles)*2*math.Pi; t += l.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(l.size+int(x*float64(l.size)+0.5), l.size+int(y*float64(l.size)+0.5), randomPaletteIndex())
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, l.delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func ExampleLissajous() {
	palette := []color.Color{
		color.RGBA{0xF4, 0xDB, 0xD6, 0xFF}, // Rosewater
		color.RGBA{0xF0, 0xC6, 0xC6, 0xFF}, // Flamingo
		color.RGBA{0xF5, 0xBD, 0xE6, 0xFF}, // Pink
		color.RGBA{0xC6, 0xA0, 0xF6, 0xFF}, // Mauve
		color.RGBA{0xED, 0x87, 0x96, 0xFF}, // Red
		color.RGBA{0xEE, 0x99, 0xA0, 0xFF}, // Maroon
		color.RGBA{0xF5, 0xA9, 0x7F, 0xFF}, // Peach
		color.RGBA{0xEE, 0xD4, 0x9F, 0xFF}, // Yellow
		color.RGBA{0xA6, 0xDA, 0x95, 0xFF}, // Green
		color.RGBA{0x8B, 0xD5, 0xCA, 0xFF}, // Teal
		color.RGBA{0x91, 0xD7, 0xE3, 0xFF}, // Sky
		color.RGBA{0x7D, 0xC4, 0xE4, 0xFF}, // Sapphire
		color.RGBA{0x8A, 0xAD, 0xF4, 0xFF}, // Blue
		color.RGBA{0xB7, 0xBD, 0xF8, 0xFF}, // Lavender

		color.RGBA{0xCA, 0xD3, 0xF5, 0xFF}, // Text
		color.RGBA{0xB8, 0xC0, 0xE0, 0xFF}, // Subtext1
		color.RGBA{0xA5, 0xAD, 0xCB, 0xFF}, // Subtext0
		color.RGBA{0x93, 0x9A, 0xB7, 0xFF}, // Overlay2
		color.RGBA{0x80, 0x87, 0xA2, 0xFF}, // Overlay1
		color.RGBA{0x6E, 0x73, 0x8D, 0xFF}, // Overlay0
		color.RGBA{0x5B, 0x60, 0x78, 0xFF}, // Surface2
		color.RGBA{0x49, 0x4D, 0x64, 0xFF}, // Surface1
		color.RGBA{0x36, 0x3A, 0x4F, 0xFF}, // Surface0
		color.RGBA{0x24, 0x27, 0x3A, 0xFF}, // Base
		color.RGBA{0x1E, 0x20, 0x30, 0xFF}, // Mantle
		color.RGBA{0x18, 0x19, 0x26, 0xFF}, // Crust
	}

	l, _ := NewLissajous(palette, 5, 0.001, 100, 64, 8)
	l.Animate(os.Stdout)
}
