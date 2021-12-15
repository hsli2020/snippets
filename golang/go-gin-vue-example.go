// ./main.go

package main

import (
	"example/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	routers.Register(e)
	e.Run(":8080")
}


// ./routers/routers.go

package routers

import (
	"net/http"

	"example/web"

	"github.com/gin-gonic/gin"
)

// HandleIndex return HTML
func HandleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		html := web.MustAsset("index.html")
		c.Data(200, "text/html; charset=UTF-8", html)
	}
}

// Register routes
func Register(e *gin.Engine) {
	h := gin.WrapH(http.FileServer(web.AssetFile()))
	e.GET("/favicon.ico", h)
	e.GET("/js/*filepath", h)
	e.GET("/css/*filepath", h)
	e.GET("/img/*filepath", h)
	e.GET("/fonts/*filepath", h)
	e.NoRoute(HandleIndex())
}


// go generate ./web


// ./web/web.go => ./web/web_gen.go
package web

//go:generate npm run build
//go:generate go-bindata -fs -o web_gen.go -ignore *.map -pkg web -prefix dist/ ./dist/...
