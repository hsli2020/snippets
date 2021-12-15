package main

import (
    "flag"
    "fmt"
    "image"
    "image/color/palette"
    "image/draw"
    "image/gif"
    "io/ioutil"
    "os"

    _ "image/jpeg"
)

var path, output string
var delay int

func main() {
    flag.StringVar(&path, "p", "", "path to the folder containing images")
    flag.StringVar(&output, "o", "output.gif", "the name of the generated GIF")
    flag.IntVar(&delay, "d", 5, "delay between frames in seconds")
    flag.Parse()

    if path == "" {
        fmt.Println("A path is required")
        flag.PrintDefaults()
        return
    }

    if delay < 1 || delay > 10 {
        fmt.Println("delay must be between 1 and 10 inclusively")
        return
    }

    // fmt.Println("This will be a GIF generator!")
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println(err)
        return
    }

    anim := gif.GIF{}
    for _, info := range files {
        f, err := os.Open(path + "/" + info.Name())
        if err != nil {
            fmt.Printf("Could not open file %s. Error: %s\n", info.Name(), err)
            return
        }
        defer f.Close()
        img, _, _ := image.Decode(f)

        // Image has to be palleted before going into the GIF
        paletted := image.NewPaletted(img.Bounds(), palette.Plan9)
        draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

        anim.Image = append(anim.Image, paletted)
        anim.Delay = append(anim.Delay, delay*100)
    }

    f, _ := os.Create(output)
    defer f.Close()
    gif.EncodeAll(f, &anim)
}