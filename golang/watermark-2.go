package main

import (
    "bytes"
    "fmt"
    "strings"

    "io/ioutil"
    "os"
    "path"

    "math"
    "math/rand"
    "time"

    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"

    "github.com/gonum/plot/vg"
    "github.com/gonum/plot/vg/vgimg"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

// WaterMark for adding a watermark on the image
func WaterMark(img image.Image, markText string) (image.Image, error) {
    // image's length to canvas's length
    bounds := img.Bounds()
    w := vg.Length(bounds.Max.X) * vg.Inch / vgimg.DefaultDPI
    h := vg.Length(bounds.Max.Y) * vg.Inch / vgimg.DefaultDPI
    diagonal := vg.Length(math.Sqrt(float64(w*w + h*h)))

    // create a canvas, which width and height are diagonal
    c := vgimg.New(diagonal, diagonal)

    // draw image on the center of canvas
    rect := vg.Rectangle{}
    rect.Min.X = diagonal/2 - w/2
    rect.Min.Y = diagonal/2 - h/2
    rect.Max.X = diagonal/2 + w/2
    rect.Max.Y = diagonal/2 + h/2
    c.DrawImage(rect, img)

    // make a fontStyle, which width is vg.Inch * 0.7
    fontStyle, _ := vg.MakeFont("Courier", vg.Inch*0.7)

    // repeat the markText
    markTextWidth := fontStyle.Width(markText)
    unitText := markText
    for markTextWidth <= diagonal {
        markText += " " + unitText
        markTextWidth = fontStyle.Width(markText)
    }

    // set the color of markText
    c.SetColor(color.RGBA{0, 0, 0, 122})

    // set a random angle between 0 and p/2
    ? := math.Pi * rand.Float64() / 2
    c.Rotate(?)

    // set the lineHeight and add the markText
    lineHeight := fontStyle.Extents().Height * 1
    for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
        c.FillString(fontStyle, vg.Point{X: 0, Y: offset}, markText)
    }

    // canvas writeto jpeg
    // canvas.img is private
    // so use a buffer to transfer
    jc := vgimg.PngCanvas{Canvas: c}
    buff := new(bytes.Buffer)
    jc.WriteTo(buff)
    img, _, err := image.Decode(buff)
    if err != nil {
        return nil, err
    }

    // get the center point of the image
    ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

    // cutout the marked image
    size := bounds.Size()
    bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
    rv := image.NewRGBA(bounds)
    draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
    return rv, nil
}

// MarkingPicture for marking picture with text
func MarkingPicture(filepath, text string) (image.Image, error) {
    f, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        return nil, err
    }

    img, err = WaterMark(img, text)
    if err != nil {
        return nil, err
    }
    return img, nil
}

func writeTo(img image.Image, ext string) (rv *bytes.Buffer, err error) {
    ext = strings.ToLower(ext)
    rv = new(bytes.Buffer)
    switch ext {
    case ".jpg", ".jpeg":
        err = jpeg.Encode(rv, img, &jpeg.Options{Quality: 100})
    case ".png":
        err = png.Encode(rv, img)
    }
    return rv, err
}

func main() {
    if len(os.Args) < 3 {
        panic(fmt.Errorf("Please using like this:$ %s [file or directory] [text]", os.Args[0]))
    }

    target := os.Args[1]
    text := os.Args[2]

    if stat, err := os.Stat(target); err == nil && stat.IsDir() {
        files, _ := ioutil.ReadDir(target)
        for _, fn := range files {
            img, err := MarkingPicture(path.Join(target, fn.Name()), text)
            if err != nil {
                continue
            }

            ext := path.Ext(fn.Name())
            base := strings.Split(fn.Name(), ".")[0] + "_marked"
            f, err := os.Create(base + ext)
            if err != nil {
                panic(err)
            }

            buff, err := writeTo(img, ext)
            if err != nil {
                panic(err)
            }
            if _, err = buff.WriteTo(f); err != nil {
                panic(err)
            }
        }
    } else {
        img, err := MarkingPicture(target, text)
        if err != nil {
            panic(err)
        }

        ext := path.Ext(target)
        base := strings.Split(path.Base(target), ".")[0] + "_marked"
        f, err := os.Create(base + ext)
        if err != nil {
            panic(err)
        }

        buff, err := writeTo(img, ext)
        if err != nil {
            panic(err)
        }

        if _, err = buff.WriteTo(f); err != nil {
            panic(err)
        }
    }
}