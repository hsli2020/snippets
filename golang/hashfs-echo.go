package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"text/template"

	"github.com/benbjohnson/hashfs"
	"github.com/labstack/echo/v4"
	"github.com/polaris1119/embed"
)

func main() {
	e := echo.New()

	e.GET("/assets/*", func(ctx echo.Context) error {
		filename, err := url.PathUnescape(ctx.Param("*"))
		if err != nil {
			return err
		}

		isHashed := false
		if base, hash := hashfs.ParseName(filename); hash != "" {
			if embed.Fsys.HashName(base) == filename {
				filename = base
				isHashed = true
			}
		}

		f, err := embed.Fsys.Open(filename)
		if os.IsNotExist(err) {
			return echo.ErrNotFound
		} else if err != nil {
			return echo.ErrInternalServerError
		}
		defer f.Close()

		// Fetch file info. Disallow directories from being displayed.
		fi, err := f.Stat()
		if err != nil {
			return echo.ErrInternalServerError
		} else if fi.IsDir() {
			return echo.ErrForbidden
		}

		contentType := "text/plain"
		// Determine content type based on file extension.
		if ext := path.Ext(filename); ext != "" {
			contentType = mime.TypeByExtension(ext)
		}

		// Cache the file aggressively if the file contains a hash.
		if isHashed {
			ctx.Response().Header().Set("Cache-Control", `public, max-age=31536000`)
		}

		// Set content length.
		ctx.Response().Header().Set("Content-Length", strconv.FormatInt(fi.Size(), 10))

		// Flush header and write content.
		buf := new(bytes.Buffer)
		if ctx.Request().Method != "HEAD" {
			io.Copy(buf, f)
		}
		return ctx.Blob(http.StatusOK, contentType, buf.Bytes())
	})

	e.GET("/", func(ctx echo.Context) error {
		tpl, err := template.New("index.html").ParseFiles("template/index.html")
		if err != nil {
			return err
		}

		var buf = new(bytes.Buffer)
		err = tpl.Execute(buf, map[string]interface{}{
			"mainjs": embed.Fsys.HashName("static/main.js"),
		})
		if err != nil {
			return err
		}
		return ctx.HTML(http.StatusOK, buf.String())
	})

	e.Logger.Fatal(e.Start(":8080"))
}
/*
package embed

import (
 "embed"
 "github.com/benbjohnson/hashfs"
)

//go:embed static
var embedFS embed.FS

// 带 hash 功能的 fs.FS
var Fsys = hashfs.NewFS(embedFS)

├── cmd
│   ├── std
│   │   └── main.go
├── embed.go
├── go.mod
├── go.sum
├── static
│   └── main.js // 主要处理该文件的嵌入、hash
├── template
│   └── index.html

index.html：

<html>
  <head>
    <title>测试 Embed Hash</title>
    <script src="/assets/{{.mainjs}}"></script>
  </head>
  <body>
    <h1>测试 Embed Hash</h1><hr>
    <div>以下内容来自 JS：</div>
    <p id="content" style="color: red;"></p>
  </body>
</html>

main.js：

window.onload = function() {
    document.querySelector('#content').innerHTML = "我是 JS 内容";
}*/
