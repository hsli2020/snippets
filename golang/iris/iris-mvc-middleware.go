File: mvc/middleware/main.go

// Package main shows how you can add middleware to an mvc Application, simply
// by using its `Router` which is a sub router(an iris.Party) of the main iris app.
package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/mvc"
)

var cacheHandler = cache.Handler(10 * time.Second)

func main() {
	app := iris.New()
	mvc.Configure(app, configure)

	// http://localhost:8080
	// http://localhost:8080/other
	//
	// refresh every 10 seconds and you'll see different time output.
	app.Run(iris.Addr(":8080"))
}

func configure(m *mvc.Application) {
	m.Router.Use(cacheHandler)
	m.Handle(&exampleController{
		timeFormat: "Mon, Jan 02 2006 15:04:05",
	})
}

type exampleController struct {
	timeFormat string
}

func (c *exampleController) Get() string {
	now := time.Now().Format(c.timeFormat)
	return "last time executed without cache: " + now
}

func (c *exampleController) GetOther() string {
	now := time.Now().Format(c.timeFormat)
	return "/other: " + now
}

File: mvc/middleware/per-method/main.go

/*
If you want to use it as middleware for the entire controller
you can use its router which is just a sub router to add it as you normally do with standard API:

I'll show you 4 different methods for adding a middleware into an mvc application,
all of those 4 do exactly the same thing, select what you prefer,
I prefer the last code-snippet when I need the middleware to be registered somewhere
else as well, otherwise I am going with the first one:

```go
// 1
mvc.Configure(app.Party("/user"), func(m *mvc.Application) {
     m.Router.Use(cache.Handler(10*time.Second))
})
```

```go
// 2
// same:
userRouter := app.Party("/user")
userRouter.Use(cache.Handler(10*time.Second))
mvc.Configure(userRouter, ...)
```

```go
// 3
// same:
userRouter := app.Party("/user", cache.Handler(10*time.Second))
mvc.Configure(userRouter, ...)
```

```go
// 4
// same:
app.PartyFunc("/user", func(r iris.Party){
    r.Use(cache.Handler(10*time.Second))
    mvc.Configure(r, ...)
})
```

If you want to use a middleware for a single route,
for a single controller's method that is already registered by the engine
and not by custom `Handle` (which you can add
the middleware there on the last parameter) and it's not depend on the `Next Handler` to do its job
then you just call it on the method:

```go
var myMiddleware := myMiddleware.New(...) // this should return an iris/context.Handler

type UserController struct{}
func (c *UserController) GetSomething(ctx iris.Context) {
    // ctx.Proceed checks if myMiddleware called `ctx.Next()`
    // inside it and returns true if so, otherwise false.
    nextCalled := ctx.Proceed(myMiddleware)
    if !nextCalled {
        return
    }

    // else do the job here, it's allowed
}
```

And last, if you want to add a middleware on a specific method
and it depends on the next and the whole chain then you have to do it
using the `AfterActivation` like the example below:
*/
package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/mvc"
)

var cacheHandler = cache.Handler(10 * time.Second)

func main() {
	app := iris.New()
	// You don't have to use .Configure if you do it all in the main func
	// mvc.Configure and mvc.New(...).Configure() are just helpers to split
	// your code better, here we use the simplest form:
	m := mvc.New(app)
	m.Handle(&exampleController{})

	app.Run(iris.Addr(":8080"))
}

type exampleController struct{}

func (c *exampleController) AfterActivation(a mvc.AfterActivation) {
	// select the route based on the method name you want to
	// modify.
	index := a.GetRoute("Get")
	// just prepend the handler(s) as middleware(s) you want to use.
	// or append for "done" handlers.
	index.Handlers = append([]iris.Handler{cacheHandler}, index.Handlers...)
}

func (c *exampleController) Get() string {
	// refresh every 10 seconds and you will see different time output.
	now := time.Now().Format("Mon, Jan 02 2006 15:04:05")
	return "last time executed without cache: " + now
}

File: mvc/middleware/without-ctx-next/main.go

/*Package main is a simple example of the behavior change of the execution flow of the handlers,
normally we need the `ctx.Next()` to call the next handler in a route's handler chain,
but with the new `ExecutionRules` we can change this default behavior.
Please read below before continue.

The `Party#SetExecutionRules` alters the execution flow of the route handlers outside of the handlers themselves.

For example, if for some reason the desired result is the (done or all) handlers to be executed no matter what
even if no `ctx.Next()` is called in the previous handlers, including the begin(`Use`),
the main(`Handle`) and the done(`Done`) handlers themselves, then:
Party#SetExecutionRules(iris.ExecutionRules {
  Begin: iris.ExecutionOptions{Force: true},
  Main:  iris.ExecutionOptions{Force: true},
  Done:  iris.ExecutionOptions{Force: true},
})

Note that if `true` then the only remained way to "break" the handler chain is by `ctx.StopExecution()` now that `ctx.Next()` does not matter.

These rules are per-party, so if a `Party` creates a child one then the same rules will be applied to that as well.
Reset of these rules (before `Party#Handle`) can be done with `Party#SetExecutionRules(iris.ExecutionRules{})`.

The most common scenario for its use can be found inside Iris MVC Applications;
when we want the `Done` handlers of that specific mvc app's `Party`
to be executed but we don't want to add `ctx.Next()` on the `exampleController#EndRequest`*/
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/example") })

	// example := app.Party("/example")
	// example.SetExecutionRules && mvc.New(example) or...
	m := mvc.New(app.Party("/example"))

	// IMPORTANT
	// All options can be filled with Force:true, they all play nice together.
	m.Router.SetExecutionRules(iris.ExecutionRules{
		// Begin:  <- from `Use[all]` to `Handle[last]` future route handlers, execute all, execute all even if `ctx.Next()` is missing.
		// Main:   <- all `Handle` future route handlers, execute all >> >>.
		Done: iris.ExecutionOptions{Force: true}, // <- from `Handle[last]` to `Done[all]` future route handlers, execute all >> >>.
	})
	m.Router.Done(doneHandler)
	// m.Router.Done(...)
	// ...
	//

	m.Handle(&exampleController{})

	app.Run(iris.Addr(":8080"))
}

func doneHandler(ctx iris.Context) {
	ctx.WriteString("\nFrom Done Handler")
}

type exampleController struct{}

func (c *exampleController) Get() string {
	return "From Main Handler"
	// Note that here we don't binding the `Context`, and we don't call its `Next()`
	// function in order to call the `doneHandler`,
	// this is done automatically for us because we changed the execution rules with the `SetExecutionRules`.
	//
	// Therefore the final output is:
	// From Main Handler
	// From Done Handler
}
