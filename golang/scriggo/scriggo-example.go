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
	products := []Product{
		{ Name: "リンゴ",	URL: "https://ja.wikipedia.org/wiki/blah-1" },
		{ Name: "みかん",	URL: "https://ja.wikipedia.org/wiki/blah-2" },
	}
	opts := &scriggo.BuildOptions{
		Globals: native.Declarations{
			"capitalize": builtin.Capitalize, // global function
			"products":   &products,
		},
	}

	template, err := scriggo.BuildTemplate(local, "template/index.html", opts)
	if err != nil {	panic(err) }

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return template.Run(c.Response().Writer, nil, nil)
	})
	e.Logger.Fatal(e.Start(":8989"))
}

// template/layout.html
<!DOCTYPE html>
<html>
<head><title>{{ Title() }}</title></head>
<body>
  {{ Body() }}
</body>
</head>


// template/banner.html
{% macro Banner %}そんなバナナ{% end %}


// template/index.html
{% extends "layout.html" %}
{% import "banners.html" %}
{% macro Title %}釣りタイトル{% end %}
{% macro Body %}
    <p>商品一覧</p>
    <ul>
      {% for product in products %}
      <li><a href="{{ product.URL }}">{{ product.Name }}</a></li>
      {% end %}
    </ul>
    {{ Banner() }}
{% end %}
