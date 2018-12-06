package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	// - standard html  | iris.HTML(...)
	// - django         | iris.Django(...)
	// - pug(jade)      | iris.Pug(...)
	// - handlebars     | iris.Handlebars(...)
	// - amber          | iris.Amber(...)

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	// default template funcs are:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}

	// register a custom template func.
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	// register the view engine to the views, this will load the templates.
	app.RegisterView(tmpl)

	app.Get("/", hi)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8")) // defaults to that but you can change it.
}

func hi(ctx iris.Context) {
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "iris") // {{.Name}} will render: iris
	// ctx.ViewData("", myCcustomStruct{})
	ctx.View("hi.html")
}
