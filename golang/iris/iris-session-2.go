package main // File: sessions/overview/main.go

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
)

func secret(ctx iris.Context) {
	// Check if user is authenticated
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// Print secret message
	ctx.WriteString("The cake is a lie!")
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)

	// Authentication goes here

	// Set user as authenticated
	session.Set("authenticated", true)
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
}

func main() {
	app := iris.New()

	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)

	app.Run(iris.Addr(":8080"))
}

package main // File: sessions/securecookie/main.go

// developers can use any library to add a custom cookie encoder/decoder.
// At this example we use the gorilla's securecookie package:

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/gorilla/securecookie"
)

func newApp() *iris.Application {
	app := iris.New()

	cookieName := "mycustomsessionid"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	mySessions := sessions.New(sessions.Config{
		Cookie:       cookieName,
		Encode:       secureCookie.Encode,
		Decode:       secureCookie.Decode,
		AllowReclaim: true,
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		s := mySessions.Start(ctx) //set session values
		s.Set("name", "iris")

		//test if setted here
		ctx.Writef("All ok session setted to: %s", s.GetString("name"))
	})
	app.Get("/get", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		s := mySessions.Start(ctx)
		name := s.GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})
	app.Get("/delete", func(ctx iris.Context) {
		s := mySessions.Start(ctx) // delete a specific key
		s.Delete("name")
	})
	app.Get("/clear", func(ctx iris.Context) {
		mySessions.Start(ctx).Clear() // removes all entries
	})
	app.Get("/update", func(ctx iris.Context) {
		mySessions.ShiftExpiration(ctx) // updates expire date with a new date
	})
	app.Get("/destroy", func(ctx iris.Context) {
		mySessions.Destroy(ctx) //destroy, removes the entire session data and cookie
	})

	// Note about destroy:
	//
	// You can destroy a session outside of a handler too, using the:
	// mySessions.DestroyByID
	// mySessions.DestroyAll

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}
