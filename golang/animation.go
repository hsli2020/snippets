package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"os"
)

func main() {

	const width, height = 300, 300
	var images []*image.Paletted
	var delays []int
	var disposals []byte

	var palette color.Palette = color.Palette{
		image.Transparent,
		image.Black,
		image.White,
		color.RGBA{0xEE, 0xEE, 0xEE, 255},
		color.RGBA{0xCC, 0xCC, 0xCC, 255},
		color.RGBA{0x99, 0x99, 0x99, 255},
		color.RGBA{0x66, 0x66, 0x66, 255},
		color.RGBA{0x33, 0x33, 0x33, 255},
	}
	dc := gg.NewContext(width, height)

	for i := 0; i < 16; i++ {
		dc.SetRGBA(1, 1, 1, 0)
		dc.Clear()

		c := math.Cos((float64(i) * math.Pi * 2) / 16)
		s := math.Sin((float64(i) * math.Pi * 2) / 16)
		x := 150 + 100*c
		y := 150 + 100*s
		dc.DrawCircle(float64(x), float64(y), 20)
		dc.SetRGBA(0, 0, 0, 1)
		dc.Fill()

		img := dc.Image()
		bounds := img.Bounds()

		dst := image.NewPaletted(bounds, palette)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
		images = append(images, dst)
		delays = append(delays, 10)
		disposals = append(disposals, gif.DisposalBackground)
	}

	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
		//Disposal: disposals,
	})
}
