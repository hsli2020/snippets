package main

import (
	"embed"

	"github.com/labstack/echo"
	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/builtin"
	"github.com/open2b/scriggo/native"
)

//go:embed template
var local embed.FS

type Product struct {
	Name string
	URL  string
}

func main() {
	products := []Product{}
	opts := &scriggo.BuildOptions{
		Globals: native.Declarations{
			"capitalize": builtin.Capitalize, // global function
			"products":   &products,
		},
	}

	template, err := scriggo.BuildTemplate(local, "template/index.html", opts)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// Uh... This is possibly race condition.
		products = []Product{
			{
				Name: "リンゴ",
				URL:  "https://ja.wikipedia.org/wiki/リンゴ",
			},
			{
				Name: "みかん",
				URL:  "https://ja.wikipedia.org/wiki/ウンシュウミカン",
			},
		}
		return template.Run(c.Response().Writer, nil, nil)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
