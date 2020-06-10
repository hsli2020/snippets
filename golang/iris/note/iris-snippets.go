    app := iris.New()

    app.Get("/", handler) // Method: "GET"
    app.Post("/", handler) // Method: "POST"
    app.Put("/", handler) // Method: "PUT"
    app.Delete("/", handler) // Method: "DELETE"
    app.Options("/", handler) // Method: "OPTIONS"
    app.Trace("/", handler) // Method: "TRACE"
    app.Connect("/", handler) // Method: "CONNECT"
    app.Head("/", handler) // Method: "HEAD"
    app.Patch("/", handler) // Method: "PATCH"

    app.Any("/", handler) // register the route for all HTTP Methods

    app.Handle("GET", "/contact", func(ctx iris.Context) {
        ctx.HTML("<h1> Hello from /contact </h1>")
    })

	// This will serve the ./static/favicons/ion_32_32.ico to: localhost:8080/favicon.ico
	app.Favicon("./static/favicons/ion_32_32.ico")

	// app.Favicon("./static/favicons/ion_32_32.ico", "/favicon_48_48.ico")
	// This will serve the ./static/favicons/ion_32_32.ico to: localhost:8080/favicon_48_48.ico

	// Grouping Routes
    users := app.Party("/users", myAuthMiddlewareHandler)
	{
		// http://myhost.com/users/42/profile
		users.Get("/{id:int}/profile", userProfileHandler)

		// http://myhost.com/users/messages/1
		users.Get("/inbox/{id:int}", userMessageHandler)
	}

	v1 := app.Party("/api/v1")
	{
		v1.Get("/", h)
		v1.Put("/put", h)
		v1.Post("/post", h)
	}

	// you can use the "string" type which is valid for a single path parameter that can be anything.
	app.Get("/username/{name}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("name"))
	}) // type is missing = {name:string}

	// Let's register our first macro attached to int macro type.
	// "min" = the function
	// "minValue" = the argument of the function
	// func(string) bool = the macro's path parameter evaluator, this executes in serve time when
	// a user requests a path which contains the :int macro type with the min(...) macro parameter function.
	app.Macros().Int.RegisterFunc("min", func(minValue int) func(string) bool {
		// do anything before serve here [...]
		// at this case we don't need to do anything
		return func(paramValue string) bool {
			n, err := strconv.Atoi(paramValue)
			if err != nil {
				return false
			}
			return n >= minValue
		}
	})

	// http://localhost:8080/profile/id>=1
	// this will throw 404 even if it's found as route on : /profile/0, /profile/blabla, /profile/-1
	// macro parameter functions are optional of course.
	app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
		// second parameter is the error but it will always nil because we use macros,
		// the validaton already happened.
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("Hello id: %d", id)
	})

	// to change the error code per route's macro evaluator:
	app.Get("/profile/{id:int min(1)}/friends/{friendid:int min(1) else 504}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		friendid, _ := ctx.Params().GetInt("friendid")
		ctx.Writef("Hello id: %d looking for friend id: ", id, friendid)
	}) // this will throw e 504 error code instead of 404 if all route's macros not passed.

	// http://localhost:8080/game/a-zA-Z/level/0-9
	// remember, alphabetical is lowercase or uppercase letters only.
	app.Get("/game/{name:alphabetical}/level/{level:int}", func(ctx iris.Context) {
		ctx.Writef("name: %s | level: %s", ctx.Params().Get("name"), ctx.Params().Get("level"))
	})

	// let's use a trivial custom regexp that validates a single path parameter
	// which its value is only lowercase letters.

	// http://localhost:8080/lowercase/anylowercase
	app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
		ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %s",
			ctx.Params().Get("name"))
	})

	// http://localhost:8080/single_file/app.js
	app.Get("/single_file/{myfile:file}", func(ctx iris.Context) {
		ctx.Writef("file type validates if the parameter value has a form of a file name, got: %s",
			ctx.Params().Get("myfile"))
	})

	// http://localhost:8080/myfiles/any/directory/here/
	// this is the only macro type that accepts any number of path segments.
	app.Get("/myfiles/{directory:path}", func(ctx iris.Context) {
		ctx.Writef("path type accepts any number of path segments, path after /myfiles/ is: %s",
			ctx.Params().Get("directory"))
    })

    // when 404 then render the template $templatedir/errors/404.html
    app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){
        ctx.View("errors/404.html")
    })

    app.OnErrorCode(500, func(ctx iris.Context){
        // ...
    })

    func handler(ctx iris.Context){
        ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
    }

    func h(ctx iris.Context) {
        ctx.Application().Logger().Infof(ctx.Path())
        ctx.Writef("Hello from %s", ctx.Path())
    }

	app.Run(iris.Addr(":8080"))


//  The standard html,
//  its template parser is the golang.org/pkg/html/template/
//
//  Django,
//  its template parser is the github.com/flosch/pongo2
//
//  Pug(Jade),
//  its template parser is the github.com/Joker/jade
//
//  Handlebars,
//  its template parser is the github.com/aymerick/raymond
//
//  Amber,
//  its template parser is the github.com/eknkc/amber
