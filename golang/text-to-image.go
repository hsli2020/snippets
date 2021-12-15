package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func main() {
	text := "美国开始推动数字美元，此举将推动全球数字货币的进程。本期节目分析美联储对数字美元由观望到积极对态度转变，以及美国政府背后的影响。"
	err := TextToImage(text, "output.jpg")
	if err != nil {
		fmt.Println(err)
	}
}

const (
	defaultSize    float64 = 45
	defaultSpacing float64 = 1.5
)

func TextToImage(text, filename string) error {
	m := image.NewRGBA(image.Rect(0, 0, 1600, 900))
	draw.Draw(m, m.Bounds(), image.Black, image.ZP, draw.Src)

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("./font/SourceHanSans-Bold.ttf")
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(defaultSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(image.White)

	pt := freetype.Pt(250, 200+int(c.PointToFixed(defaultSize)>>6))
	if _, err := c.DrawString(text, pt); err != nil {
		return err
	}

	fd, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = jpeg.Encode(fd, m, nil)
	if err != nil {
		return err
	}

	return nil
}
