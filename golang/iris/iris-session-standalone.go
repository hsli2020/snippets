package main // File: sessions/standalone/main.go

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type businessModel struct {
	Name string
}

func main() {
	app := iris.New()
	sess := sessions.New(sessions.Config{
		// Cookie string, the session's client cookie name, for example: "mysessionid"
		//
		// Defaults to "irissessionid"
		Cookie: "mysessionid",
		// it's time.Duration, from the time cookie is created, how long it can be alive?
		// 0 means no expire.
		// -1 means expire when browser closes
		// or set a value, like 2 hours:
		Expires: time.Hour * 2,
		// if you want to invalid cookies on different subdomains
		// of the same host, then enable it.
		// Defaults to false.
		DisableSubdomainPersistence: true,
		// AllowReclaim will allow to
		// Destroy and Start a session in the same request handler.
		// it removes the cookie for both `Request` and `ResponseWriter` while `Destroy`
		// or add a new cookie to `Request` while `Start`.
		//
		// Defaults to false.
		AllowReclaim: true,
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx) //set session values.
		s.Set("name", "iris")

		//test if setted here.
		ctx.Writef("All ok session setted to: %s", s.GetString("name"))

		// Set will set the value as-it-is,
		// if it's a slice or map
		// you will be able to change it on .Get directly!
		// Keep note that I don't recommend saving big data neither slices or maps on a session
		// but if you really need it then use the `SetImmutable` instead of `Set`.
		// Use `SetImmutable` consistently, it's slower.
		// Read more about muttable and immutable go types: https://stackoverflow.com/a/8021081
	})

	app.Get("/get", func(ctx iris.Context) {
		// get a specific value, as string, if not found then it returns just an empty string.
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		sess.Start(ctx).Delete("name") // delete a specific key
	})

	app.Get("/clear", func(ctx iris.Context) {
		sess.Start(ctx).Clear() // removes all entries.
	})

	app.Get("/update", func(ctx iris.Context) {
		sess.ShiftExpiration(ctx) // updates expire date.
	})

	app.Get("/destroy", func(ctx iris.Context) {
		sess.Destroy(ctx) //destroy, removes the entire session data and cookie
	})
	// Note about Destroy:
	//
	// You can destroy a session outside of a handler too, using the:
	// mySessions.DestroyByID
	// mySessions.DestroyAll

	// remember: slices and maps are muttable by-design
	// The `SetImmutable` makes sure that they will be stored and received
	// as immutable, so you can't change them directly by mistake.
	//
	// Use `SetImmutable` consistently, it's slower than `Set`.
	// Read more about muttable and immutable go types: https://stackoverflow.com/a/8021081
	app.Get("/set_immutable", func(ctx iris.Context) {
		business := []businessModel{{Name: "Edward"}, {Name: "value 2"}}
		s := sess.Start(ctx)
		s.SetImmutable("businessEdit", business)
		businessGet := s.Get("businessEdit").([]businessModel)

		// try to change it, if we used `Set` instead of `SetImmutable` this
		// change will affect the underline array of the session's value "businessEdit",
		// but now it will not.
		businessGet[0].Name = "Gabriel"

	})

	app.Get("/get_immutable", func(ctx iris.Context) {
		valSlice := sess.Start(ctx).Get("businessEdit")
		if valSlice == nil {
			ctx.HTML("please navigate to the <a href='/set_immutable'>/set_immutable</a> first")
			return
		}

		firstModel := valSlice.([]businessModel)[0]
		// businessGet[0].Name is equal to Edward initially
		if firstModel.Name != "Edward" {
			panic("immutable data cannot be changed from the caller without re-SetImmutable")
		}

		ctx.Writef("[]businessModel[0].Name remains: %s", firstModel.Name)

		// the name should remains "Edward"
	})

	app.Run(iris.Addr(":8080"))
}
