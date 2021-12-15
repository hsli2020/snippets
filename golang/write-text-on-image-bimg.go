package main

import (
	"fmt"
	"github.com/ghiac/bimg"
	"os"
)

func main() {
	buffer, err := bimg.Read("image.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	image := bimg.NewImage(buffer)
	s, _ := image.Size()
	newImage, err := bimg.NewImage(buffer).AddText(bimg.AddText{
		Width:      s.Width,
		Height:     s.Height,
		DPI:        200,
		Top:        200,
		Left:       20,
		Text:       "هیچی نیستی",
		Font:       "sans bold 12",
		Background: bimg.Color{R: 255, G: 255, B: 255},
		Opacity:    1,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == 800 && size.Height == 600 {
		fmt.Println("The image size is valid")
	}

	bimg.Write("new.jpg", newImage)
}
