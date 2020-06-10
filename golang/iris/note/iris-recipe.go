APIdoc

File: apidoc/yaag/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
)

/*
	go get github.com/betacraft/yaag/...
*/

type myXML struct {
	Result string `xml:"result"`
}

func main() {
	app := iris.New()

	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"result": "Hello World!"})
	})

	app.Get("/plain", func(ctx iris.Context) {
		ctx.Text("Hello World!")
	})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(myXML{Result: "Hello World!"})
	})

	app.Get("/complex", func(ctx iris.Context) {
		value := ctx.URLParam("key")
		ctx.JSON(iris.Map{"value": value})
	})

	// Run our HTTP Server.
	//
	// Documentation of "yaag" doesn't note the follow, but in Iris we are careful on what
	// we provide to you.
	//
	// Each incoming request results on re-generation and update of the "apidoc.html" file.
	// Recommentation:
	// Write tests that calls those handlers, save the generated "apidoc.html".
	// Turn off the yaag middleware when in production.
	//
	// Example usage:
	// Visit all paths and open the generated "apidoc.html" file to see the API's automated docs.
	app.Run(iris.Addr(":8080"))
}

Authentication

Basicauth

File: authentication/basicauth/main.go

package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		Users:   map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global app.Use(authentication) (or app.UseGlobal before the .Run)
	// to routes
	/*
		app.Get("/mysecret", authentication, h)
	*/

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })

	// to party

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}

	return app
}

func main() {
	app := newApp()
	// open http://localhost:8080/admin
	app.Run(iris.Addr(":8080"))
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter it will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

File: authentication/basicauth/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestBasicAuth(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	// redirects to /admin without basic auth
	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	// without basic auth
	e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)

	// with valid basic auth
	e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
	e.GET("/admin/profile").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/profile myusername:mypassword")
	e.GET("/admin/settings").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/settings myusername:mypassword")

	// with invalid basic auth
	e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
		Expect().Status(httptest.StatusUnauthorized)
}

Oauth2

File: authentication/oauth2/main.go

package main

// Any OAuth2 (even the pure golang/x/net/oauth2) package
// can be used with iris but at this example we will see the markbates' goth:
//
// $ go get github.com/markbates/goth/...
//
// This OAuth2 example works with sessions, so we will need
// to attach a session manager.
// Optionally: for even more secure session values,
// developers can use any third-party package to add a custom cookie encoder/decoder.
// At this example we will use the gorilla's securecookie:
//
// $ go get github.com/gorilla/securecookie
// Example of securecookie can be found at "sessions/securecookie" example folder.

// Notes:
// The whole example is converted by markbates/goth/example/main.go.
// It's tested with my own TWITTER application and it worked, even for localhost.
// I guess that everything else works as expected, all bugs reported by goth library's community
// are fixed in the time I wrote that example, have fun!
import (
	"errors"
	"os"
	"sort"

	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"

	"github.com/gorilla/securecookie" // optionally, used for session's encoder/decoder

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/amazon"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/bitbucket"
	"github.com/markbates/goth/providers/box"
	"github.com/markbates/goth/providers/dailymotion"
	"github.com/markbates/goth/providers/deezer"
	"github.com/markbates/goth/providers/digitalocean"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/dropbox"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/fitbit"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/heroku"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/intercom"
	"github.com/markbates/goth/providers/lastfm"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/meetup"
	"github.com/markbates/goth/providers/onedrive"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/paypal"
	"github.com/markbates/goth/providers/salesforce"
	"github.com/markbates/goth/providers/slack"
	"github.com/markbates/goth/providers/soundcloud"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/steam"
	"github.com/markbates/goth/providers/stripe"
	"github.com/markbates/goth/providers/twitch"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/uber"
	"github.com/markbates/goth/providers/wepay"
	"github.com/markbates/goth/providers/xero"
	"github.com/markbates/goth/providers/yahoo"
	"github.com/markbates/goth/providers/yammer"
)

var sessionsManager *sessions.Sessions

func init() {
	// attach a session manager
	cookieName := "mycustomsessionid"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	sessionsManager = sessions.New(sessions.Config{
		Cookie: cookieName,
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})
}

// These are some function helpers that you may use if you want

// GetProviderName is a function used to get the name of a provider
// for a given request. By default, this provider is fetched from
// the URL query string. If you provide it in a different way,
// assign your own function to this variable that returns the provider
// name for your request.
var GetProviderName = func(ctx iris.Context) (string, error) {
	// try to get it from the url param "provider"
	if p := ctx.URLParam("provider"); p != "" {
		return p, nil
	}

	// try to get it from the url PATH parameter "{provider} or :provider or {provider:string} or {provider:alphabetical}"
	if p := ctx.Params().Get("provider"); p != "" {
		return p, nil
	}

	// try to get it from context's per-request storage
	if p := ctx.Values().GetString("provider"); p != "" {
		return p, nil
	}
	// if not found then return an empty string with the corresponding error
	return "", errors.New("you must select a provider")
}

/*
BeginAuthHandler is a convenience handler for starting the authentication process.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

BeginAuthHandler will redirect the user to the appropriate authentication end-point
for the requested provider.

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
func BeginAuthHandler(ctx iris.Context) {
	url, err := GetAuthURL(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Writef("%v", err)
		return
	}

	ctx.Redirect(url, iris.StatusTemporaryRedirect)
}

/*
GetAuthURL starts the authentication process with the requested provided.
It will return a URL that should be used to send users to.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider" or from the context's value of "provider" key.

I would recommend using the BeginAuthHandler instead of doing all of these steps
yourself, but that's entirely up to you.
*/
func GetAuthURL(ctx iris.Context) (string, error) {
	providerName, err := GetProviderName(ctx)
	if err != nil {
		return "", err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth(SetState(ctx))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	session := sessionsManager.Start(ctx)
	session.Set(providerName, sess.Marshal())
	return url, nil
}

// SetState sets the state string associated with the given request.
// If no state string is associated with the request, one will be generated.
// This state is sent to the provider and can be retrieved during the
// callback.
var SetState = func(ctx iris.Context) string {
	state := ctx.URLParam("state")
	if len(state) > 0 {
		return state
	}

	return "state"

}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
var GetState = func(ctx iris.Context) string {
	return ctx.URLParam("state")
}

/*
CompleteUserAuth does what it says on the tin. It completes the authentication
process and fetches all of the basic information about the user from the provider.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
var CompleteUserAuth = func(ctx iris.Context) (goth.User, error) {
	providerName, err := GetProviderName(ctx)
	if err != nil {
		return goth.User{}, err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}
	session := sessionsManager.Start(ctx)
	value := session.GetString(providerName)
	if value == "" {
		return goth.User{}, errors.New("session value for " + providerName + " not found")
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, ctx.Request().URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	session.Set(providerName, sess.Marshal())
	return provider.FetchUser(sess)
}

// Logout invalidates a user session.
func Logout(ctx iris.Context) error {
	providerName, err := GetProviderName(ctx)
	if err != nil {
		return err
	}
	session := sessionsManager.Start(ctx)
	session.Delete(providerName)
	return nil
}

// End of the "some function helpers".

func main() {
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),

		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		fitbit.New(os.Getenv("FITBIT_KEY"), os.Getenv("FITBIT_SECRET"), "http://localhost:3000/auth/fitbit/callback"),
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3000/auth/gplus/callback"),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		spotify.New(os.Getenv("SPOTIFY_KEY"), os.Getenv("SPOTIFY_SECRET"), "http://localhost:3000/auth/spotify/callback"),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
		lastfm.New(os.Getenv("LASTFM_KEY"), os.Getenv("LASTFM_SECRET"), "http://localhost:3000/auth/lastfm/callback"),
		twitch.New(os.Getenv("TWITCH_KEY"), os.Getenv("TWITCH_SECRET"), "http://localhost:3000/auth/twitch/callback"),
		dropbox.New(os.Getenv("DROPBOX_KEY"), os.Getenv("DROPBOX_SECRET"), "http://localhost:3000/auth/dropbox/callback"),
		digitalocean.New(os.Getenv("DIGITALOCEAN_KEY"), os.Getenv("DIGITALOCEAN_SECRET"), "http://localhost:3000/auth/digitalocean/callback", "read"),
		bitbucket.New(os.Getenv("BITBUCKET_KEY"), os.Getenv("BITBUCKET_SECRET"), "http://localhost:3000/auth/bitbucket/callback"),
		instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), "http://localhost:3000/auth/instagram/callback"),
		intercom.New(os.Getenv("INTERCOM_KEY"), os.Getenv("INTERCOM_SECRET"), "http://localhost:3000/auth/intercom/callback"),
		box.New(os.Getenv("BOX_KEY"), os.Getenv("BOX_SECRET"), "http://localhost:3000/auth/box/callback"),
		salesforce.New(os.Getenv("SALESFORCE_KEY"), os.Getenv("SALESFORCE_SECRET"), "http://localhost:3000/auth/salesforce/callback"),
		amazon.New(os.Getenv("AMAZON_KEY"), os.Getenv("AMAZON_SECRET"), "http://localhost:3000/auth/amazon/callback"),
		yammer.New(os.Getenv("YAMMER_KEY"), os.Getenv("YAMMER_SECRET"), "http://localhost:3000/auth/yammer/callback"),
		onedrive.New(os.Getenv("ONEDRIVE_KEY"), os.Getenv("ONEDRIVE_SECRET"), "http://localhost:3000/auth/onedrive/callback"),

		//Pointed localhost.com to http://localhost:3000/auth/yahoo/callback through proxy as yahoo
		// does not allow to put custom ports in redirection uri
		yahoo.New(os.Getenv("YAHOO_KEY"), os.Getenv("YAHOO_SECRET"), "http://localhost.com"),
		slack.New(os.Getenv("SLACK_KEY"), os.Getenv("SLACK_SECRET"), "http://localhost:3000/auth/slack/callback"),
		stripe.New(os.Getenv("STRIPE_KEY"), os.Getenv("STRIPE_SECRET"), "http://localhost:3000/auth/stripe/callback"),
		wepay.New(os.Getenv("WEPAY_KEY"), os.Getenv("WEPAY_SECRET"), "http://localhost:3000/auth/wepay/callback", "view_user"),
		//By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		//in sandbox environment
		paypal.New(os.Getenv("PAYPAL_KEY"), os.Getenv("PAYPAL_SECRET"), "http://localhost:3000/auth/paypal/callback"),
		steam.New(os.Getenv("STEAM_KEY"), "http://localhost:3000/auth/steam/callback"),
		heroku.New(os.Getenv("HEROKU_KEY"), os.Getenv("HEROKU_SECRET"), "http://localhost:3000/auth/heroku/callback"),
		uber.New(os.Getenv("UBER_KEY"), os.Getenv("UBER_SECRET"), "http://localhost:3000/auth/uber/callback"),
		soundcloud.New(os.Getenv("SOUNDCLOUD_KEY"), os.Getenv("SOUNDCLOUD_SECRET"), "http://localhost:3000/auth/soundcloud/callback"),
		gitlab.New(os.Getenv("GITLAB_KEY"), os.Getenv("GITLAB_SECRET"), "http://localhost:3000/auth/gitlab/callback"),
		dailymotion.New(os.Getenv("DAILYMOTION_KEY"), os.Getenv("DAILYMOTION_SECRET"), "http://localhost:3000/auth/dailymotion/callback", "email"),
		deezer.New(os.Getenv("DEEZER_KEY"), os.Getenv("DEEZER_SECRET"), "http://localhost:3000/auth/deezer/callback", "email"),
		discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(os.Getenv("MEETUP_KEY"), os.Getenv("MEETUP_SECRET"), "http://localhost:3000/auth/meetup/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(os.Getenv("AUTH0_KEY"), os.Getenv("AUTH0_SECRET"), "http://localhost:3000/auth/auth0/callback", os.Getenv("AUTH0_DOMAIN")),
		xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/xero/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := make(map[string]string)
	m["amazon"] = "Amazon"
	m["bitbucket"] = "Bitbucket"
	m["box"] = "Box"
	m["dailymotion"] = "Dailymotion"
	m["deezer"] = "Deezer"
	m["digitalocean"] = "Digital Ocean"
	m["discord"] = "Discord"
	m["dropbox"] = "Dropbox"
	m["facebook"] = "Facebook"
	m["fitbit"] = "Fitbit"
	m["github"] = "Github"
	m["gitlab"] = "Gitlab"
	m["soundcloud"] = "SoundCloud"
	m["spotify"] = "Spotify"
	m["steam"] = "Steam"
	m["stripe"] = "Stripe"
	m["twitch"] = "Twitch"
	m["uber"] = "Uber"
	m["wepay"] = "Wepay"
	m["yahoo"] = "Yahoo"
	m["yammer"] = "Yammer"
	m["gplus"] = "Google Plus"
	m["heroku"] = "Heroku"
	m["instagram"] = "Instagram"
	m["intercom"] = "Intercom"
	m["lastfm"] = "Last FM"
	m["linkedin"] = "Linkedin"
	m["onedrive"] = "Onedrive"
	m["paypal"] = "Paypal"
	m["twitter"] = "Twitter"
	m["salesforce"] = "Salesforce"
	m["slack"] = "Slack"
	m["meetup"] = "Meetup.com"
	m["auth0"] = "Auth0"
	m["openid-connect"] = "OpenID Connect"
	m["xero"] = "Xero"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	// create our app,
	// set a view
	// set sessions
	// and setup the router for the showcase
	app := iris.New()

	// attach and build our templates
	app.RegisterView(iris.HTML("./templates", ".html"))

	// start of the router

	app.Get("/auth/{provider}/callback", func(ctx iris.Context) {

		user, err := CompleteUserAuth(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("%v", err)
			return
		}
		ctx.ViewData("", user)
		if err := ctx.View("user.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	app.Get("/logout/{provider}", func(ctx iris.Context) {
		Logout(ctx)
		ctx.Redirect("/", iris.StatusTemporaryRedirect)
	})

	app.Get("/auth/{provider}", func(ctx iris.Context) {
		// try to get the user without re-authenticating
		if gothUser, err := CompleteUserAuth(ctx); err == nil {
			ctx.ViewData("", gothUser)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			BeginAuthHandler(ctx)
		}
	})

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("", providerIndex)

		if err := ctx.View("index.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	// http://localhost:3000
	app.Run(iris.Addr("localhost:3000"))
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

File: authentication/oauth2/templates/index.html

{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}

File: authentication/oauth2/templates/user.html

<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>

Cache
Client Side
File: cache/client-side/main.go

// Package main shows how you can use the `WriteWithExpiration`
// based on the "modtime", if it's newer than the request header then
// it will refresh the contents, otherwise will let the client (99.9% the browser)
// to handle the cache mechanism, it's faster than iris.Cache because server-side
// has nothing to do and no need to store the responses in the memory.
package main

import (
	"time"

	"github.com/kataras/iris"
)

const refreshEvery = 10 * time.Second

func main() {
	app := iris.New()
	app.Use(iris.Cache304(refreshEvery))
	// same as:
	// app.Use(func(ctx iris.Context) {
	// 	now := time.Now()
	// 	if modified, err := ctx.CheckIfModifiedSince(now.Add(-refresh)); !modified && err == nil {
	// 		ctx.WriteNotModified()
	// 		return
	// 	}

	// 	ctx.SetLastModified(now)

	// 	ctx.Next()
	// })

	app.Get("/", greet)
	app.Run(iris.Addr(":8080"))
}

func greet(ctx iris.Context) {
	ctx.Header("X-Custom", "my  custom header")
	ctx.Writef("Hello World! %s", time.Now())
}

Simple
File: cache/simple/main.go

package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/cache"
)

var markdownContents = []byte(`## Hello Markdown

This is a sample of Markdown contents



Features
--------

All features of Sundown are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Routine safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.

	[this is a link](https://github.com/kataras/iris) `)

// Cache should not be used on handlers that contain dynamic data.
// Cache is a good and a must-feature on static content, i.e "about page" or for a whole blog site.
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", cache.Handler(10*time.Second), writeMarkdown)
	// saves its content on the first request and serves it instead of re-calculating the content.
	// After 10 seconds it will be cleared and resetted.

	app.Run(iris.Addr(":8080"))
}

func writeMarkdown(ctx iris.Context) {
	// tap multiple times the browser's refresh button and you will
	// see this println only once every 10 seconds.
	println("Handler executed. Content refreshed.")

	ctx.Markdown(markdownContents)
}

/* Note that `StaticWeb` does use the browser's disk caching by-default
therefore, register the cache handler AFTER any StaticWeb calls,
for a faster solution that server doesn't need to keep track of the response
navigate to https://github.com/kataras/iris/blob/master/_examples/cache/client-side/main.go */

Configuration
? click here to read the introduction section
From Configuration Structure
File: configuration/from-configuration-structure/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you want to modify the whole configuration.
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

	// or before Run:
	// app.Configure(iris.WithConfiguration(iris.Configuration{...}))
}

From Toml File
File: configuration/from-toml-file/configs/iris.tml

DisablePathCorrection = false

EnablePathEscape = false

FireMethodNotAllowed = true

DisableBodyConsumptionOnUnmarshal = false

TimeFormat = "Mon, 01 Jan 2006 15:04:05 GMT"

Charset = "UTF-8"



[Other]

	MyServerName = "iris"

File: configuration/from-toml-file/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.tml")))

	// or before run:
	// app.Configure(iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
	// app.Run(iris.Addr(":8080"))
}

From Yaml File
File: configuration/from-yaml-file/configs/iris.yml

DisablePathCorrection: false

EnablePathEscape: false

FireMethodNotAllowed: true

DisableBodyConsumptionOnUnmarshal: true

TimeFormat: Mon, 01 Jan 2006 15:04:05 GMT

Charset: UTF-8

File: configuration/from-yaml-file/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	// If iris.YAML's input string argument is "~" then it loads the configuration from the home directory
	// and can be shared between many iris instances.
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

	// or before run:
	// app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))
	// app.Run(iris.Addr(":8080"))
}

File: configuration/from-yaml-file/shared-configuration/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you share configuration between multiple iris instances.
	// This configuration file lives in your $HOME/iris.yml for unix hosts
	// or %HOMEDRIVE%+%HOMEPATH%/iris.yml for windows hosts, and you can modify it.
	app.Run(iris.Addr(":8080"), iris.WithGlobalConfiguration)
	// or before run:
	// app.Configure(iris.WithGlobalConfiguration)
	// app.Run(iris.Addr(":8080"))
}

Functional
File: configuration/functional/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you want to change some of the configuration's field.
	// Prefix: "With", code editors will help you navigate through all
	// configuration options without even a glitch to the documentation.

	app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"))

	// or before run:
	// app.Configure(iris.WithoutStartupLog, iris.WithCharset("UTF-8"))
	// app.Run(iris.Addr(":8080"))
}

Convert Handlers
Negroni Like
File: convert-handlers/negroni-like/main.go

package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	irisMiddleware := iris.FromStd(negronilikeTestMiddleware)
	app.Use(irisMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Home </h1>")
		// this will print an error,
		// this route's handler will never be executed because the middleware's criteria not passed.
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx iris.Context) {
		ctx.Writef("Hello world!")
		// this will print "OK. Hello world!".
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(iris.Addr(":8080"))
}

func negronilikeTestMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/ok" && r.Method == "GET" {
		w.Write([]byte("OK. "))
		next(w, r) // go to the next route's handler
		return
	}
	// else print an error and do not forward to the route's handler.
	w.WriteHeader(iris.StatusBadRequest)
	w.Write([]byte("Bad request"))
}

// Look "routing/custom-context" if you want to convert a custom handler with a custom Context
// to a context.Handler.

NetHTTP

File: convert-handlers/nethttp/main.go

package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	irisMiddleware := iris.FromStd(nativeTestMiddleware)
	app.Use(irisMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Home")
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx iris.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(iris.Addr(":8080"))
}

func nativeTestMiddleware(w http.ResponseWriter, r *http.Request) {
	println("Request path: " + r.URL.Path)
}

// Look "routing/custom-context" if you want to convert a custom handler with a custom Context
// to a context.Handler.

Real Usecase Raven

File: convert-handlers/real-usecase-raven/wrapping-the-router/main.go

package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/kataras/iris"

	"github.com/getsentry/raven-go"
)

// https://docs.sentry.io/clients/go/integrations/http/
func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hi")
	})

	// Example for WrapRouter is already here:
	// https://github.com/kataras/iris/blob/master/_examples/routing/custom-wrapper/main.go#L53
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, irisRouter http.HandlerFunc) {
		// Exactly the same source code:
		// https://github.com/getsentry/raven-go/blob/379f8d0a68ca237cf8893a1cdfd4f574125e2c51/http.go#L70

		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		irisRouter(w, r)
	})

	app.Run(iris.Addr(":8080"))
}

File: convert-handlers/real-usecase-raven/writing-middleware/main.go

package main

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/kataras/iris"

	"github.com/getsentry/raven-go"
)

// At this example you will see how to convert any net/http middleware
// that has the form of `(HandlerFunc) HandlerFunc`.
// If the `raven.RecoveryHandler` had the form of
// `(http.HandlerFunc)` or `(http.HandlerFunc, next http.HandlerFunc)`
// you could just use the `irisMiddleware := iris.FromStd(nativeHandler)`
// but it doesn't, however as you already know Iris can work with net/http directly
// because of the `ctx.ResponseWriter()` and `ctx.Request()` are the original
// http.ResponseWriter and *http.Request.
// (this one is a big advantage, as a result you can use Iris for ever :)).
//
// The source code of the native middleware does not change at all.
// https://github.com/getsentry/raven-go/blob/379f8d0a68ca237cf8893a1cdfd4f574125e2c51/http.go#L70
// The only addition is the Line 18 and Line 39 (instead of handler(w,r))
// and you have a new iris middleware ready to use!
func irisRavenMiddleware(ctx iris.Context) {
	w, r := ctx.ResponseWriter(), ctx.Request()

	defer func() {
		if rval := recover(); rval != nil {
			debug.PrintStack()
			rvalStr := fmt.Sprint(rval)
			packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
			raven.Capture(packet, nil)
			w.WriteHeader(iris.StatusInternalServerError)
		}
	}()

	ctx.Next()
}

// https://docs.sentry.io/clients/go/integrations/http/
func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func main() {
	app := iris.New()
	app.Use(irisRavenMiddleware)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hi")
	})

	app.Run(iris.Addr(":8080"))
}

Cookies

Basic

File: cookies/basic/main.go

package main

import "github.com/kataras/iris"

func newApp() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, value) // <--
		// Alternatively: ctx.SetCookie(&http.Cookie{...})
		//
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))
		//
		// If you want to be visible only to current request path:
		// (note that client should be responsible for that if server sent an empty cookie's path, all browsers are compatible)
		// ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)
		// More:
		//                              iris.CookieExpires(time.Duration)
		//                              iris.CookieHTTPOnly(false)

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name) // <--
		// If you want more than the value then:
		// cookie, err := ctx.Request().Cookie(name)
		// if err != nil {
		//  handle error.
		// }

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main() {
	app := newApp()

	// GET:    http://localhost:8080/cookies/my_name/my_value
	// GET:    http://localhost:8080/cookies/my_name
	// DELETE: http://localhost:8080/cookies/my_name
	app.Run(iris.Addr(":8080"))
}

File: cookies/basic/main_test.go

package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestCookiesBasic(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.URL("http://example.com"))

	cookieName, cookieValue := "my_cookie_name", "my_cookie_value"

	// Test Set A Cookie.
	t1 := e.GET(fmt.Sprintf("/cookies/%s/%s", cookieName, cookieValue)).Expect().Status(httptest.StatusOK)
	t1.Cookie(cookieName).Value().Equal(cookieValue) // validate cookie's existence, it should be there now.
	t1.Body().Contains(cookieValue)

	// Test Retrieve A Cookie.
	t2 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t2.Body().Equal(cookieValue)

	// Test Remove A Cookie.
	t3 := e.DELETE(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t3.Body().Contains(cookieName)

	t4 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t4.Cookies().Empty()
	t4.Body().Empty()
}

Securecookie

File: cookies/securecookie/main.go

package main

// developers can use any library to add a custom cookie encoder/decoder.
// At this example we use the gorilla's securecookie package:
// $ go get github.com/gorilla/securecookie
// $ go run main.go

import (
	"github.com/kataras/iris"

	"github.com/gorilla/securecookie"
)

var (
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

func newApp() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, value, iris.CookieEncode(sc.Encode)) // <--

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name, iris.CookieDecode(sc.Decode)) // <--

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

File: cookies/securecookie/main_test.go

package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestCookiesBasic(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.URL("http://example.com"))

	cookieName, cookieValue := "my_cookie_name", "my_cookie_value"

	// Test Set A Cookie.
	t1 := e.GET(fmt.Sprintf("/cookies/%s/%s", cookieName, cookieValue)).Expect().Status(httptest.StatusOK)
	// note that this will not work because it doesn't always returns the same value:
	// cookieValueEncoded, _ := sc.Encode(cookieName, cookieValue)
	t1.Cookie(cookieName).Value().NotEqual(cookieValue) // validate cookie's existence and value is not on its raw form.
	t1.Body().Contains(cookieValue)

	// Test Retrieve A Cookie.
	t2 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t2.Body().Equal(cookieValue)

	// Test Remove A Cookie.
	t3 := e.DELETE(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t3.Body().Contains(cookieName)

	t4 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t4.Cookies().Empty()
	t4.Body().Empty()
}

Experimental Handlers

Casbin

File: experimental-handlers/casbin/middleware/casbinmodel.conf

[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")

File: experimental-handlers/casbin/middleware/casbinpolicy.csv

p, alice, /dataset1/*, GET
p, alice, /dataset1/resource1, POST
p, bob, /dataset2/resource1, *
p, bob, /dataset2/resource2, GET
p, bob, /dataset2/folder1/*, POST

File: experimental-handlers/casbin/middleware/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/casbin/casbin"
	cm "github.com/iris-contrib/middleware/casbin"
)

// $ go get github.com/casbin/casbin
// $ go run main.go

// Enforcer maps the model and the policy for the casbin service, we use this variable on the main_test too.
var Enforcer = casbin.NewEnforcer("casbinmodel.conf", "casbinpolicy.csv")

func newApp() *iris.Application {
	casbinMiddleware := cm.New(Enforcer)

	app := iris.New()
	app.Use(casbinMiddleware.ServeHTTP)

	app.Get("/", hi)

	app.Get("/dataset1/{p:path}", hi) // p, alice, /dataset1/*, GET

	app.Post("/dataset1/resource1", hi)

	app.Get("/dataset2/resource2", hi)
	app.Post("/dataset2/folder1/{p:path}", hi)

	app.Any("/dataset2/resource1", hi)

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func hi(ctx iris.Context) {
	ctx.Writef("Hello %s", cm.Username(ctx.Request()))
}

File: experimental-handlers/casbin/middleware/main_test.go

package main

import (
	"testing"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris/httptest"
)

func TestCasbinMiddleware(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.Debug(false))

	type ttcasbin struct {
		username string
		path     string
		method   string
		status   int
	}

	tt := []ttcasbin{
		{"alice", "/dataset1/resource1", "GET", 200},
		{"alice", "/dataset1/resource1", "POST", 200},
		{"alice", "/dataset1/resource2", "GET", 200},
		{"alice", "/dataset1/resource2", "POST", 404},

		{"bob", "/dataset2/resource1", "GET", 200},
		{"bob", "/dataset2/resource1", "POST", 200},
		{"bob", "/dataset2/resource1", "DELETE", 200},
		{"bob", "/dataset2/resource2", "GET", 200},
		{"bob", "/dataset2/resource2", "POST", 404},
		{"bob", "/dataset2/resource2", "DELETE", 404},

		{"bob", "/dataset2/folder1/item1", "GET", 404},
		{"bob", "/dataset2/folder1/item1", "POST", 200},
		{"bob", "/dataset2/folder1/item1", "DELETE", 404},
		{"bob", "/dataset2/folder1/item2", "GET", 404},
		{"bob", "/dataset2/folder1/item2", "POST", 200},
		{"bob", "/dataset2/folder1/item2", "DELETE", 404},
	}

	for _, tt := range tt {
		check(e, tt.method, tt.path, tt.username, tt.status)
	}
}

func check(e *httpexpect.Expect, method, path, username string, status int) {
	e.Request(method, path).WithBasicAuth(username, "password").Expect().Status(status)
}

File: experimental-handlers/casbin/wrapper/casbinmodel.conf

[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")

File: experimental-handlers/casbin/wrapper/casbinpolicy.csv

p, alice, /dataset1/*, GET
p, alice, /dataset1/resource1, POST
p, bob, /dataset2/resource1, *
p, bob, /dataset2/resource2, GET
p, bob, /dataset2/folder1/*, POST
p, dataset1_admin, /dataset1/*, *
g, cathrin, dataset1_admin

File: experimental-handlers/casbin/wrapper/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/casbin/casbin"
	cm "github.com/iris-contrib/middleware/casbin"
)

// $ go get github.com/casbin/casbin
// $ go run main.go

// Enforcer maps the model and the policy for the casbin service, we use this variable on the main_test too.
var Enforcer = casbin.NewEnforcer("casbinmodel.conf", "casbinpolicy.csv")

func newApp() *iris.Application {
	casbinMiddleware := cm.New(Enforcer)

	app := iris.New()
	app.WrapRouter(casbinMiddleware.Wrapper())

	app.Get("/", hi)

	app.Any("/dataset1/{p:path}", hi) // p, dataset1_admin, /dataset1/*, * && p, alice, /dataset1/*, GET

	app.Post("/dataset1/resource1", hi)

	app.Get("/dataset2/resource2", hi)
	app.Post("/dataset2/folder1/{p:path}", hi)

	app.Any("/dataset2/resource1", hi)

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func hi(ctx iris.Context) {
	ctx.Writef("Hello %s", cm.Username(ctx.Request()))
}

File: experimental-handlers/casbin/wrapper/main_test.go

package main

import (
	"testing"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris/httptest"
)

func TestCasbinWrapper(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	type ttcasbin struct {
		username string
		path     string
		method   string
		status   int
	}

	tt := []ttcasbin{
		{"alice", "/dataset1/resource1", "GET", 200},
		{"alice", "/dataset1/resource1", "POST", 200},
		{"alice", "/dataset1/resource2", "GET", 200},
		{"alice", "/dataset1/resource2", "POST", 403},

		{"bob", "/dataset2/resource1", "GET", 200},
		{"bob", "/dataset2/resource1", "POST", 200},
		{"bob", "/dataset2/resource1", "DELETE", 200},
		{"bob", "/dataset2/resource2", "GET", 200},
		{"bob", "/dataset2/resource2", "POST", 403},
		{"bob", "/dataset2/resource2", "DELETE", 403},

		{"bob", "/dataset2/folder1/item1", "GET", 403},
		{"bob", "/dataset2/folder1/item1", "POST", 200},
		{"bob", "/dataset2/folder1/item1", "DELETE", 403},
		{"bob", "/dataset2/folder1/item2", "GET", 403},
		{"bob", "/dataset2/folder1/item2", "POST", 200},
		{"bob", "/dataset2/folder1/item2", "DELETE", 403},
	}

	for _, tt := range tt {
		check(e, tt.method, tt.path, tt.username, tt.status)
	}

	ttAdmin := []ttcasbin{
		{"cathrin", "/dataset1/item", "GET", 200},
		{"cathrin", "/dataset1/item", "POST", 200},
		{"cathrin", "/dataset1/item", "DELETE", 200},
		{"cathrin", "/dataset2/item", "GET", 403},
		{"cathrin", "/dataset2/item", "POST", 403},
		{"cathrin", "/dataset2/item", "DELETE", 403},
	}

	for _, tt := range ttAdmin {
		check(e, tt.method, tt.path, tt.username, tt.status)
	}

	Enforcer.DeleteRolesForUser("cathrin")

	ttAdminDeleted := []ttcasbin{
		{"cathrin", "/dataset1/item", "GET", 403},
		{"cathrin", "/dataset1/item", "POST", 403},
		{"cathrin", "/dataset1/item", "DELETE", 403},
		{"cathrin", "/dataset2/item", "GET", 403},
		{"cathrin", "/dataset2/item", "POST", 403},
		{"cathrin", "/dataset2/item", "DELETE", 403},
	}

	for _, tt := range ttAdminDeleted {
		check(e, tt.method, tt.path, tt.username, tt.status)
	}

}

func check(e *httpexpect.Expect, method, path, username string, status int) {
	e.Request(method, path).WithBasicAuth(username, "password").Expect().Status(status)
}

Cloudwatch

File: experimental-handlers/cloudwatch/simple/main.go

package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	cw "github.com/iris-contrib/middleware/cloudwatch"
)

// $ go get github.com/aws/aws-sdk-go/...
// $ go run main.go

func main() {
	app := iris.New()
	app.Use(cw.New("us-east-1", "test").ServeHTTP)

	app.Get("/", func(ctx iris.Context) {
		put := cw.GetPutFunc(ctx)

		put([]*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("MyMetric"),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("ThingOne"),
						Value: aws.String("something"),
					},
					{
						Name:  aws.String("ThingTwo"),
						Value: aws.String("other"),
					},
				},
				Timestamp: aws.Time(time.Now()),
				Unit:      aws.String("Count"),
				Value:     aws.Float64(42),
			},
		})

		ctx.StatusCode(iris.StatusOK)
		ctx.Text("success!\n")
	})

	// http://localhost:8080
	// should give: NoCredentialProviders
	// which is correct, you have to authorize your aws, we asumme that you know how to.
	app.Run(iris.Addr(":8080"))
}

Cors

File: experimental-handlers/cors/simple/main.go

package main

// go get -u github.com/iris-contrib/middleware/...

import (
	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/cors"
)

func main() {
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		v1.Get("/home", func(ctx iris.Context) {
			ctx.WriteString("Hello from /home")
		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Run(iris.Addr("localhost:8080"))
}

Csrf

File: experimental-handlers/csrf/main.go

// This middleware provides Cross-Site Request Forgery
// protection.
//
// It securely generates a masked (unique-per-request) token that
// can be embedded in the HTTP response (e.g. form field or HTTP header).
// The original (unmasked) token is stored in the session, which is inaccessible
// by an attacker (provided you are using HTTPS). Subsequent requests are
// expected to include this token, which is compared against the session token.
// Requests that do not provide a matching token are served with a HTTP 403
// 'Forbidden' error response.
package main

// $ go get -u github.com/iris-contrib/middleware/...

import (
	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/csrf"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	// Note that the authentication key provided should be 32 bytes
	// long and persist across application restarts.
	protect := csrf.Protect([]byte("9AB0F421E53A477C084477AEA06096F5"),
		csrf.Secure(false)) // Defaults to true, but pass `false` while no https (devmode).

	users := app.Party("/user", protect)
	{
		users.Get("/signup", getSignupForm)
		// // POST requests without a valid token will return a HTTP 403 Forbidden.
		users.Post("/signup", postSignupForm)
	}

	// GET: http://localhost:8080/user/signup
	// POST: http://localhost:8080/user/signup
	app.Run(iris.Addr(":8080"))
}

func getSignupForm(ctx iris.Context) {
	// views/user/signup.html just needs a {{ .csrfField }} template tag for
	// csrf.TemplateField to inject the CSRF token into. Easy!
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("user/signup.html")

	// We could also retrieve the token directly from csrf.Token(ctx) and
	// set it in the request header - ctx.GetHeader("X-CSRF-Token", token)
	// This is useful if you're sending JSON to clients or a front-end JavaScript
	// framework.
}

func postSignupForm(ctx iris.Context) {
	ctx.Writef("You're welcome mate!")
}

File: experimental-handlers/csrf/views/user/signup.html

<form method="POST" action="/user/signup">
    {{ .csrfField }}
<button type="submit">Proceed</button>
</form>

JWT

File: experimental-handlers/jwt/main.go

// iris provides some basic middleware, most for your learning courve.
// You can use any net/http compatible middleware with iris.FromStd wrapper.
//
// JWT net/http video tutorial for golang newcomers: https://www.youtube.com/watch?v=dgJFeqeXVKw
//
// This middleware is the only one cloned from external source: https://github.com/auth0/go-jwt-middleware
// (because it used "context" to define the user but we don't need that so a simple iris.FromStd wouldn't work as expected.)
package main

// $ go get -u github.com/dgrijalva/jwt-go
// $ go run main.go

import (
	"github.com/kataras/iris"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func myHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	ctx.Writef("%s", user.Signature)
}

func main() {
	app := iris.New()

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Use(jwtHandler.Serve)

	app.Get("/ping", myHandler)
	app.Run(iris.Addr("localhost:3001"))
} // don't forget to look ../jwt_test.go to seee how to set your own custom claims

Newrelic

File: experimental-handlers/newrelic/simple/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/newrelic"
)

func main() {
	app := iris.New()
	config := newrelic.Config("APP_SERVER_NAME", "NEWRELIC_LICENSE_KEY")
	config.Enabled = true
	m, err := newrelic.New(config)
	if err != nil {
		app.Logger().Fatal(err)
	}
	app.Use(m.ServeHTTP)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("success!\n")
	})

	app.Run(iris.Addr(":8080"))
}

Prometheus

File: experimental-handlers/prometheus/simple/main.go

package main

import (
	"math/rand"
	"time"

	"github.com/kataras/iris"

	prometheusMiddleware "github.com/iris-contrib/middleware/prometheus"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	app := iris.New()
	m := prometheusMiddleware.New("serviceName", 300, 1200, 5000)

	app.Use(m.ServeHTTP)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		// error code handlers are not sharing the same middleware as other routes, so we have
		// to call them inside their body.
		m.ServeHTTP(ctx)

		ctx.Writef("Not Found")
	})

	app.Get("/", func(ctx iris.Context) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ctx.Writef("Slept for %d milliseconds", sleep)
	})

	app.Get("/metrics", iris.FromStd(prometheus.Handler()))

	// http://localhost:8080/
	// http://localhost:8080/anotfound
	// http://localhost:8080/metrics
	app.Run(iris.Addr(":8080"))
}

Secure

File: experimental-handlers/secure/simple/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/iris-contrib/middleware/secure"
)

func main() {
	s := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                                                                                                                         // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		SSLRedirect:             true,                                                                                                                                                // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                                                                                                                               // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                                                                                                                                   // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"},                                                                                                     // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                                                                                                                           // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                                                                                                                                // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                                                                                                                                // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                                                                                                                               // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                                                                                                                                // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                                                                                                                        // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option.
		ContentTypeNosniff:      true,                                                                                                                                                // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXSSFilter:        true,                                                                                                                                                // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		ContentSecurityPolicy:   "default-src 'self'",                                                                                                                                // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		PublicKey:               `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".

		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	app := iris.New()
	app.Use(s.Serve)

	app.Get("/home", func(ctx iris.Context) {
		ctx.Writef("Hello from /home")
	})

	app.Run(iris.Addr(":8080"))
}

Tollboothic

File: experimental-handlers/tollboothic/limit-handler/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/tollboothic"
)

// $ go get github.com/didip/tollbooth
// $ go run main.go

func main() {
	app := iris.New()

	limiter := tollbooth.NewLimiter(1, nil)
	//
	// or create a limiter with expirable token buckets
	// This setting means:
	// create a 1 request/second limiter and
	// every token bucket in it will expire 1 hour after it was initially set.
	// limiter := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	app.Get("/", tollboothic.LimitHandler(limiter), func(ctx iris.Context) {
		ctx.HTML("<b>Hello, world!</b>")
	})

	app.Run(iris.Addr(":8080"))
}

// Read more at: https://github.com/didip/tollbooth

File Server

Basic

File: file-server/basic/assets/css/main.css

body {
    background-color: black;
}

File: file-server/basic/assets/index.html

<h1>Hello index</h1>

File: file-server/basic/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Favicon("./assets/favicon.ico")

	// enable gzip, optionally:
	// if used before the `StaticXXX` handlers then
	// the content byte range feature is gone.
	// recommend: turn off for large files especially
	// when server has low memory,
	// turn on for medium-sized files
	// or for large-sized files if they are zipped already,
	// i.e "zippedDir/file.gz"
	//
	// app.Use(iris.Gzip)

	// first parameter is the request path
	// second is the system directory
	//
	// app.StaticWeb("/css", "./assets/css")
	// app.StaticWeb("/js", "./assets/js")
	//
	app.StaticWeb("/static", "./assets")

	// http://localhost:8080/static/css/main.css
	// http://localhost:8080/static/js/jquery-2.1.1.js
	// http://localhost:8080/static/favicon.ico
	app.Run(iris.Addr(":8080"))

	// Note:
	// Routing doesn't allows something .StaticWeb("/", "./assets")
	//
	// To see how you can wrap the router in order to achieve
	// wildcard on root path, see "single-page-application".
}

Embedding Files Into App

File: file-server/embedding-files-into-app/main.go

package main

import (
	"github.com/kataras/iris"
)

// Follow these steps first:
// $ go get -u github.com/shuLhan/go-bindata/...
// $ go-bindata ./assets/...
// $ go build
// $ ./embedding-files-into-app
// "physical" files are not used, you can delete the "assets" folder and run the example.
//
// See `file-server/embedding-gziped-files-into-app` example as well.
func newApp() *iris.Application {
	app := iris.New()

	app.StaticEmbedded("/static", "./assets", Asset, AssetNames)

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080/static/css/bootstrap.min.css
	// http://localhost:8080/static/js/jquery-2.1.1.js
	// http://localhost:8080/static/favicon.ico
	app.Run(iris.Addr(":8080"))
}

File: file-server/embedding-files-into-app/main_test.go

package main

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
)

type resource string

// content types that are used in the ./assets,
// we could use the detectContentType that iris do but it's better
// to do it manually so we can test if that returns the correct result on embedding files.
func (r resource) contentType() string {
	switch filepath.Ext(r.String()) {
	case ".js":
		return "application/javascript"
	case ".css":
		return "text/css"
	case ".ico":
		return "image/x-icon"
	case ".html":
		return "text/html"
	default:
		return "text/plain"
	}
}

func (r resource) String() string {
	return string(r)
}

func (r resource) strip(strip string) string {
	s := r.String()
	return strings.TrimPrefix(s, strip)
}

func (r resource) loadFromBase(dir string) string {
	filename := r.String()

	filename = r.strip("/static")

	fullpath := filepath.Join(dir, filename)

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(fullpath + " failed with error: " + err.Error())
	}

	result := string(b)

	if runtime.GOOS != "windows" {
		result = strings.Replace(result, "\n", "\r\n", -1)
	}
	return result
}

var urls = []resource{
	"/static/css/bootstrap.min.css",
	"/static/js/jquery-2.1.1.js",
	"/static/favicon.ico",
}

// if bindata's values matches with the assets/... contents
// and secondly if the StaticEmbedded had successfully registered
// the routes and gave the correct response.
func TestEmbeddingFilesIntoApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	if runtime.GOOS != "windows" {
		// remove the embedded static favicon for !windows,
		// it should be built for unix-specific in order to be work
		urls = urls[0 : len(urls)-1]
	}

	for _, u := range urls {
		url := u.String()
		contents := u.loadFromBase("./assets")

		e.GET(url).Expect().
			Status(httptest.StatusOK).
			ContentType(u.contentType(), app.ConfigurationReadOnly().GetCharset()).
			Body().Equal(contents)
	}
}

Embedding Gziped Files Into App

File: file-server/embedding-gziped-files-into-app/main.go

package main

import (
	"github.com/kataras/iris"
)

// NOTE: need different tool than the "embedding-files-into-app" example.
//
// Follow these steps first:
// $ go get -u github.com/kataras/bindata/cmd/bindata
// $ bindata ./assets/...
// $ go build
// $ ./embedding-gziped-files-into-app
// "physical" files are not used, you can delete the "assets" folder and run the example.

func newApp() *iris.Application {
	app := iris.New()

	// Note the `GzipAsset` and `GzipAssetNames` are different from `go-bindata`'s `Asset` and `AssetNames,
	// that means that you can use both `go-bindata` and `bindata` tools,
	// the `go-bindata` can be used for the view engine's `Binary` method
	// and the `bindata` with the `StaticEmbeddedGzip` (x8 times faster than the StaticEmbeded with `go-bindata`).
	app.StaticEmbeddedGzip("/static", "./assets", GzipAsset, GzipAssetNames)

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080/static/css/bootstrap.min.css
	// http://localhost:8080/static/js/jquery-2.1.1.js
	// http://localhost:8080/static/favicon.ico
	app.Run(iris.Addr(":8080"))
}

File: file-server/embedding-gziped-files-into-app/main_test.go

package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
	"github.com/klauspost/compress/gzip"
)

type resource string

// content types that are used in the ./assets,
// we could use the detectContentType that iris do but it's better
// to do it manually so we can test if that returns the correct result on embedding files.
func (r resource) contentType() string {
	switch filepath.Ext(r.String()) {
	case ".js":
		return "application/javascript"
	case ".css":
		return "text/css"
	case ".ico":
		return "image/x-icon"
	case ".html":
		return "text/html"
	default:
		return "text/plain"
	}
}

func (r resource) String() string {
	return string(r)
}

func (r resource) strip(strip string) string {
	s := r.String()
	return strings.TrimPrefix(s, strip)
}

func (r resource) loadFromBase(dir string) string {
	filename := r.String()

	filename = r.strip("/static")

	fullpath := filepath.Join(dir, filename)

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(fullpath + " failed with error: " + err.Error())
	}
	result := string(b)

	if runtime.GOOS != "windows" {
		result = strings.Replace(result, "\n", "\r\n", -1)
	}
	return result
}

var urls = []resource{
	"/static/css/bootstrap.min.css",
	"/static/js/jquery-2.1.1.js",
	"/static/favicon.ico",
}

// if bindata's values matches with the assets/... contents
// and secondly if the StaticEmbedded had successfully registered
// the routes and gave the correct response.
func TestEmbeddingGzipFilesIntoApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	if runtime.GOOS != "windows" {
		// remove the embedded static favicon for !windows,
		// it should be built for unix-specific in order to be work
		urls = urls[0 : len(urls)-1]
	}

	for i, u := range urls {
		url := u.String()
		rawContents := u.loadFromBase("./assets")

		response := e.GET(url).Expect()
		response.ContentType(u.contentType(), app.ConfigurationReadOnly().GetCharset())

		if expected, got := response.Raw().StatusCode, httptest.StatusOK; expected != got {
			t.Fatalf("[%d] of '%s': expected %d status code but got %d", i, url, expected, got)
		}

		func() {
			reader, err := gzip.NewReader(bytes.NewBuffer(response.Content))
			defer reader.Close()
			if err != nil {
				t.Fatalf("[%d] of '%s': %v", i, url, err)
			}
			buf := new(bytes.Buffer)
			reader.WriteTo(buf)
			if rawContents != buf.String() {
				t.Fatalf("[%d] of '%s': expected body:\n%s but got:\n%s", i, url, rawContents, buf.String())
			}
		}()
	}
}

Favicon

File: file-server/favicon/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// This will serve the ./static/favicons/favicon.ico to: localhost:8080/favicon.ico
	app.Favicon("./static/favicons/favicon.ico")

	// app.Favicon("./static/favicons/favicon.\\.ico", "/favicon_16_16.ico")
	// This will serve the ./static/favicons/favicon.ico to: localhost:8080/favicon_16_16.ico

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(`<a href="/favicon.ico"> press here to see the favicon.ico</a>.
		 At some browsers like chrome, it should be visible at the top-left side of the browser's window,
		 because some browsers make requests to the /favicon.ico automatically,
		  so iris serves your favicon in that path too (you can change it).`)
	}) // if favicon doesn't show to you, try to clear your browser's cache.

	app.Run(iris.Addr(":8080"))
}

Send Files

File: file-server/send-files/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		file := "./files/first.zip"
		ctx.SendFile(file, "c.zip")
	})

	app.Run(iris.Addr(":8080"))
}

Single Page Application

File: file-server/single-page-application/basic/main.go

package main

import (
	"github.com/kataras/iris"
)

// same as embedded-single-page-application but without go-bindata, the files are "physical" stored in the
// current system directory.

var page = struct {
	Title string
}{"Welcome"}

func newApp() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./public", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	// or just serve index.html as it is:
	// app.Get("/{f:path}", func(ctx iris.Context) {
	// 	ctx.ServeFile("index.html", false)
	// })

	assetHandler := app.StaticHandler("./public", false, false)
	// as an alternative of SPA you can take a look at the /routing/dynamic-path/root-wildcard
	// example too
	app.SPA(assetHandler)

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080
	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	app.Run(iris.Addr(":8080"))
}

File: file-server/single-page-application/basic/main_test.go

package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
)

type resource string

func (r resource) String() string {
	return string(r)
}

func (r resource) strip(strip string) string {
	s := r.String()
	return strings.TrimPrefix(s, strip)
}

func (r resource) loadFromBase(dir string) string {
	filename := r.String()

	if filename == "/" {
		filename = "/index.html"
	}

	fullpath := filepath.Join(dir, filename)

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(fullpath + " failed with error: " + err.Error())
	}

	result := string(b)

	return result
}

var urls = []resource{
	"/",
	"/index.html",
	"/app.js",
	"/css/main.css",
}

func TestSPA(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.Debug(false))

	for _, u := range urls {
		url := u.String()
		contents := u.loadFromBase("./public")
		contents = strings.Replace(contents, "{{ .Page.Title }}", page.Title, 1)

		e.GET(url).Expect().
			Status(httptest.StatusOK).
			Body().Equal(contents)
	}
}

File: file-server/single-page-application/basic/public/app.js

window.alert("app.js loaded from \"/");

File: file-server/single-page-application/basic/public/css/main.css

body {
    background-color: black;
}

File: file-server/single-page-application/basic/public/index.html

<html>

<head>
    <title>{{ .Page.Title }}</title>
</head>

<body>
    <h1> Hello from index.html </h1>


    <script src="/app.js">  </script>
</body>

</html>

File: file-server/single-page-application/embedded-single-page-application/main.go

package main

import (
	"github.com/kataras/iris"
)

// $ go get -u github.com/shuLhan/go-bindata/...
// $ go-bindata ./public/...
// $ go build
// $ ./embedded-single-page-application

var page = struct {
	Title string
}{"Welcome"}

func newApp() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./public", ".html").Binary(Asset, AssetNames))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	assetHandler := iris.StaticEmbeddedHandler("./public", Asset, AssetNames, false) // keep that false if you use the `go-bindata` tool.
	// as an alternative of SPA you can take a look at the /routing/dynamic-path/root-wildcard
	// example too
	// or
	// app.StaticEmbedded if you don't want to redirect on index.html and simple serve your SPA app (recommended).

	// public/index.html is a dynamic view, it's handlded by root,
	// and we don't want to be visible as a raw data, so we will
	// the return value of `app.SPA` to modify the `IndexNames` by;
	app.SPA(assetHandler).AddIndexName("index.html")

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080
	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	app.Run(iris.Addr(":8080"))
}

// Note that app.Use/UseGlobal/Done will be executed
// only to the registered routes like our index (app.Get("/", ..)).
// The file server is clean, but you can still add middleware to that by wrapping its "assetHandler".
//
// With this method, unlike StaticWeb("/" , "./public") which is not working by-design anymore,
// all custom http errors and all routes are working fine with a file server that is registered
// to the root path of the server.

File: file-server/single-page-application/embedded-single-page-application/main_test.go

package main

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
)

type resource string

func (r resource) contentType() string {
	switch filepath.Ext(r.String()) {
	case ".js":
		return "application/javascript"
	case ".css":
		return "text/css"
	default:
		return "text/html"
	}
}

func (r resource) String() string {
	return string(r)
}

func (r resource) strip(strip string) string {
	s := r.String()
	return strings.TrimPrefix(s, strip)
}

func (r resource) loadFromBase(dir string) string {
	filename := r.String()

	if filename == "/" {
		filename = "/index.html"
	}

	fullpath := filepath.Join(dir, filename)

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(fullpath + " failed with error: " + err.Error())
	}
	result := string(b)
	if runtime.GOOS != "windows" {
		result = strings.Replace(result, "\n", "\r\n", -1)
	}
	return result
}

var urls = []resource{
	"/",
	"/index.html",
	"/app.js",
	"/css/main.css",
}

func TestSPAEmbedded(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	for _, u := range urls {
		url := u.String()
		contents := u.loadFromBase("./public")
		contents = strings.Replace(contents, "{{ .Page.Title }}", page.Title, 1)

		e.GET(url).Expect().
			Status(httptest.StatusOK).
			ContentType(u.contentType(), app.ConfigurationReadOnly().GetCharset()).
			Body().Equal(contents)
	}
}

File: file-server/single-page-application/embedded-single-page-application/public/app.js

window.alert("app.js loaded from \"/");

File: file-server/single-page-application/embedded-single-page-application/public/css/main.css

body {
    background-color: black;
}

File: file-server/single-page-application/embedded-single-page-application/public/index.html

<html>

<head>
    <title>{{ .Page.Title }}</title>
</head>

<body>
    <h1> Hello from index.html </h1>


    <script src="/app.js">  </script>
</body>

</html>

File: file-server/single-page-application/embedded-single-page-application-with-other-routes/main.go

package main

import "github.com/kataras/iris"

// $ go get -u github.com/shuLhan/go-bindata/...
// $ go-bindata ./public/...
// $ go build
// $ ./embedded-single-page-application-with-other-routes

func newApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.Writef("404 not found here")
	})

	app.StaticEmbedded("/", "./public", Asset, AssetNames)

	// Note:
	// if you want a dynamic index page then see the file-server/embedded-single-page-application
	// which is registering a view engine based on bindata as well and a root route.

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	app.Get("/.well-known", func(ctx iris.Context) {
		ctx.WriteString("well-known")
	})
	app.Get(".well-known/ready", func(ctx iris.Context) {
		ctx.WriteString("ready")
	})
	app.Get(".well-known/live", func(ctx iris.Context) {
		ctx.WriteString("live")
	})
	app.Get(".well-known/metrics", func(ctx iris.Context) {
		ctx.Writef("metrics")
	})
	return app
}

func main() {
	app := newApp()

	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	//
	// http://localhost:8080/ping
	// http://localhost:8080/.well-known
	// http://localhost:8080/.well-known/ready
	// http://localhost:8080/.well-known/live
	// http://localhost:8080/.well-known/metrics
	//
	// Remember: we could use the root wildcard `app.Get("/{param:path}")` and serve the files manually as well.
	app.Run(iris.Addr(":8080"))
}

File: file-server/single-page-application/embedded-single-page-application-with-other-routes/public/app.js

window.alert("app.js loaded from \"/");

File: file-server/single-page-application/embedded-single-page-application-with-other-routes/public/css/main.css

body {
    background-color: black;
}

File: file-server/single-page-application/embedded-single-page-application-with-other-routes/public/index.html

<html>

<head>
    <title>Hello from static page</title>
</head>

<body>
    <h1> Hello from index.html </h1>


    <script src="/app.js">  </script>
</body>

</html>

Hero

Basic

File: hero/basic/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func main() {

	app := iris.New()

	// 1
	helloHandler := hero.Handler(hello)
	app.Get("/{to:string}", helloHandler)

	// 2
	hero.Register(&myTestService{
		prefix: "Service: Hello",
	})

	helloServiceHandler := hero.Handler(helloService)
	app.Get("/service/{to:string}", helloServiceHandler)

	// 3
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		// it binds the "form" with a
		// x-www-form-urlencoded form data and returns it.
		ctx.ReadForm(&form)
		return
	})

	loginHandler := hero.Handler(login)
	app.Post("/login", loginHandler)

	// http://localhost:8080/your_name
	// http://localhost:8080/service/your_name
	app.Run(iris.Addr(":8080"))
}

func hello(to string) string {
	return "Hello " + to
}

type Service interface {
	SayHello(to string) string
}

type myTestService struct {
	prefix string
}

func (s *myTestService) SayHello(to string) string {
	return s.prefix + " " + to
}

func helloService(to string, service Service) string {
	return service.SayHello(to)
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func login(form LoginForm) string {
	return "Hello " + form.Username
}

Overview

File: hero/overview/datamodels/movie.go

// file: datamodels/movie.go

package datamodels

// Movie is our sample data structure.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/movie.go"
// which could wrap by embedding the datamodels.Movie or
// declare new fields instead butwe will use this datamodel
// as the only one Movie model in our application,
// for the shake of simplicty.
type Movie struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

File: hero/overview/datasource/movies.go

// file: datasource/movies.go

package datasource

import "github.com/kataras/iris/_examples/hero/overview/datamodels"

// Movies is our imaginary data source.
var Movies = map[int64]datamodels.Movie{
	1: {
		ID:     1,
		Name:   "Casablanca",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	2: {
		ID:     2,
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	3: {
		ID:     3,
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	4: {
		ID:     4,
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
	5: {
		ID:     5,
		Name:   "North by Northwest",
		Year:   1959,
		Genre:  "Thriller",
		Poster: "https://iris-go.com/images/examples/mvc-movies/5.jpg",
	},
}

File: hero/overview/main.go

// file: main.go

package main

import (
	"github.com/kataras/iris/_examples/hero/overview/datasource"
	"github.com/kataras/iris/_examples/hero/overview/repositories"
	"github.com/kataras/iris/_examples/hero/overview/services"
	"github.com/kataras/iris/_examples/hero/overview/web/middleware"
	"github.com/kataras/iris/_examples/hero/overview/web/routes"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	movieService := services.NewMovieService(repo)
	hero.Register(movieService)

	// Serve our routes with hero handlers.
	app.PartyFunc("/hello", func(r iris.Party) {
		r.Get("/", hero.Handler(routes.Hello))
		r.Get("/{name}", hero.Handler(routes.HelloName))
	})

	app.PartyFunc("/movies", func(r iris.Party) {
		// Add the basic authentication(admin:password) middleware
		// for the /movies based requests.
		r.Use(middleware.BasicAuth)

		r.Get("/", hero.Handler(routes.Movies))
		r.Get("/{id:long}", hero.Handler(routes.MovieByID))
		r.Put("/{id:long}", hero.Handler(routes.UpdateMovieByID))
		r.Delete("/{id:long}", hero.Handler(routes.DeleteMovieByID))
	})

	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

File: hero/overview/repositories/movie_repository.go

// file: repositories/movie_repository.go

package repositories

import (
	"errors"
	"sync"

	"github.com/kataras/iris/_examples/hero/overview/datamodels"
)

// Query represents the visitor and action queries.
type Query func(datamodels.Movie) bool

// MovieRepository handles the basic operations of a movie entity/model.
// It's an interface in order to be testable, i.e a memory movie repository or
// a connected to an sql database.
type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewMovieRepository returns a new movie memory-based repository,
// the one and only repository type in our example.
func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

// movieMemoryRepository is a "MovieRepository"
// which manages the movies using the memory data source (map).
type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

// Select receives a query function
// which is fired for every single movie model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known "found" value
// and the last known movie model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exec(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty datamodels.Movie if not found at all.
	if !found {
		movie = datamodels.Movie{}
	}

	return
}

// SelectMany same as Select but returns one or more datamodels.Movie as a slice.
// If limit <=0 then it returns everything.
func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Movie) {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate adds or updates a movie to the (memory) storage.
//
// Returns the new movie and an error if any.
func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) (datamodels.Movie, error) {
	id := movie.ID

	if id == 0 { // Create new action
		var lastID int64
		// find the biggest ID in order to not have duplications
		// in productions apps you can use a third-party
		// library to generate a UUID as string.
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		movie.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = movie
		r.mu.Unlock()

		return movie, nil
	}

	// Update action based on the movie.ID,
	// here we will allow updating the poster and genre if not empty.
	// Alternatively we could do pure replace instead:
	// r.source[id] = movie
	// and comment the code below;
	current, exists := r.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})

	if !exists { // ID is not a real one, return an error.
		return datamodels.Movie{}, errors.New("failed to update a nonexistent movie")
	}

	// or comment these and r.source[id] = m for pure replace
	if movie.Poster != "" {
		current.Poster = movie.Poster
	}

	if movie.Genre != "" {
		current.Genre = movie.Genre
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return movie, nil
}

func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Movie) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}

File: hero/overview/services/movie_service.go

// file: services/movie_service.go

package services

import (
	"github.com/kataras/iris/_examples/hero/overview/datamodels"
	"github.com/kataras/iris/_examples/hero/overview/repositories"
)

// MovieService handles some of the CRUID operations of the movie datamodel.
// It depends on a movie repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type MovieService interface {
	GetAll() []datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error)
}

// NewMovieService returns the default movie service.
func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo repositories.MovieRepository
}

// GetAll returns all movies.
func (s *movieService) GetAll() []datamodels.Movie {
	return s.repo.SelectMany(func(_ datamodels.Movie) bool {
		return true
	}, -1)
}

// GetByID returns a movie based on its id.
func (s *movieService) GetByID(id int64) (datamodels.Movie, bool) {
	return s.repo.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID updates a movie's poster and genre.
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

// DeleteByID deletes a movie by its id.
//
// Returns true if deleted otherwise false.
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	}, 1)
}

File: hero/overview/web/middleware/basicauth.go

// file: web/middleware/basicauth.go

package middleware

import "github.com/kataras/iris/middleware/basicauth"

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

File: hero/overview/web/routes/hello.go

// file: web/routes/hello.go

package routes

import (
	"errors"

	"github.com/kataras/iris/hero"
)

var helloView = hero.View{
	Name: "hello/index.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

// Hello will return a predefined view with bind data.
//
// `hero.Result` is just an interface with a `Dispatch` function.
// `hero.Response` and `hero.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func Hello() hero.Result {
	return helloView
}

// you can define a standard error in order to re-use anywhere in your app.
var errBadName = errors.New("bad name")

// you can just return it as error or even better
// wrap this error with an hero.Response to make it an hero.Result compatible type.
var badName = hero.Response{Err: errBadName, Code: 400}

// HelloName returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/hello/iris
// curl -i http://localhost:8080/hello/anything
func HelloName(name string) hero.Result {
	if name != "iris" {
		return badName
	}

	// return hero.Response{Text: "Hello " + name} OR:
	return hero.View{
		Name: "hello/name.html",
		Data: name,
	}
}

File: hero/overview/web/routes/movies.go

// file: web/routes/movie.go

package routes

import (
	"errors"

	"github.com/kataras/iris/_examples/hero/overview/datamodels"
	"github.com/kataras/iris/_examples/hero/overview/services"

	"github.com/kataras/iris"
)

// Movies returns list of the movies.
// Demo:
// curl -i http://localhost:8080/movies
func Movies(service services.MovieService) (results []datamodels.Movie) {
	return service.GetAll()
}

// MovieByID returns a movie.
// Demo:
// curl -i http://localhost:8080/movies/1
func MovieByID(service services.MovieService, id int64) (movie datamodels.Movie, found bool) {
	return service.GetByID(id) // it will throw 404 if not found.
}

// UpdateMovieByID updates a movie.
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func UpdateMovieByID(ctx iris.Context, service services.MovieService, id int64) (datamodels.Movie, error) {
	// get the request data for poster and genre
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Movie{}, errors.New("failed due form file 'poster' missing")
	}
	// we don't need the file so close it now.
	file.Close()

	// imagine that is the url of the uploaded file...
	poster := info.Filename
	genre := ctx.FormValue("genre")

	return service.UpdatePosterAndGenreByID(id, poster, genre)
}

// DeleteMovieByID deletes a movie.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func DeleteMovieByID(service services.MovieService, id int64) interface{} {
	wasDel := service.DeleteByID(id)
	if wasDel {
		// return the deleted movie's ID
		return iris.Map{"deleted": id}
	}
	// right here we can see that a method function can return any of those two types(map or int),
	// we don't have to specify the return type to a specific type.
	return iris.StatusBadRequest
}

File: hero/overview/web/views/hello/index.html

<!-- file: web/views/hello/index.html -->
<html>

<head>
    <title>{{.Title}} - My App</title>
</head>

<body>
    <p>{{.MyMessage}}</p>
</body>

</html>

File: hero/overview/web/views/hello/name.html

<!-- file: web/views/hello/name.html -->
<html>

<head>
    <title>{{.}}' Portfolio - My App</title>
</head>

<body>
    <h1>Hello {{.}}</h1>
</body>

</html>

HTTP Listening

Custom HTTPserver

File: http-listening/custom-httpserver/easy-way/main.go

package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// Any custom fields here. Handler and ErrorLog are setted to the server automatically
	srv := &http.Server{Addr: ":8080"}

	// http://localhost:8080/
	// http://localhost:8080/mypath
	app.Run(iris.Server(srv)) // same as app.Run(iris.Addr(":8080"))

	// More:
	// see "multi" if you need to use more than one server at the same app.
	//
	// for a custom listener use: iris.Listener(net.Listener) or
	// iris.TLS(cert,key) or iris.AutoTLS(), see "custom-listener" example for those.
}

File: http-listening/custom-httpserver/multi/main.go

package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// Note: It's not needed if the first action is "go app.Run".
	if err := app.Build(); err != nil {
		panic(err)
	}

	// start a secondary server listening on localhost:9090.
	// use "go" keyword for Listen functions if you need to use more than one server at the same app.
	//
	// http://localhost:9090/
	// http://localhost:9090/mypath
	srv1 := &http.Server{Addr: ":9090", Handler: app}
	go srv1.ListenAndServe()
	println("Start a server listening on http://localhost:9090")

	// start a "second-secondary" server listening on localhost:5050.
	//
	// http://localhost:5050/
	// http://localhost:5050/mypath
	srv2 := &http.Server{Addr: ":5050", Handler: app}
	go srv2.ListenAndServe()
	println("Start a server listening on http://localhost:5050")

	// Note: app.Run is totally optional, we have already built the app with app.Build,
	// you can just make a new http.Server instead.
	// http://localhost:8080/
	// http://localhost:8080/mypath
	app.Run(iris.Addr(":8080")) // Block here.
}

File: http-listening/custom-httpserver/std-way/main.go

package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// call .Build before use the 'app' as a http.Handler on a custom http.Server
	app.Build()

	// create our custom server and assign the Handler/Router
	srv := &http.Server{Handler: app, Addr: ":8080"} // you have to set Handler:app and Addr, see "iris-way" which does this automatically.
	// http://localhost:8080/
	// http://localhost:8080/mypath
	println("Start a server listening on http://localhost:8080")
	srv.ListenAndServe() // same as app.Run(iris.Addr(":8080"))

	// Notes:
	// Banner is not shown at all. Same for the Interrupt Handler, even if app's configuration allows them.
	//
	// `.Run` is the only one function that cares about those three.

	// More:
	// see "multi" if you need to use more than one server at the same app.
	//
	// for a custom listener use: iris.Listener(net.Listener) or
	// iris.TLS(cert,key) or iris.AutoTLS(), see "custom-listener" example for those.
}

Custom Listener

File: http-listening/custom-listener/main.go

package main

import (
	"net"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// create any custom tcp listener, unix sock file or tls tcp listener.
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	// use of the custom listener
	app.Run(iris.Listener(l))
}

File: http-listening/custom-listener/unix-reuseport/main.go

// +build linux darwin dragonfly freebsd netbsd openbsd rumprun

package main

import (
	// Package tcplisten provides customizable TCP net.Listener with various
	// performance-related options:
	//
	//   - SO_REUSEPORT. This option allows linear scaling server performance
	//     on multi-CPU servers.
	//     See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for details.
	//
	//   - TCP_DEFER_ACCEPT. This option expects the server reads from the accepted
	//     connection before writing to them.
	//
	//   - TCP_FASTOPEN. See https://lwn.net/Articles/508865/ for details.
	"github.com/valyala/tcplisten"

	"github.com/kataras/iris"
)

// $ go get github.com/valyala/tcplisten
// $ go run main.go

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello World!</b>")
	})

	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}

	l, err := listenerCfg.NewListener("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	app.Run(iris.Listener(l))
}

File: http-listening/custom-listener/unix-reuseport/main_windows.go

// +build windows

package main

func main() {
	panic("windows operating system does not support this feature")
}

Graceful Shutdown

File: http-listening/graceful-shutdown/custom-notifier/main.go

package main

import (
	stdContext "context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hi, I just exist in order to see if the server is closed</h1>")
	})

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")

			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()

	// Start the server and disable the default interrupt handler in order to
	// handle it clear and simple by our own, without any issues.
	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}

File: http-listening/graceful-shutdown/default-notifier/main.go

package main

import (
	stdContext "context"
	"time"

	"github.com/kataras/iris"
)

// Before continue:
//
// Gracefully Shutdown on control+C/command+C or when kill command sent is ENABLED BY-DEFAULT.
//
// In order to manually manage what to do when app is interrupted,
// We have to disable the default behavior with the option `WithoutInterruptHandler`
// and register a new interrupt handler (globally, across all possible hosts).
func main() {
	app := iris.New()

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}

Iris Configurator And Host Configurator

File: http-listening/iris-configurator-and-host-configurator/counter/configurator.go

package counter

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
)

func Configurator(app *iris.Application) {
	counterValue := 0

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			counterValue++
		}

		app.ConfigureHost(func(h *host.Supervisor) { // <- HERE: IMPORTANT
			h.RegisterOnShutdown(func() {
				ticker.Stop()
			})
		}) // or put the ticker outside of the gofunc and put the configurator before or after the app.Get, outside of this gofunc
	}()

	app.Get("/counter", func(ctx iris.Context) {
		ctx.Writef("Counter value = %d", counterValue)
	})
}

File: http-listening/iris-configurator-and-host-configurator/main.go

package main

import (
	"github.com/kataras/iris/_examples/http-listening/iris-configurator-and-host-configurator/counter"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Configure(counter.Configurator)

	app.Run(iris.Addr(":8080"))
}

Listen Addr

File: http-listening/listen-addr/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

File: http-listening/listen-addr/omit-server-errors/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		// do something
	}
	// same as:
	// err := app.Run(iris.Addr(":8080"))
	// if err != nil && (err != iris.ErrServerClosed || err.Error() != iris.ErrServerClosed.Error()) {
	//     [...]
	// }
}

File: http-listening/listen-addr/omit-server-errors/main_test.go

package main

import (
	"bytes"
	stdContext "context"
	"strings"
	"testing"
	"time"

	"github.com/kataras/iris"
)

func logger(app *iris.Application) *bytes.Buffer {
	buf := &bytes.Buffer{}

	app.Logger().SetOutput(buf)

	// disable the "Now running at...." in order to have a clean log of the error.
	// we could attach that on `Run` but better to keep things simple here.
	app.Configure(iris.WithoutStartupLog)
	return buf
}

func TestListenAddr(t *testing.T) {
	app := iris.New()
	// we keep the logger running as well but in a controlled way.
	log := logger(app)

	// close the server at 3-6 seconds
	go func() {
		time.Sleep(3 * time.Second)
		ctx, cancel := stdContext.WithTimeout(stdContext.TODO(), 3*time.Second)
		defer cancel()
		app.Shutdown(ctx)
	}()

	err := app.Run(iris.Addr(":9829"))
	// in this case the error should be logged and return as well.
	if err != iris.ErrServerClosed {
		t.Fatalf("expecting err to be `iris.ErrServerClosed` but got: %v", err)
	}

	expectedMessage := iris.ErrServerClosed.Error()

	if got := log.String(); !strings.Contains(got, expectedMessage) {
		t.Fatalf("expecting to log to contains the:\n'%s'\ninstead of:\n'%s'", expectedMessage, got)
	}

}

func TestListenAddrWithoutServerErr(t *testing.T) {
	app := iris.New()
	// we keep the logger running as well but in a controlled way.
	log := logger(app)

	// close the server at 3-6 seconds
	go func() {
		time.Sleep(3 * time.Second)
		ctx, cancel := stdContext.WithTimeout(stdContext.TODO(), 3*time.Second)
		defer cancel()
		app.Shutdown(ctx)
	}()

	// we disable the ErrServerClosed, so the error should be nil when server is closed by `app.Shutdown`.

	// so in this case the iris/http.ErrServerClosed should be NOT logged and NOT return.
	err := app.Run(iris.Addr(":9827"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		t.Fatalf("expecting err to be nil but got: %v", err)
	}

	if got := log.String(); got != "" {
		t.Fatalf("expecting to log nothing but logged: '%s'", got)
	}
}

Listen Letsencrypt

File: http-listening/listen-letsencrypt/main.go

// Package main provide one-line integration with letsencrypt.org
package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from SECURE SERVER!")
	})

	app.Get("/test2", func(ctx iris.Context) {
		ctx.Writef("Welcome to secure server from /test2!")
	})

	app.Get("/redirect", func(ctx iris.Context) {
		ctx.Redirect("/test2")
	})

	// NOTE: This will not work on domains like this,
	// use real whitelisted domain(or domains split by whitespaces)
	// and a non-public e-mail instead.
	app.Run(iris.AutoTLS(":443", "example.com", "mail@example.com"))
}

Listen Tls

File: http-listening/listen-tls/main.go

package main

import (
	"net/url"

	"github.com/kataras/iris"

	"github.com/kataras/iris/core/host"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the SECURE server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})

	// to start a new server listening at :80 and redirects
	// to the secure address, then:
	target, _ := url.Parse("https://127.0.1:443")
	go host.NewProxy("127.0.0.1:80", target).ListenAndServe()

	// start the server (HTTPS) on port 443, this is a blocking func
	app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))

}

File: http-listening/listen-tls/mycert.cert

-----BEGIN CERTIFICATE-----

MIIDAzCCAeugAwIBAgIJAPDsxtKV4v3uMA0GCSqGSIb3DQEBBQUAMBgxFjAUBgNV
BAMMDTEyNy4wLjAuMTo0NDMwHhcNMTYwNjI5MTMxMjU4WhcNMjYwNjI3MTMxMjU4
WjAYMRYwFAYDVQQDDA0xMjcuMC4wLjE6NDQzMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA0KtAOHKrcbLwWJXgRX7XSFyu4HHHpSty4bliv8ET4sLJpbZH
XeVX05Foex7PnrurDP6e+0H5TgqqcpQM17/ZlFcyKrJcHSCgV0ZDB3Sb8RLQSLns
8a+MOSbn1WZ7TkC7d/cWlKmasQRHQ2V/cWlGooyKNEPoGaEz8MbY0wn2spyIJwsB
dciERC6317VTXbiZdoD8QbAsT+tBvEHM2m2A7B7PQmHNehtyFNbSV5uZNodvv1uv
ZTnDa6IqpjFLb1b2HNFgwmaVPmmkLuy1l9PN+o6/DUnXKKBrfPAx4JOlqTKEQpWs
pnfacTE3sWkkmOSSFltAXfkXIJFKdS/hy5J/KQIDAQABo1AwTjAdBgNVHQ4EFgQU
zr1df/c9+NyTpmyiQO8g3a8NswYwHwYDVR0jBBgwFoAUzr1df/c9+NyTpmyiQO8g
3a8NswYwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEACG5shtMSDgCd
MNjOF+YmD+PX3Wy9J9zehgaDJ1K1oDvBbQFl7EOJl8lRMWITSws22Wxwh8UXVibL
sscKBp14dR3e7DdbwCVIX/JyrJyOaCfy2nNBdf1B06jYFsIHvP3vtBAb9bPNOTBQ
QE0Ztu9kCqgsmu0//sHuBEeA3d3E7wvDhlqRSxTLcLtgC1NSgkFvBw0JvwgpkX6s
M5WpSBZwZv8qpplxhFfqNy8Uf+xrpSW0pGfkHumehkQGC6/Ry7raganS0aHhDPK9
Z1bEJ2com1bFFAQsm9yIXrRVMGGCtihB2Au0Q4jpEjUbzWYM+ItZyvRAGRM6Qex6
s/jogMeRsw==

-----END CERTIFICATE-----

File: http-listening/listen-tls/mykey.key

-----BEGIN RSA PRIVATE KEY-----

MIIEpQIBAAKCAQEA0KtAOHKrcbLwWJXgRX7XSFyu4HHHpSty4bliv8ET4sLJpbZH
XeVX05Foex7PnrurDP6e+0H5TgqqcpQM17/ZlFcyKrJcHSCgV0ZDB3Sb8RLQSLns
8a+MOSbn1WZ7TkC7d/cWlKmasQRHQ2V/cWlGooyKNEPoGaEz8MbY0wn2spyIJwsB
dciERC6317VTXbiZdoD8QbAsT+tBvEHM2m2A7B7PQmHNehtyFNbSV5uZNodvv1uv
ZTnDa6IqpjFLb1b2HNFgwmaVPmmkLuy1l9PN+o6/DUnXKKBrfPAx4JOlqTKEQpWs
pnfacTE3sWkkmOSSFltAXfkXIJFKdS/hy5J/KQIDAQABAoIBAQDCd+bo9I0s8Fun
4z3Y5oYSDTZ5O/CY0O5GyXPrSzCSM4Cj7EWEj1mTdb9Ohv9tam7WNHHLrcd+4NfK
4ok5hLVs1vqM6h6IksB7taKATz+Jo0PzkzrsXvMqzERhEBo4aoGMIv2rXIkrEdas
S+pCsp8+nAWtAeBMCn0Slu65d16vQxwgfod6YZfvMKbvfhOIOShl9ejQ+JxVZcMw
Ti8sgvYmFUrdrEH3nCgptARwbx4QwlHGaw/cLGHdepfFsVaNQsEzc7m61fSO70m4
NYJv48ZgjOooF5AccbEcQW9IxxikwNc+wpFYy5vDGzrBwS5zLZQFpoyMWFhtWdjx
hbmNn1jlAoGBAPs0ZjqsfDrH5ja4dQIdu5ErOccsmoHfMMBftMRqNG5neQXEmoLc
Uz8WeQ/QDf302aTua6E9iSjd7gglbFukVwMndQ1Q8Rwxz10jkXfiE32lFnqK0csx
ltruU6hOeSGSJhtGWBuNrT93G2lmy23fSG6BqOzdU4rn/2GPXy5zaxM/AoGBANSm
/E96RcBUiI6rDVqKhY+7M1yjLB41JrErL9a0Qfa6kYnaXMr84pOqVN11IjhNNTgl
g1lwxlpXZcZh7rYu9b7EEMdiWrJDQV7OxLDHopqUWkQ+3MHwqs6CxchyCq7kv9Df
IKqat7Me6Cyeo0MqcW+UMxlCRBxKQ9jqC7hDfZuXAoGBAJmyS8ImerP0TtS4M08i
JfsCOY21qqs/hbKOXCm42W+be56d1fOvHngBJf0YzRbO0sNo5Q14ew04DEWLsCq5
+EsDv0hwd7VKfJd+BakV99ruQTyk5wutwaEeJK1bph12MD6L4aiqHJAyLeFldZ45
+TUzu8mA+XaJz+U/NXtUPvU9AoGBALtl9M+tdy6I0Fa50ujJTe5eEGNAwK5WNKTI
5D2XWNIvk/Yh4shXlux+nI8UnHV1RMMX++qkAYi3oE71GsKeG55jdk3fFQInVsJQ
APGw3FDRD8M4ip62ki+u+tEr/tIlcAyHtWfjNKO7RuubWVDlZFXqCiXmSdOMdsH/
bxiREW49AoGACWev/eOzBoQJCRN6EvU2OV0s3b6f1QsPvcaH0zc6bgbBFOGmJU8v
pXhD88tsu9exptLkGVoYZjR0n0QT/2Kkyu93jVDW/80P7VCz8DKYyAJDa4CVwZxO
MlobQSunSDKx/CCJhWkbytCyh1bngAtwSAYLXavYIlJbAzx6FvtAIw4=

-----END RSA PRIVATE KEY-----

Listen Unix

File: http-listening/listen-unix/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/netutil"
)

func main() {
	app := iris.New()

	l, err := netutil.UNIX("/tmpl/srv.sock", 0666) // see its code to see how you can manually create a new file listener, it's easy.
	if err != nil {
		panic(err)
	}

	app.Run(iris.Listener(l))
}

// Look "custom-listener/unix-reuseport" too.

Notify On Shutdown

File: http-listening/notify-on-shutdown/main.go

package main

import (
	stdContext "context"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello, try to refresh the page after ~10 secs</h1>")
	})

	// app.ConfigureHost(configureHost) -> or pass "configureHost" as `app.Addr` argument, same result.

	app.Logger().Info("Wait 10 seconds and check your terminal again")
	// simulate a shutdown action here...
	go func() {
		<-time.After(10 * time.Second)
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts, this will notify the callback we had register
		// inside the `configureHost` func.
		app.Shutdown(ctx)
	}()

	// start the server as usual, the only difference is that
	// we're adding a second (optional) function
	// to configure the just-created host supervisor.
	//
	// http://localhost:8080
	// wait 10 seconds and check your terminal.
	app.Run(iris.Addr(":8080", configureHost), iris.WithoutServerError(iris.ErrServerClosed))

}

func configureHost(su *host.Supervisor) {
	// here we have full access to the host that will be created
	// inside the `app.Run` function or `NewHost`.
	//
	// we're registering a shutdown "event" callback here:
	su.RegisterOnShutdown(func() {
		println("server is closed")
	})
	// su.RegisterOnError
	// su.RegisterOnServe
}

HTTP Request

Read Custom Per Type

File: http_request/read-custom-per-type/main.go

package main

import (
	"gopkg.in/yaml.v2"

	"github.com/kataras/iris"
)

func main() {
	app := newApp()

	// use Postman or whatever to do a POST request
	// (however you are always free to use app.Get and GET http method requests to read body of course)
	// to the http://localhost:8080 with RAW BODY:
	/*
		addr: localhost:8080
		serverName: Iris
	*/
	//
	// The response should be:
	// Received: main.config{Addr:"localhost:8080", ServerName:"Iris"}
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func newApp() *iris.Application {
	app := iris.New()
	app.Post("/", handler)

	return app
}

// simple yaml stuff, read more at https://github.com/go-yaml/yaml
type config struct {
	Addr       string `yaml:"addr"`
	ServerName string `yaml:"serverName"`
}

// Decode implements the `kataras/iris/context#BodyDecoder` optional interface
// that any go type can implement in order to be self-decoded when reading the request's body.
func (c *config) Decode(body []byte) error {
	return yaml.Unmarshal(body, c)
}

func handler(ctx iris.Context) {
	var c config

	//
	// Note:
	// second parameter is nil because our &c implements the `context#BodyDecoder`
	// which has a priority over the context#Unmarshaler (which can be a more global option for reading request's body)
	// see the `http_request/read-custom-via-unmarshaler/main.go` example to learn how to use the context#Unmarshaler too.
	//
	// Note 2:
	// If you need to read the body again for any reason
	// you should disable the body consumption via `app.Run(..., iris.WithoutBodyConsumptionOnUnmarshal)`.
	//

	if err := ctx.UnmarshalBody(&c, nil); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v", c)
}

File: http_request/read-custom-per-type/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestReadCustomPerType(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	expectedResponse := `Received: main.config{Addr:"localhost:8080", ServerName:"Iris"}`

	e.POST("/").WithText("addr: localhost:8080\nserverName: Iris").Expect().
		Status(httptest.StatusOK).Body().Equal(expectedResponse)
}

Read Custom Via Unmarshaler

File: http_request/read-custom-via-unmarshaler/main.go

package main

import (
	"gopkg.in/yaml.v2"

	"github.com/kataras/iris"
)

func main() {
	app := newApp()

	// use Postman or whatever to do a POST request
	// (however you are always free to use app.Get and GET http method requests to read body of course)
	// to the http://localhost:8080 with RAW BODY:
	/*
		addr: localhost:8080
		serverName: Iris
	*/
	//
	// The response should be:
	// Received: main.config{Addr:"localhost:8080", ServerName:"Iris"}
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func newApp() *iris.Application {
	app := iris.New()
	app.Post("/", handler)

	return app
}

// simple yaml stuff, read more at https://github.com/go-yaml/yaml
type config struct {
	Addr       string `yaml:"addr"`
	ServerName string `yaml:"serverName"`
}

/*
type myBodyDecoder struct{}

var DefaultBodyDecoder = myBodyDecoder{}

// Implements the `kataras/iris/context#Unmarshaler` but at our example
// we will use the simplest `context#UnmarshalerFunc` to pass just the yaml.Unmarshal.
//
// Can be used as: ctx.UnmarshalBody(&c, DefaultBodyDecoder)
func (r *myBodyDecoder) Unmarshal(data []byte, outPtr interface{}) error {
	return yaml.Unmarshal(data, outPtr)
}
*/

func handler(ctx iris.Context) {
	var c config

	//
	// Note:
	// yaml.Unmarshal already implements the `context#Unmarshaler`
	// so we can use it directly, like the json.Unmarshal(ctx.ReadJSON), xml.Unmarshal(ctx.ReadXML)
	// and every library which follows the best practises and is aligned with the Go standards.
	//
	// Note 2:
	// If you need to read the body again for any reason
	// you should disable the body consumption via `app.Run(..., iris.WithoutBodyConsumptionOnUnmarshal)`.
	//

	if err := ctx.UnmarshalBody(&c, iris.UnmarshalerFunc(yaml.Unmarshal)); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v", c)
}

File: http_request/read-custom-via-unmarshaler/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestReadCustomViaUnmarshaler(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	expectedResponse := `Received: main.config{Addr:"localhost:8080", ServerName:"Iris"}`

	e.POST("/").WithText("addr: localhost:8080\nserverName: Iris").Expect().
		Status(httptest.StatusOK).Body().Equal(expectedResponse)
}

Read Form

File: http_request/read-form/main.go

// package main contains an example on how to use the ReadForm, but with the same way you can do the ReadJSON & ReadJSON
package main

import (
	"github.com/kataras/iris"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {
	app := iris.New()

	// set the view html template engine
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})

	app.Post("/form_action", func(ctx iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}

		ctx.Writef("Visitor: %#v", visitor)
	})

	app.Post("/post_value", func(ctx iris.Context) {
		username := ctx.PostValueDefault("Username", "iris")
		ctx.Writef("Username: %s", username)
	})

	app.Run(iris.Addr(":8080"))
}

File: http_request/read-form/templates/form.html

<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
	<form action="/form_action" method="post">
		Username: <input type="text" name="Username" /> <br />
		Mail: <input type="text" name="Mail" /> <br />
		Select one or more:  <br/>
		<select multiple="multiple" name="mydata">
			<option value='one'>One</option>
			<option value='two'>Two</option>
			<option value='three'>Three</option>
			<option value='four'>Four</option>
		</select>

		<hr />
		<input type="submit" value="Send data" />

	</form>
</body>
</html>

Read Json

File: http_request/read-json/main.go

package main

import (
	"github.com/kataras/iris"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx iris.Context) {
	var c Company

	if err := ctx.ReadJSON(&c); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", c)
}

// simple json stuff, read more at https://golang.org/pkg/encoding/json
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// MyHandler2 reads a collection of Person from JSON post body.
func MyHandler2(ctx iris.Context) {
	var persons []Person
	err := ctx.ReadJSON(&persons)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", persons)
}

func main() {
	app := iris.New()

	app.Post("/", MyHandler)
	app.Post("/slice", MyHandler2)

	// use Postman or whatever to do a POST request
	// to the http://localhost:8080 with RAW BODY:
	/*
		{
			"Name": "iris-Go",
			"City": "New York",
			"Other": "Something here"
		}
	*/
	// and Content-Type to application/json (optionally but good practise)
	//
	// The response should be:
	// Received: main.Company{Name:"iris-Go", City:"New York", Other:"Something here"}
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

Read Xml

File: http_request/read-xml/main.go

package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
)

func main() {
	app := newApp()

	// use Postman or whatever to do a POST request
	// to the http://localhost:8080 with RAW BODY:
	/*
		<person name="Winston Churchill" age="90">
			<description>Description of this person, the body of this inner element.</description>
		</person>
	*/
	// and Content-Type to application/xml (optionally but good practise)
	//
	// The response should be:
	// Received: main.person{XMLName:xml.Name{Space:"", Local:"person"}, Name:"Winston Churchill", Age:90, Description:"Description of this person, the body of this inner element."}
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func newApp() *iris.Application {
	app := iris.New()
	app.Post("/", handler)

	return app
}

// simple xml stuff, read more at https://golang.org/pkg/encoding/xml
type person struct {
	XMLName     xml.Name `xml:"person"`      // element name
	Name        string   `xml:"name,attr"`   // ,attr for attribute.
	Age         int      `xml:"age,attr"`    // ,attr attribute.
	Description string   `xml:"description"` // inner element name, value is its body.
}

func handler(ctx iris.Context) {
	var p person
	if err := ctx.ReadXML(&p); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v", p)
}

File: http_request/read-xml/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestReadXML(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	expectedResponse := `Received: main.person{XMLName:xml.Name{Space:"", Local:"person"}, Name:"Winston Churchill", Age:90, Description:"Description of this person, the body of this inner element."}`
	send := `<person name="Winston Churchill" age="90"><description>Description of this person, the body of this inner element.</description></person>`

	e.POST("/").WithText(send).Expect().
		Status(httptest.StatusOK).Body().Equal(expectedResponse)
}

Request Logger

File: http_request/request-logger/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	app := iris.New()

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	h := func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	}
	app.Get("/", h)

	app.Get("/1", h)

	app.Get("/2", h)

	// http errors have their own handlers, therefore
	// registering a middleare should be done manually.
	/*
	 app.OnErrorCode(404 ,customLogger, func(ctx iris.Context) {
	 	ctx.Writef("My Custom 404 error page ")
	 })
	*/
	// or catch all http errors:
	app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
		// this should be added to the logs, at the end because of the `logger.Config#MessageContextKey`
		ctx.Values().Set("logger_message",
			"a dynamic message passed to the logs")
		ctx.Writef("My Custom error page")
	})

	// http://localhost:8080
	// http://localhost:8080/1
	// http://localhost:8080/2
	// http://lcoalhost:8080/notfoundhere
	// see the output on the console.
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}

File: http_request/request-logger/request-logger-file/main.go

package main

import (
	"os"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

const deleteFileOnExit = true

func main() {
	app := iris.New()
	r, close := newRequestLogger()
	defer close()

	app.Use(r)
	app.OnAnyErrorCode(r, func(ctx iris.Context) {
		ctx.HTML("<h1> Error: Please try <a href ='/'> this </a> instead.</h1>")
	})

	h := func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	}

	app.Get("/", h)

	app.Get("/1", h)

	app.Get("/2", h)

	// http://localhost:8080
	// http://localhost:8080/1
	// http://localhost:8080/2
	// http://lcoalhost:8080/notfoundhere
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func newRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }

	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	logFile := newLogFile()
	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}

	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}

	//	we don't want to use the logger
	// to log requests to assets and etc
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})

	h = logger.New(c)

	return
}

Upload File

File: http_request/upload-file/main.go

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

const maxSize = 5 << 20 // 5MB

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html"))

	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).

		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// render the form with the token for any use you'd like.
		// ctx.ViewData("", token)
		// or add second argument to the `View` method.
		// Token will be passed as {{.}} in the template.
		ctx.View("upload_form.html", token)
	})

	/* Read before continue.

	0. The default post max size is 32MB,
	you can extend it to read more data using the `iris.WithPostMaxMemory(maxSize)` configurator at `app.Run`,
	note that this will not be enough for your needs, read below.

	1. The faster way to check the size is using the `ctx.GetContentLength()` which returns the whole request's size
	(plus a logical number like 2MB or even 10MB for the rest of the size like headers). You can create a
	middleware to adapt this to any necessary handler.

	myLimiter := func(ctx iris.Context) {
		if ctx.GetContentLength() > maxSize { // + 2 << 20 {
			ctx.StatusCode(iris.StatusRequestEntityTooLarge)
			return
		}
		ctx.Next()
	}

	app.Post("/upload", myLimiter, myUploadHandler)

	Most clients will set the "Content-Length" header (like browsers) but it's always better to make sure that any client
	can't send data that your server can't or doesn't want to handle. This can be happen using
	the `app.Use(LimitRequestBodySize(maxSize))` (as app or route middleware)
	or the `ctx.SetMaxRequestBodySize(maxSize)` to limit the request based on a customized logic inside a particular handler, they're the same,
	read below.

	2. You can force-limit the request body size inside a handler using the `ctx.SetMaxRequestBodySize(maxSize)`,
	this will force the connection to close if the incoming data are larger (most clients will receive it as "connection reset"),
	use that to make sure that the client will not send data that your server can't or doesn't want to accept, as a fallback.

	app.Post("/upload", iris.LimitRequestBodySize(maxSize), myUploadHandler)

	OR

	app.Post("/upload", func(ctx iris.Context){
		ctx.SetMaxRequestBodySize(maxSize)

		// [...]
	})

	3. Another way is to receive the data and check the second return value's `Size` value of the `ctx.FormFile`, i.e `info.Size`, this will give you
	the exact file size, not the whole incoming request data length.

	app.Post("/", func(ctx iris.Context){
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()

		if info.Size > maxSize {
			ctx.StatusCode(iris.StatusRequestEntityTooLarge)
			return
		}

		// [...]
	})
	*/

	// Handle the post request from the upload_form.html to the server
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		// Get the file from the request.
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile("./uploads/"+fname,
			os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()

		io.Copy(out, file)
	})

	// start the server at http://localhost:8080 with post limit at 5 MB.
	app.Run(iris.Addr(":8080") /* 0.*/, iris.WithPostMaxMemory(maxSize))
}

File: http_request/upload-file/templates/upload_form.html

<html>

<head>
	<title>Upload file</title>
</head>

<body>
	<form enctype="multipart/form-data" action="http://127.0.0.1:8080/upload" method="POST">
		<input type="file" name="uploadfile" />
		<input type="hidden" name="token" value="{{.}}" />
		<input type="submit" value="upload" />
	</form>
</body>

</html>

Upload Files

File: http_request/upload-files/main.go

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html"))

	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).

		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// render the form with the token for any use you'd like.
		ctx.View("upload_form.html", token)
	})

	// Handle the post request from the upload_form.html to the server.
	app.Post("/upload", func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files (multiple property on the form input).
		//

		// second argument is totally optionally,
		// it can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		ctx.UploadFormFiles("./uploads", beforeSave)
	})

	// start the server at http://localhost:8080 with post limit at 32 MB.
	app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(32<<20))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}

File: http_request/upload-files/templates/upload_form.html

<html>
<head>
<title>Upload file</title>
</head>
<body>
	<form enctype="multipart/form-data"
		action="http://127.0.0.1:8080/upload" method="POST">
		<input type="file" name="uploadfile" multiple/> <input type="hidden"
			name="token" value="{{.}}" /> <input type="submit" value="upload" />
	</form>
</body>
</html>

HTTP Responsewriter

Herotemplate

File: http_responsewriter/herotemplate/app.go

package main

import (
	"bytes"

	"github.com/kataras/iris/_examples/http_responsewriter/herotemplate/template"

	"github.com/kataras/iris"
)

// $ go get -u github.com/shiyanhui/hero/hero
// $ go run app.go
//
// Read more at https://github.com/shiyanhui/hero/hero

func main() {

	app := iris.New()

	app.Get("/users", func(ctx iris.Context) {
		ctx.Gzip(true)
		ctx.ContentType("text/html")

		var userList = []string{
			"Alice",
			"Bob",
			"Tom",
		}

		// Had better use buffer sync.Pool.
		// Hero(github.com/shiyanhui/hero/hero) exports GetBuffer and PutBuffer for this.
		//
		// buffer := hero.GetBuffer()
		// defer hero.PutBuffer(buffer)
		// buffer := new(bytes.Buffer)
		// template.UserList(userList, buffer)
		// ctx.Write(buffer.Bytes())

		// using an io.Writer for automatic buffer management (i.e. hero built-in buffer pool),
		// iris context implements the io.Writer by its ResponseWriter
		// which is an enhanced version of the standard http.ResponseWriter
		// but still 100% compatible, GzipResponseWriter too:
		// _, err := template.UserListToWriter(userList, ctx.GzipResponseWriter())
		buffer := new(bytes.Buffer)
		template.UserList(userList, buffer)

		_, err := ctx.Write(buffer.Bytes())
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})

	app.Run(iris.Addr(":8080"))
}

File: http_responsewriter/herotemplate/template/index.html

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>

    <body>
        <%@ body { %>
        <% } %>
    </body>
</html>

File: http_responsewriter/herotemplate/template/index.html.go

// Code generated by hero.
// DO NOT EDIT!
package template

File: http_responsewriter/herotemplate/template/user.html

<li>
    <%= user %>
</li>

File: http_responsewriter/herotemplate/template/user.html.go

// Code generated by hero.
// DO NOT EDIT!
package template

File: http_responsewriter/herotemplate/template/userlist.html

<%: func UserList(userList []string, buffer *bytes.Buffer) %>

<%~ "index.html" %>

<%@ body { %>
    <% for _, user := range userList { %>
        <ul>
            <%+ "user.html" %>
        </ul>
    <% } %>
<% } %>

File: http_responsewriter/herotemplate/template/userlist.html.go

// Code generated by hero.
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func UserList(userList []string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>

    <body>
        `)
	for _, user := range userList {
		buffer.WriteString(`
        <ul>
            `)
		buffer.WriteString(`<li>
    `)
		hero.EscapeHTML(user, buffer)
		buffer.WriteString(`
</li>
`)

		buffer.WriteString(`
        </ul>
    `)
	}

	buffer.WriteString(`
    </body>
</html>
`)

}

File: http_responsewriter/herotemplate/template/userlistwriter.html

<%: func UserListToWriter(userList []string, w io.Writer) (int, error)%>

<%~ "index.html" %>

<%@ body { %>
    <% for _, user := range userList { %>
        <ul>
            <%+ "user.html" %>
        </ul>
    <% } %>
<% } %>

File: http_responsewriter/herotemplate/template/userlistwriter.html.go

// Code generated by hero.
// DO NOT EDIT!
package template

import (
	"io"

	"github.com/shiyanhui/hero"
)

func UserListToWriter(userList []string, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>

    <body>
        `)
	for _, user := range userList {
		_buffer.WriteString(`
        <ul>
            `)
		_buffer.WriteString(`<li>
    `)
		hero.EscapeHTML(user, _buffer)
		_buffer.WriteString(`
</li>
`)

		_buffer.WriteString(`
        </ul>
    `)
	}

	_buffer.WriteString(`
    </body>
</html>
`)
	return w.Write(_buffer.Bytes())

}

Quicktemplate

File: http_responsewriter/quicktemplate/controllers/execute_template.go

package controllers

import (
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/kataras/iris"
)

// ExecuteTemplate renders a "tmpl" partial template to the `context#ResponseWriter`.
func ExecuteTemplate(ctx iris.Context, tmpl templates.Partial) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	templates.WriteTemplate(ctx.ResponseWriter(), tmpl)
}

File: http_responsewriter/quicktemplate/controllers/hello.go

package controllers

import (
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/kataras/iris"
)

// Hello renders our ../templates/hello.qtpl file using the compiled ../templates/hello.qtpl.go file.
func Hello(ctx iris.Context) {
	// vars := make(map[string]interface{})
	// vars["message"] = "Hello World!"
	// vars["name"] = ctx.Params().Get("name")
	// [...]
	// &templates.Hello{ Vars: vars }
	// [...]

	// However, as an alternative, we recommend that you should the `ctx.ViewData(key, value)`
	// in order to be able modify the `templates.Hello#Vars` from a middleware(other handlers) as well.
	ctx.ViewData("message", "Hello World!")
	ctx.ViewData("name", ctx.Params().Get("name"))

	// set view data to the `Vars` template's field
	tmpl := &templates.Hello{
		Vars: ctx.GetViewData(),
	}

	// render the template
	ExecuteTemplate(ctx, tmpl)
}

File: http_responsewriter/quicktemplate/controllers/index.go

package controllers

import (
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/kataras/iris"
)

// Index renders our ../templates/index.qtpl file using the compiled ../templates/index.qtpl.go file.
func Index(ctx iris.Context) {
	tmpl := &templates.Index{}

	// render the template
	ExecuteTemplate(ctx, tmpl)
}

File: http_responsewriter/quicktemplate/main.go

package main

import (
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/controllers"

	"github.com/kataras/iris"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/", controllers.Index)
	app.Get("/{name}", controllers.Hello)

	return app
}

func main() {
	app := newApp()
	// http://localhost:8080
	// http://localhost:8080/yourname
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

File: http_responsewriter/quicktemplate/main_test.go

package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestResponseWriterQuicktemplate(t *testing.T) {
	baseRawBody := `
<html>
	<head>
		<title>Quicktemplate integration with Iris</title>
	</head>
	<body>
		<div>
			Header contents here...
		</div>

		<div style="margin:10px;">

	<h1>%s</h1>
	<div>
		%s
	</div>

		</div>

	</body>
	<footer>
		Footer contents here...
	</footer>
</html>
`

	expectedIndexRawBody := fmt.Sprintf(baseRawBody, "Index Page", "This is our index page's body.")
	name := "yourname"
	expectedHelloRawBody := fmt.Sprintf(baseRawBody, "Hello World!", "Hello <b>"+name+"!</b>")

	app := newApp()

	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal(expectedIndexRawBody)
	e.GET("/" + name).Expect().Status(httptest.StatusOK).Body().Equal(expectedHelloRawBody)
}

File: http_responsewriter/quicktemplate/templates/base.qtpl

This is our templates' base implementation.

{% interface
Partial {
	Body()
}
%}


Template writes a template implementing the Partial interface.
{% func Template(p Partial) %}
<html>
	<head>
		<title>Quicktemplate integration with Iris</title>
	</head>
	<body>
		<div>
			Header contents here...
		</div>

		<div style="margin:10px;">
			{%= p.Body() %}
		</div>

	</body>
	<footer>
		Footer contents here...
	</footer>
</html>
{% endfunc %}


Base template implementation. Other pages may inherit from it if they need
overriding only certain Partial methods.
{% code type Base struct {} %}
{% func (b *Base) Body() %}This is the base body{% endfunc %}

File: http_responsewriter/quicktemplate/templates/hello.qtpl

Hello template, implements the Partial's methods.

{% code
type Hello struct {
  Vars map[string]interface{}
}
%}

{% func (h *Hello) Body() %}
	<h1>{%v h.Vars["message"] %}</h1>
	<div>
		Hello <b>{%v h.Vars["name"] %}!</b>
	</div>
{% endfunc %}

File: http_responsewriter/quicktemplate/templates/index.qtpl

Index template, implements the Partial's methods.

{% code
type Index struct {}
%}

{% func (i *Index) Body() %}
	<h1>Index Page</h1>
	<div>
		This is our index page's body.
	</div>
{% endfunc %}

Sse Third Party

File: http_responsewriter/sse-third-party/main.go

package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/r3labs/sse"
)

// First of all install the sse third-party package (you can use other if you don't like this approach)
// $ go get -u github.com/r3labs/sse
func main() {
	app := iris.New()
	s := sse.New()
	/*
		This creates a new stream inside of the scheduler.
		Seeing as there are no consumers, publishing a message
		to this channel will do nothing.
		Clients can connect to this stream once the iris handler is started
		by specifying stream as a url parameter, like so:
		http://localhost:8080/events?stream=messages
	*/
	s.CreateStream("messages")

	app.Any("/events", iris.FromStd(s.HTTPHandler))

	go func() {
		// You design when to send messages to the client,
		// here we just wait 5 seconds to send the first message
		// in order to give u time to open a browser window...
		time.Sleep(5 * time.Second)
		// Publish a payload to the stream.
		s.Publish("messages", &sse.Event{
			Data: []byte("ping"),
		})

		time.Sleep(3 * time.Second)
		s.Publish("messages", &sse.Event{
			Data: []byte("second message"),
		})
		time.Sleep(2 * time.Second)
		s.Publish("messages", &sse.Event{
			Data: []byte("third message"),
		})

	}() // ...

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

/* For a golang SSE client you can look at: https://github.com/r3labs/sse#example-client */

Stream Writer

File: http_responsewriter/stream-writer/main.go

package main

import (
	"fmt" // just an optional helper
	"io"
	"time" // showcase the delay

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.ContentType("text/html")
		ctx.Header("Transfer-Encoding", "chunked")
		i := 0
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		// Send the response in chunks and wait for half a second between each chunk.
		ctx.StreamWriter(func(w io.Writer) bool {
			fmt.Fprintf(w, "Message number %d<br>", ints[i])
			time.Sleep(500 * time.Millisecond) // simulate delay.
			if i == len(ints)-1 {
				return false // close and flush
			}
			i++
			return true // continue write
		})
	})

	type messageNumber struct {
		Number int `json:"number"`
	}

	app.Get("/alternative", func(ctx iris.Context) {
		ctx.ContentType("application/json")
		ctx.Header("Transfer-Encoding", "chunked")
		i := 0
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		// Send the response in chunks and wait for half a second between each chunk.
		for {
			ctx.JSON(messageNumber{Number: ints[i]})
			ctx.WriteString("\n")
			time.Sleep(500 * time.Millisecond) // simulate delay.
			if i == len(ints)-1 {
				break
			}
			i++
			ctx.ResponseWriter().Flush()
		}
	})

	app.Run(iris.Addr(":8080"))
}

Transactions

File: http_responsewriter/transactions/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	// subdomains works with all available routers, like other features too.

	app.Get("/", func(ctx context.Context) {
		ctx.BeginTransaction(func(t *context.Transaction) {
			// OPTIONAl STEP: , if true then the next transictions will not be executed if this transiction fails
			// t.SetScope(context.RequestTransactionScope)

			// OPTIONAL STEP:
			// create a new custom type of error here to keep track of the status code and reason message
			err := context.NewTransactionErrResult()

			// we should use t.Context if we want to rollback on any errors lives inside this function clojure.
			t.Context().Text("Blablabla this should not be sent to the client because we will fill the err with a message and status")

			// virtualize a fake error here, for the shake of the example
			fail := true
			if fail {
				err.StatusCode = iris.StatusInternalServerError
				// NOTE: if empty reason then the default or the custom http error will be fired (like ctx.FireStatusCode)
				err.Reason = "Error: Virtual failure!!"
			}

			// OPTIONAl STEP:
			// but useful if we want to post back an error message to the client if the transaction failed.
			// if the reason is empty then the transaction completed successfully,
			// otherwise we rollback the whole response writer's body,
			// headers and cookies, status code and everything lives inside this transaction
			t.Complete(err)
		})

		ctx.BeginTransaction(func(t *context.Transaction) {
			t.Context().HTML("<h1>This will sent at all cases because it lives on different transaction and it doesn't fails</h1>")
			// * if we don't have any 'throw error' logic then no need of scope.Complete()
		})

		// OPTIONALLY, depends on the usage:
		// at any case, what ever happens inside the context's transactions send this to the client
		ctx.HTML("<h1>Let's add a second html message to the response, " +
			"if the transaction was failed and it was request scoped then this message would " +
			"not been shown. But it has a transient scope(default) so, it is visible as expected!</h1>")
	})

	app.Run(iris.Addr(":8080"))
}

Write Gzip

File: http_responsewriter/write-gzip/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteGzip([]byte("Hello World!"))
		ctx.Header("X-Custom",
			"Headers can be set here after WriteGzip as well, because the data are kept before sent to the client when using the context's GzipResponseWriter and ResponseRecorder.")
	})

	app.Get("/2", func(ctx iris.Context) {
		// same as the `WriteGzip`.
		// However GzipResponseWriter gives you more options, like
		// reset data, disable and more, look its methods.
		ctx.GzipResponseWriter().WriteString("Hello World!")
	})

	app.Run(iris.Addr(":8080"))
}

Write Rest

File: http_responsewriter/write-rest/main.go

package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// User bind struct
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

// ExampleXML just a test struct to view represents xml content-type
type ExampleXML struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

func main() {
	app := iris.New()

	// Read
	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)

		ctx.Writef("%s %s is %d years old and comes from %s!", user.Firstname, user.Lastname, user.Age, user.City)
	})

	// Write
	app.Get("/encode", func(ctx iris.Context) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}

		// Manually setting a content type: ctx.ContentType("application/javascript")
		ctx.JSON(peter)
	})

	// Other content types,

	app.Get("/binary", func(ctx iris.Context) {
		// useful when you want force-download of contents of raw bytes form.
		ctx.Binary([]byte("Some binary data here."))
	})

	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Plain text here")
	})

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
	})

	app.Get("/jsonp", func(ctx iris.Context) {
		ctx.JSONP(map[string]string{"hello": "jsonp"}, context.JSONP{Callback: "callbackName"})
	})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(ExampleXML{One: "hello", Two: "xml"}) // or iris.Map{"One":"hello"...}
	})

	app.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
	})

	// http://localhost:8080/decode
	// http://localhost:8080/encode
	//
	// http://localhost:8080/binary
	// http://localhost:8080/text
	// http://localhost:8080/json
	// http://localhost:8080/jsonp
	// http://localhost:8080/xml
	// http://localhost:8080/markdown
	//
	// `iris.WithOptimizations` is an optional configurator,
	// if passed to the `Run` then it will ensure that the application
	// response to the client as fast as possible.
	//
	//
	// `iris.WithoutServerError` is an optional configurator,
	// if passed to the `Run` then it will not print its passed error as an actual server error.
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

Miscellaneous

File Logger

File: miscellaneous/file-logger/main.go

package main

import (
	"os"
	"time"

	"github.com/kataras/iris"
)

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func main() {
	f := newLogFile()
	defer f.Close()

	app := iris.New()
	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(newLogFile())

	app.Get("/", func(ctx iris.Context) {
		// for the sake of simplicity, in order see the logs at the ./_today_.txt
		ctx.Application().Logger().Info("Request path: " + ctx.Path())
		ctx.Writef("hello")
	})

	// navigate to http://localhost:8080
	// and open the ./logs.txt file
	if err := app.Run(iris.Addr(":8080"), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}

I18n

File: miscellaneous/i18n/locales/locale_el-GR.ini

hi = ?e?a, %s

File: miscellaneous/i18n/locales/locale_en-US.ini

hi = hello, %s

File: miscellaneous/i18n/locales/locale_multi_first_el-GR.ini

key1 = a?t? e??a? ??a t??? ap? t? p??t? a??e??: locale_multi_first

File: miscellaneous/i18n/locales/locale_multi_first_en-US.ini

key1 = this is a value from the first file: locale_multi_first

File: miscellaneous/i18n/locales/locale_multi_second_el-GR.ini

key2 = a?t? e??a? ??a t??? ap? t? de?te?? a??e?? ?et?f?as??: locale_multi_second

File: miscellaneous/i18n/locales/locale_multi_second_en-US.ini

key2 = this is a value from the second file: locale_multi_second

File: miscellaneous/i18n/locales/locale_zh-CN.ini

hi = ??,%s

File: miscellaneous/i18n/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
)

func newApp() *iris.Application {
	app := iris.New()

	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"el-GR": "./locales/locale_el-GR.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	app.Get("/", func(ctx iris.Context) {

		// it tries to find the language by:
		// ctx.Values().GetString("language")
		// if that was empty then
		// it tries to find from the URLParameter setted on the configuration
		// if not found then
		// it tries to find the language by the "language" cookie
		// if didn't found then it it set to the Default setted on the configuration

		// hi is the key, 'iris' is the %s on the .ini file
		// the second parameter is optional

		// hi := ctx.Translate("hi", "iris")
		// or:
		hi := i18n.Translate(ctx, "hi", "iris")

		language := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())
		// return is form of 'en-US'

		// The first succeed language found saved at the cookie with name ("language"),
		//  you can change that by changing the value of the:  iris.TranslateLanguageContextKey
		ctx.Writef("From the language %s translated output: %s", language, hi)
	})

	multiLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_multi_first_en-US.ini, ./locales/locale_multi_second_en-US.ini",
			"el-GR": "./locales/locale_multi_first_el-GR.ini, ./locales/locale_multi_second_el-GR.ini"}})

	app.Get("/multi", multiLocale, func(ctx iris.Context) {
		language := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())

		fromFirstFileValue := i18n.Translate(ctx, "key1")
		fromSecondFileValue := i18n.Translate(ctx, "key2")
		ctx.Writef("From the language: %s, translated output:\n%s=%s\n%s=%s",
			language, "key1", fromFirstFileValue,
			"key2", fromSecondFileValue)
	})

	return app
}

func main() {
	app := newApp()

	// go to http://localhost:8080/?lang=el-GR
	// or http://localhost:8080 (default is en-US)
	// or http://localhost:8080/?lang=zh-CN
	//
	// go to http://localhost:8080/multi?lang=el-GR
	// or http://localhost:8080/multi (default is en-US)
	// or http://localhost:8080/multi?lang=en-US
	//
	// or use cookies to set the language.
	app.Run(iris.Addr(":8080"))
}

File: miscellaneous/i18n/main_test.go

package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestI18n(t *testing.T) {
	app := newApp()

	expectedf := "From the language %s translated output: %s"
	var (
		elgr = fmt.Sprintf(expectedf, "el-GR", "?e?a, iris")
		enus = fmt.Sprintf(expectedf, "en-US", "hello, iris")
		zhcn = fmt.Sprintf(expectedf, "zh-CN", "??,iris")

		elgrMulti = fmt.Sprintf("From the language: %s, translated output:\n%s=%s\n%s=%s", "el-GR",
			"key1",
			"a?t? e??a? ??a t??? ap? t? p??t? a??e??: locale_multi_first",
			"key2",
			"a?t? e??a? ??a t??? ap? t? de?te?? a??e?? ?et?f?as??: locale_multi_second")
		enusMulti = fmt.Sprintf("From the language: %s, translated output:\n%s=%s\n%s=%s", "en-US",
			"key1",
			"this is a value from the first file: locale_multi_first",
			"key2",
			"this is a value from the second file: locale_multi_second")
	)

	e := httptest.New(t, app)
	// default is en-US
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal(enus)
	// default is en-US if lang query unable to be found
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal(enus)

	e.GET("/").WithQueryString("lang=el-GR").Expect().Status(httptest.StatusOK).
		Body().Equal(elgr)
	e.GET("/").WithQueryString("lang=en-US").Expect().Status(httptest.StatusOK).
		Body().Equal(enus)
	e.GET("/").WithQueryString("lang=zh-CN").Expect().Status(httptest.StatusOK).
		Body().Equal(zhcn)

	e.GET("/multi").WithQueryString("lang=el-GR").Expect().Status(httptest.StatusOK).
		Body().Equal(elgrMulti)
	e.GET("/multi").WithQueryString("lang=en-US").Expect().Status(httptest.StatusOK).
		Body().Equal(enusMulti)

}

Pprof

File: miscellaneous/pprof/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/pprof"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Any("/debug/pprof/{action:path}", pprof.New())
	//                              ___________
	app.Run(iris.Addr(":8080"))
}

Recaptcha

File: miscellaneous/recaptcha/custom_form/main.go

package main

import (
	"fmt"

	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/recaptcha"
)

// keys should be obtained by https://www.google.com/recaptcha
const (
	recaptchaPublic = "6Lf3WywUAAAAAKNfAm5DP2J5ahqedtZdHTYaKkJ6"
	recaptchaSecret = "6Lf3WywUAAAAAJpArb8nW_LCL_PuPuokmEABFfgw"
)

func main() {
	app := iris.New()

	r := recaptcha.New(recaptchaSecret)

	app.Get("/comment", showRecaptchaForm)

	// pass the middleware before the main handler or use the `recaptcha.SiteVerify`.
	app.Post("/comment", r, postComment)

	app.Run(iris.Addr(":8080"))
}

var htmlForm = `<form action="/comment" method="POST">
	    <script src="https://www.google.com/recaptcha/api.js"></script>
		<div class="g-recaptcha" data-sitekey="%s"></div>
    	<input type="submit" name="button" value="Verify">
</form>`

func showRecaptchaForm(ctx iris.Context) {
	contents := fmt.Sprintf(htmlForm, recaptchaPublic)
	ctx.HTML(contents)
}

func postComment(ctx iris.Context) {
	// [...]
	ctx.JSON(iris.Map{"success": true})
}

File: miscellaneous/recaptcha/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recaptcha"
)

// keys should be obtained by https://www.google.com/recaptcha
const (
	recaptchaPublic = ""
	recaptchaSecret = ""
)

func showRecaptchaForm(ctx iris.Context, path string) {
	ctx.HTML(recaptcha.GetFormHTML(recaptchaPublic, path))
}

func main() {
	app := iris.New()

	// On both Get and Post on this example, so you can easly
	// use a single route to show a form and the main subject if recaptcha's validation result succeed.
	app.HandleMany("GET POST", "/", func(ctx iris.Context) {
		if ctx.Method() == iris.MethodGet {
			showRecaptchaForm(ctx, "/")
			return
		}

		result := recaptcha.SiteFerify(ctx, recaptchaSecret)
		if !result.Success {
			/* redirect here if u want or do nothing */
			ctx.HTML("<b> failed please try again </b>")
			return
		}

		ctx.Writef("succeed.")
	})

	app.Run(iris.Addr(":8080"))
}

Recover

File: miscellaneous/recover/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	// use this recover(y) middleware
	app.Use(recover.New())

	i := 0
	// let's simulate a panic every next request
	app.Get("/", func(ctx iris.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
		}
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	// http://localhost:8080, refresh it 5-6 times.
	app.Run(iris.Addr(":8080"))
}

// Note:
// app := iris.Default() instead of iris.New() makes use of the recovery middleware automatically.

MVC

Basic

File: mvc/basic/main.go

package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("/basic"), basicMVC)

	app.Run(iris.Addr(":8080"))
}

func basicMVC(app *mvc.Application) {
	// You can use normal middlewares at MVC apps of course.
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	// Register dependencies which will be binding to the controller(s),
	// can be either a function which accepts an iris.Context and returns a single value (dynamic binding)
	// or a static struct value (service).
	app.Register(
		sessions.New(sessions.Config{}).Start,
		&prefixedLogger{prefix: "DEV"},
	)

	// GET: http://localhost:8080/basic
	// GET: http://localhost:8080/basic/custom
	app.Handle(new(basicController))

	// All dependencies of the parent *mvc.Application
	// are cloned to this new child,
	// thefore it has access to the same session as well.
	// GET: http://localhost:8080/basic/sub
	app.Party("/sub").
		Handle(new(basicSubController))
}

// If controller's fields (or even its functions) expecting an interface
// but a struct value is binded then it will check
// if that struct value implements
// the interface and if true then it will add this to the
// available bindings, as expected, before the server ran of course,
// remember? Iris always uses the best possible way to reduce load
// on serving web resources.

type LoggerService interface {
	Log(string)
}

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}

type basicController struct {
	Logger LoggerService

	Session *sessions.Session
}

func (c *basicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/custom", "Custom")
}

func (c *basicController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}

func (c *basicController) Get() string {
	count := c.Session.Increment("count", 1)

	body := fmt.Sprintf("Hello from basicController\nTotal visits from you: %d", count)
	c.Logger.Log(body)
	return body
}

func (c *basicController) Custom() string {
	return "custom"
}

type basicSubController struct {
	Session *sessions.Session
}

func (c *basicSubController) Get() string {
	count := c.Session.GetIntDefault("count", 1)
	return fmt.Sprintf("Hello from basicSubController.\nRead-only visits count: %d", count)
}

Hello World

File: mvc/hello-world/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// This example is equivalent to the
// https://github.com/kataras/iris/blob/master/_examples/hello-world/main.go
//
// It seems that additional code you
// have to write doesn't worth it
// but remember that, this example
// does not make use of iris mvc features like
// the Model, Persistence or the View engine neither the Session,
// it's very simple for learning purposes,
// probably you'll never use such
// as simple controller anywhere in your app.
//
// The cost we have on this example for using MVC
// on the "/hello" path which serves JSON
// is ~2MB per 20MB throughput on my personal laptop,
// it's tolerated for the majority of the applications
// but you can choose
// what suits you best with Iris, low-level handlers: performance
// or high-level controllers: easier to maintain and smaller codebase on large applications.

// Of course you can put all these to main func, it's just a separate function
// for the main_test.go.
func newApp() *iris.Application {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Serve a controller based on the root Router, "/".
	mvc.New(app).Handle(new(ExampleController))
	return app
}

func main() {
	app := newApp()

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhost:8080/custom_path
	app.Run(iris.Addr(":8080"))
}

// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

// BeforeActivation called once, before the controller adapted to the main application
// and of course before the server ran.
// After version 9 you can also add custom routes for a specific controller's methods.
// Here you can register custom method's handlers
// use the standard router with `ca.Router` to do something that you can do without mvc as well,
// and add dependencies that will be binded to a controller's fields or method function's input arguments.
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)

	// or even add a global middleware based on this controller's router,
	// which in this example is the root "/":
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

// GetUserBy serves
// Method:   GET
// Resource: http://localhost:8080/user/{username:string}
// By is a reserved "keyword" to tell the framework that you're going to
// bind path parameters in the function's input arguments, and it also
// helps to have "Get" and "GetBy" in the same controller.
//
// func (c *ExampleController) GetUserBy(username string) mvc.Result {
// 	return mvc.View{
// 		Name: "user/username.html",
// 		Data: username,
// 	}
// }

/* Can use more than one, the factory will make sure
that the correct http methods are being registered for each route
for this controller, uncomment these if you want:

func (c *ExampleController) Post() {}
func (c *ExampleController) Put() {}
func (c *ExampleController) Delete() {}
func (c *ExampleController) Connect() {}
func (c *ExampleController) Head() {}
func (c *ExampleController) Patch() {}
func (c *ExampleController) Options() {}
func (c *ExampleController) Trace() {}
*/

/*
func (c *ExampleController) All() {}
//        OR
func (c *ExampleController) Any() {}



func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	// 1 -> the HTTP Method
	// 2 -> the route's path
	// 3 -> this controller's method name that should be handler for that route.
	b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
}

// After activation, all dependencies are set-ed - so read only access on them
// but still possible to add custom controller or simple standard handlers.
func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}

*/

File: mvc/hello-world/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestMVCHelloWorld(t *testing.T) {
	e := httptest.New(t, newApp())

	e.GET("/").Expect().Status(httptest.StatusOK).
		ContentType("text/html", "utf-8").Body().Equal("<h1>Welcome</h1>")

	e.GET("/ping").Expect().Status(httptest.StatusOK).
		Body().Equal("pong")

	e.GET("/hello").Expect().Status(httptest.StatusOK).
		JSON().Object().Value("message").Equal("Hello Iris!")

	e.GET("/custom_path").Expect().Status(httptest.StatusOK).
		Body().Equal("hello from the custom handler without following the naming guide")
}

Login

File: mvc/login/datamodels/user.go

package datamodels

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the shake
// of the example, we will use this datamodel
// as the only one User model in our application.
type User struct {
	ID             int64     `json:"id" form:"id"`
	Firstname      string    `json:"firstname" form:"firstname"`
	Username       string    `json:"username" form:"username"`
	HashedPassword []byte    `json:"-" form:"-"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
}

// IsValid can do some very very simple "low-level" data validations.
func (u User) IsValid() bool {
	return u.ID > 0
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

File: mvc/login/datasource/users.go

// file: datasource/users.go

package datasource

import (
	"errors"

	"github.com/kataras/iris/_examples/mvc/login/datamodels"
)

// Engine is from where to fetch the data, in this case the users.
type Engine uint32

const (
	// Memory stands for simple memory location;
	// map[int64] datamodels.User ready to use, it's our source in this example.
	Memory Engine = iota
	// Bolt for boltdb source location.
	Bolt
	// MySQL for mysql-compatible source location.
	MySQL
)

// LoadUsers returns all users(empty map) from the memory, for the shake of simplicty.
func LoadUsers(engine Engine) (map[int64]datamodels.User, error) {
	if engine != Memory {
		return nil, errors.New("for the shake of simplicity we're using a simple map as the data source")
	}

	return make(map[int64]datamodels.User), nil
}

File: mvc/login/main.go

// file: main.go

package main

import (
	"time"

	"github.com/kataras/iris/_examples/mvc/login/datasource"
	"github.com/kataras/iris/_examples/mvc/login/repositories"
	"github.com/kataras/iris/_examples/mvc/login/services"
	"github.com/kataras/iris/_examples/mvc/login/web/controllers"
	"github.com/kataras/iris/_examples/mvc/login/web/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.StaticWeb("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	// ---- Serve our controllers. ----

	// Prepare our repositories and services.
	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		app.Logger().Fatalf("error while loading the users: %v", err)
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	// "/users" based mvc application.
	users := mvc.New(app.Party("/users"))
	// Add the basic authentication(admin:password) middleware
	// for the /users based requests.
	users.Router.Use(middleware.BasicAuth)
	// Bind the "userService" to the UserController's Service (interface) field.
	users.Register(userService)
	users.Handle(new(controllers.UsersController))

	// "/user" based mvc application.
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})
	user := mvc.New(app.Party("/user"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controllers.UserController))

	// http://localhost:8080/noexist
	// and all controller's methods like
	// http://localhost:8080/users/1
	// http://localhost:8080/user/register
	// http://localhost:8080/user/login
	// http://localhost:8080/user/me
	// http://localhost:8080/user/logout
	// basic auth: "admin", "password", see "./middleware/basicauth.go" source file.
	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Disables the updater.
		iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

File: mvc/login/repositories/user_repository.go

package repositories

import (
	"errors"
	"sync"

	"github.com/kataras/iris/_examples/mvc/login/datamodels"
)

// Query represents the visitor and action queries.
type Query func(datamodels.User) bool

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type UserRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (user datamodels.User, found bool)
	SelectMany(query Query, limit int) (results []datamodels.User)

	InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewUserRepository returns a new user memory-based repository,
// the one and only repository type in our example.
func NewUserRepository(source map[int64]datamodels.User) UserRepository {
	return &userMemoryRepository{source: source}
}

// userMemoryRepository is a "UserRepository"
// which manages the users using the memory data source (map).
type userMemoryRepository struct {
	source map[int64]datamodels.User
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *userMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.source {
		ok = query(user)
		if ok {
			if action(user) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

// Select receives a query function
// which is fired for every single user model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known boolean value
// and the last known user model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *userMemoryRepository) Select(query Query) (user datamodels.User, found bool) {
	found = r.Exec(query, func(m datamodels.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty datamodels.User if not found at all.
	if !found {
		user = datamodels.User{}
	}

	return
}

// SelectMany same as Select but returns one or more datamodels.User as a slice.
// If limit <=0 then it returns everything.
func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate adds or updates a user to the (memory) storage.
//
// Returns the new user and an error if any.
func (r *userMemoryRepository) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {
	id := user.ID

	if id == 0 { // Create new action
		var lastID int64
		// find the biggest ID in order to not have duplications
		// in productions apps you can use a third-party
		// library to generate a UUID as string.
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()

		return user, nil
	}

	// Update action based on the user.ID,
	// here we will allow updating the poster and genre if not empty.
	// Alternatively we could do pure replace instead:
	// r.source[id] = user
	// and comment the code below;
	current, exists := r.Select(func(m datamodels.User) bool {
		return m.ID == id
	})

	if !exists { // ID is not a real one, return an error.
		return datamodels.User{}, errors.New("failed to update a nonexistent user")
	}

	// or comment these and r.source[id] = user for pure replace
	if user.Username != "" {
		current.Username = user.Username
	}

	if user.Firstname != "" {
		current.Firstname = user.Firstname
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}

File: mvc/login/services/user_service.go

package services

import (
	"errors"

	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/repositories"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type UserService interface {
	GetAll() []datamodels.User
	GetByID(id int64) (datamodels.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (datamodels.User, bool)
	DeleteByID(id int64) bool

	Update(id int64, user datamodels.User) (datamodels.User, error)
	UpdatePassword(id int64, newPassword string) (datamodels.User, error)
	UpdateUsername(id int64, newUsername string) (datamodels.User, error)

	Create(userPassword string, user datamodels.User) (datamodels.User, error)
}

// NewUserService returns the default user service.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll returns all users.
func (s *userService) GetAll() []datamodels.User {
	return s.repo.SelectMany(func(_ datamodels.User) bool {
		return true
	}, -1)
}

// GetByID returns a user based on its id.
func (s *userService) GetByID(id int64) (datamodels.User, bool) {
	return s.repo.Select(func(m datamodels.User) bool {
		return m.ID == id
	})
}

// GetByUsernameAndPassword returns a user based on its username and passowrd,
// used for authentication.
func (s *userService) GetByUsernameAndPassword(username, userPassword string) (datamodels.User, bool) {
	if username == "" || userPassword == "" {
		return datamodels.User{}, false
	}

	return s.repo.Select(func(m datamodels.User) bool {
		if m.Username == username {
			hashed := m.HashedPassword
			if ok, _ := datamodels.ValidatePassword(userPassword, hashed); ok {
				return true
			}
		}
		return false
	})
}

// Update updates every field from an existing User,
// it's not safe to be used via public API,
// however we will use it on the web/controllers/user_controller.go#PutBy
// in order to show you how it works.
func (s *userService) Update(id int64, user datamodels.User) (datamodels.User, error) {
	user.ID = id
	return s.repo.InsertOrUpdate(user)
}

// UpdatePassword updates a user's password.
func (s *userService) UpdatePassword(id int64, newPassword string) (datamodels.User, error) {
	// update the user and return it.
	hashed, err := datamodels.GeneratePassword(newPassword)
	if err != nil {
		return datamodels.User{}, err
	}

	return s.Update(id, datamodels.User{
		HashedPassword: hashed,
	})
}

// UpdateUsername updates a user's username.
func (s *userService) UpdateUsername(id int64, newUsername string) (datamodels.User, error) {
	return s.Update(id, datamodels.User{
		Username: newUsername,
	})
}

// Create inserts a new User,
// the userPassword is the client-typed password
// it will be hashed before the insertion to our repository.
func (s *userService) Create(userPassword string, user datamodels.User) (datamodels.User, error) {
	if user.ID > 0 || userPassword == "" || user.Firstname == "" || user.Username == "" {
		return datamodels.User{}, errors.New("unable to create this user")
	}

	hashed, err := datamodels.GeneratePassword(userPassword)
	if err != nil {
		return datamodels.User{}, err
	}
	user.HashedPassword = hashed

	return s.repo.InsertOrUpdate(user)
}

// DeleteByID deletes a user by its id.
//
// Returns true if deleted otherwise false.
func (s *userService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.User) bool {
		return m.ID == id
	}, 1)
}

File: mvc/login/web/controllers/user_controller.go

// file: controllers/user_controller.go

package controllers

import (
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// All HTTP Methods /user/logout
type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UserController) logout() {
	c.Session.Destroy()
}

var registerStaticView = mvc.View{
	Name: "user/register.html",
	Data: iris.Map{"Title": "User Registration"},
}

// GetRegister handles GET: http://localhost:8080/user/register.
func (c *UserController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		c.logout()
	}

	return registerStaticView
}

// PostRegister handles POST: http://localhost:8080/user/register.
func (c *UserController) PostRegister() mvc.Result {
	// get firstname, username and password from the form.
	var (
		firstname = c.Ctx.FormValue("firstname")
		username  = c.Ctx.FormValue("username")
		password  = c.Ctx.FormValue("password")
	)

	// create the new user, the password will be hashed by the service.
	u, err := c.Service.Create(password, datamodels.User{
		Username:  username,
		Firstname: firstname,
	})

	// set the user's id to this session even if err != nil,
	// the zero id doesn't matters because .getCurrentUserID() checks for that.
	// If err != nil then it will be shown, see below on mvc.Response.Err: err.
	c.Session.Set(userIDKey, u.ID)

	return mvc.Response{
		// if not nil then this error will be shown instead.
		Err: err,
		// redirect to /user/me.
		Path: "/user/me",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}

}

var loginStaticView = mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"Title": "User Login"},
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		c.logout()
	}

	return loginStaticView
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (c *UserController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, found := c.Service.GetByUsernameAndPassword(username, password)

	if !found {
		return mvc.Response{
			Path: "/user/register",
		}
	}

	c.Session.Set(userIDKey, u.ID)

	return mvc.Response{
		Path: "/user/me",
	}
}

// GetMe handles GET: http://localhost:8080/user/me.
func (c *UserController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		// if it's not logged in then redirect user to the login page.
		return mvc.Response{Path: "/user/login"}
	}

	u, found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout and re-execute the function, it will redirect the client to the
		// /user/login page.
		c.logout()
		return c.GetMe()
	}

	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

// AnyLogout handles All/Any HTTP Methods for: http://localhost:8080/user/logout.
func (c *UserController) AnyLogout() {
	if c.isLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/user/login")
}

File: mvc/login/web/controllers/users_controller.go

package controllers

import (
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/services"

	"github.com/kataras/iris"
)

// UsersController is our /users API controller.
// GET				/users  | get all
// GET				/users/{id:long} | get by id
// PUT				/users/{id:long} | update by id
// DELETE			/users/{id:long} | delete by id
// Requires basic authentication.
type UsersController struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService
}

// Get returns list of the users.
// Demo:
// curl -i -u admin:password http://localhost:8080/users
//
// The correct way if you have sensitive data:
// func (c *UsersController) Get() (results []viewmodels.User) {
// 	data := c.Service.GetAll()
//
// 	for _, user := range data {
// 		results = append(results, viewmodels.User{user})
// 	}
// 	return
// }
// otherwise just return the datamodels.
func (c *UsersController) Get() (results []datamodels.User) {
	return c.Service.GetAll()
}

// GetBy returns a user.
// Demo:
// curl -i -u admin:password http://localhost:8080/users/1
func (c *UsersController) GetBy(id int64) (user datamodels.User, found bool) {
	u, found := c.Service.GetByID(id)
	if !found {
		// this message will be binded to the
		// main.go -> app.OnAnyErrorCode -> NotFound -> shared/error.html -> .Message text.
		c.Ctx.Values().Set("message", "User couldn't be found!")
	}
	return u, found // it will throw/emit 404 if found == false.
}

// PutBy updates a user.
// Demo:
// curl -i -X PUT -u admin:password -F "username=kataras"
// -F "password=rawPasswordIsNotSafeIfOrNotHTTPs_You_Should_Use_A_client_side_lib_for_hash_as_well"
// http://localhost:8080/users/1
func (c *UsersController) PutBy(id int64) (datamodels.User, error) {
	// username := c.Ctx.FormValue("username")
	// password := c.Ctx.FormValue("password")
	u := datamodels.User{}
	if err := c.Ctx.ReadForm(&u); err != nil {
		return u, err
	}

	return c.Service.Update(id, u)
}

// DeleteBy deletes a user.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/users/1
func (c *UsersController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		// return the deleted user's ID
		return map[string]interface{}{"deleted": id}
	}
	// right here we can see that a method function
	// can return any of those two types(map or int),
	// we don't have to specify the return type to a specific type.
	return iris.StatusBadRequest // same as 400.
}

File: mvc/login/web/middleware/basicauth.go

// file: middleware/basicauth.go

package middleware

import "github.com/kataras/iris/middleware/basicauth"

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

File: mvc/login/web/public/css/site.css

/* Bordered form */
form {
    border: 3px solid #f1f1f1;
}

/* Full-width inputs */
input[type=text], input[type=password] {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    box-sizing: border-box;
}

/* Set a style for all buttons */
button {
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    width: 100%;
}

/* Add a hover effect for buttons */
button:hover {
    opacity: 0.8;
}

/* Extra style for the cancel button (red) */
.cancelbtn {
    width: auto;
    padding: 10px 18px;
    background-color: #f44336;
}

/* Center the container */

/* Add padding to containers */
.container {
    padding: 16px;
}

/* The "Forgot password" text */
span.psw {
    float: right;
    padding-top: 16px;
}

/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
    span.psw {
        display: block;
        float: none;
    }
    .cancelbtn {
        width: 100%;
    }
}

File: mvc/login/web/views/shared/error.html

<h1>Error.</h1>
<h2>An error occurred while processing your request.</h2>

<h3>{{.Message}}</h3>

<footer>
    <h2>Sitemap</h2>
    <a href="http://localhost:8080/user/register">/user/register</a><br/>
    <a href="http://localhost:8080/user/login">/user/login</a><br/>
    <a href="http://localhost:8080/user/logout">/user/logout</a><br/>
    <a href="http://localhost:8080/user/me">/user/me</a><br/>
    <h3>requires authentication</h3><br/>
    <a href="http://localhost:8080/users">/users</a><br/>
    <a href="http://localhost:8080/users/1">/users/{id}</a><br/>
</footer>

File: mvc/login/web/views/shared/layout.html

<html>

<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css" />
</head>

<body>
    {{ yield }}
</body>

</html>

File: mvc/login/web/views/user/login.html

<form action="/user/login" method="POST">
    <div class="container">
        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="username" required>

        <label><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="password" required>

        <button type="submit">Login</button>
    </div>
</form>

File: mvc/login/web/views/user/me.html

<p>
    Welcome back <strong>{{.User.Firstname}}</strong>!
</p>

File: mvc/login/web/views/user/register.html

<form action="/user/register" method="POST">
    <div class="container">
        <label><b>Firstname</b></label>
        <input type="text" placeholder="Enter Firstname" name="firstname" required>

        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="username" required>

        <label><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="password" required>

        <button type="submit">Register</button>
    </div>
</form>

Middleware

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

Overview

File: mvc/overview/datamodels/movie.go

// file: datamodels/movie.go

package datamodels

// Movie is our sample data structure.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/movie.go"
// which could wrap by embedding the datamodels.Movie or
// declare new fields instead butwe will use this datamodel
// as the only one Movie model in our application,
// for the shake of simplicty.
type Movie struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

File: mvc/overview/datasource/movies.go

// file: datasource/movies.go

package datasource

import "github.com/kataras/iris/_examples/mvc/overview/datamodels"

// Movies is our imaginary data source.
var Movies = map[int64]datamodels.Movie{
	1: {
		ID:     1,
		Name:   "Casablanca",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	2: {
		ID:     2,
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	3: {
		ID:     3,
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	4: {
		ID:     4,
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
	5: {
		ID:     5,
		Name:   "North by Northwest",
		Year:   1959,
		Genre:  "Thriller",
		Poster: "https://iris-go.com/images/examples/mvc-movies/5.jpg",
	},
}

File: mvc/overview/main.go

// file: main.go

package main

import (
	"github.com/kataras/iris/_examples/mvc/overview/datasource"
	"github.com/kataras/iris/_examples/mvc/overview/repositories"
	"github.com/kataras/iris/_examples/mvc/overview/services"
	"github.com/kataras/iris/_examples/mvc/overview/web/controllers"
	"github.com/kataras/iris/_examples/mvc/overview/web/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Serve our controllers.
	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	mvc.Configure(app.Party("/movies"), movies)

	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// note the mvc.Application, it's not iris.Application.
func movies(app *mvc.Application) {
	// Add the basic authentication(admin:password) middleware
	// for the /movies based requests.
	app.Router.Use(middleware.BasicAuth)

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	movieService := services.NewMovieService(repo)
	app.Register(movieService)

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controllers.MovieController))
}

File: mvc/overview/repositories/movie_repository.go

// file: repositories/movie_repository.go

package repositories

import (
	"errors"
	"sync"

	"github.com/kataras/iris/_examples/mvc/overview/datamodels"
)

// Query represents the visitor and action queries.
type Query func(datamodels.Movie) bool

// MovieRepository handles the basic operations of a movie entity/model.
// It's an interface in order to be testable, i.e a memory movie repository or
// a connected to an sql database.
type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewMovieRepository returns a new movie memory-based repository,
// the one and only repository type in our example.
func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

// movieMemoryRepository is a "MovieRepository"
// which manages the movies using the memory data source (map).
type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

// Select receives a query function
// which is fired for every single movie model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known "found" value
// and the last known movie model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exec(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty datamodels.Movie if not found at all.
	if !found {
		movie = datamodels.Movie{}
	}

	return
}

// SelectMany same as Select but returns one or more datamodels.Movie as a slice.
// If limit <=0 then it returns everything.
func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Movie) {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate adds or updates a movie to the (memory) storage.
//
// Returns the new movie and an error if any.
func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) (datamodels.Movie, error) {
	id := movie.ID

	if id == 0 { // Create new action
		var lastID int64
		// find the biggest ID in order to not have duplications
		// in productions apps you can use a third-party
		// library to generate a UUID as string.
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		movie.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = movie
		r.mu.Unlock()

		return movie, nil
	}

	// Update action based on the movie.ID,
	// here we will allow updating the poster and genre if not empty.
	// Alternatively we could do pure replace instead:
	// r.source[id] = movie
	// and comment the code below;
	current, exists := r.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})

	if !exists { // ID is not a real one, return an error.
		return datamodels.Movie{}, errors.New("failed to update a nonexistent movie")
	}

	// or comment these and r.source[id] = m for pure replace
	if movie.Poster != "" {
		current.Poster = movie.Poster
	}

	if movie.Genre != "" {
		current.Genre = movie.Genre
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return movie, nil
}

func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Movie) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}

File: mvc/overview/services/movie_service.go

// file: services/movie_service.go

package services

import (
	"github.com/kataras/iris/_examples/mvc/overview/datamodels"
	"github.com/kataras/iris/_examples/mvc/overview/repositories"
)

// MovieService handles some of the CRUID operations of the movie datamodel.
// It depends on a movie repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type MovieService interface {
	GetAll() []datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error)
}

// NewMovieService returns the default movie service.
func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo repositories.MovieRepository
}

// GetAll returns all movies.
func (s *movieService) GetAll() []datamodels.Movie {
	return s.repo.SelectMany(func(_ datamodels.Movie) bool {
		return true
	}, -1)
}

// GetByID returns a movie based on its id.
func (s *movieService) GetByID(id int64) (datamodels.Movie, bool) {
	return s.repo.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID updates a movie's poster and genre.
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

// DeleteByID deletes a movie by its id.
//
// Returns true if deleted otherwise false.
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	}, 1)
}

File: mvc/overview/web/controllers/hello_controller.go

// file: web/controllers/hello_controller.go

package controllers

import (
	"errors"

	"github.com/kataras/iris/mvc"
)

// HelloController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type HelloController struct{}

var helloView = mvc.View{
	Name: "hello/index.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *HelloController) Get() mvc.Result {
	return helloView
}

// you can define a standard error in order to re-use anywhere in your app.
var errBadName = errors.New("bad name")

// you can just return it as error or even better
// wrap this error with an mvc.Response to make it an mvc.Result compatible type.
var badName = mvc.Response{Err: errBadName, Code: 400}

// GetBy returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/hello/iris
// curl -i http://localhost:8080/hello/anything
func (c *HelloController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
		// or
		// GetBy(name string) (mvc.Result, error) {
		//	return nil, errBadName
		// }
	}

	// return mvc.Response{Text: "Hello " + name} OR:
	return mvc.View{
		Name: "hello/name.html",
		Data: name,
	}
}

File: mvc/overview/web/controllers/movie_controller.go

// file: web/controllers/movie_controller.go

package controllers

import (
	"errors"

	"github.com/kataras/iris/_examples/mvc/overview/datamodels"
	"github.com/kataras/iris/_examples/mvc/overview/services"

	"github.com/kataras/iris"
)

// MovieController is our /movies controller.
type MovieController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	Service services.MovieService
}

// Get returns list of the movies.
// Demo:
// curl -i http://localhost:8080/movies
//
// The correct way if you have sensitive data:
// func (c *MovieController) Get() (results []viewmodels.Movie) {
// 	data := c.Service.GetAll()
//
// 	for _, movie := range data {
// 		results = append(results, viewmodels.Movie{movie})
// 	}
// 	return
// }
// otherwise just return the datamodels.
func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.Service.GetAll()
}

// GetBy returns a movie.
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}

// PutBy updates a movie.
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels.Movie, error) {
	// get the request data for poster and genre
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Movie{}, errors.New("failed due form file 'poster' missing")
	}
	// we don't need the file so close it now.
	file.Close()

	// imagine that is the url of the uploaded file...
	poster := info.Filename
	genre := ctx.FormValue("genre")

	return c.Service.UpdatePosterAndGenreByID(id, poster, genre)
}

// DeleteBy deletes a movie.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		// return the deleted movie's ID
		return iris.Map{"deleted": id}
	}
	// right here we can see that a method function can return any of those two types(map or int),
	// we don't have to specify the return type to a specific type.
	return iris.StatusBadRequest
}

File: mvc/overview/web/middleware/basicauth.go

// file: web/middleware/basicauth.go

package middleware

import "github.com/kataras/iris/middleware/basicauth"

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

File: mvc/overview/web/views/hello/index.html

<!-- file: web/views/hello/index.html -->
<html>

<head>
    <title>{{.Title}} - My App</title>
</head>

<body>
    <p>{{.MyMessage}}</p>
</body>

</html>

File: mvc/overview/web/views/hello/name.html

<!-- file: web/views/hello/name.html -->
<html>

<head>
    <title>{{.}}' Portfolio - My App</title>
</head>

<body>
    <h1>Hello {{.}}</h1>
</body>

</html>

Session Controller

File: mvc/session-controller/main.go

// +build go1.9

package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// VisitController handles the root route.
type VisitController struct {
	// the current request session,
	// its initialization happens by the dependency function that we've added to the `visitApp`.
	Session *sessions.Session

	// A time.time which is binded from the MVC,
	// order of binded fields doesn't matter.
	StartTime time.Time
}

// Get handles
// Method: GET
// Path: http://localhost:8080
func (c *VisitController) Get() string {
	// it increments a "visits" value of integer by one,
	// if the entry with key 'visits' doesn't exist it will create it for you.
	visits := c.Session.Increment("visits", 1)
	// write the current, updated visits.
	since := time.Now().Sub(c.StartTime).Seconds()
	return fmt.Sprintf("%d visit from my current session in %0.1f seconds of server's up-time",
		visits, since)
}

func newApp() *iris.Application {
	app := iris.New()
	sess := sessions.New(sessions.Config{Cookie: "mysession_cookie_name"})

	visitApp := mvc.New(app.Party("/"))
	// bind the current *session.Session, which is required, to the `VisitController.Session`
	// and the time.Now() to the `VisitController.StartTime`.
	visitApp.Register(
		// if dependency is a function which accepts
		// a Context and returns a single value
		// then the result type of this function is resolved by the controller
		// and on each request it will call the function with its Context
		// and set the result(the *sessions.Session here) to the controller's field.
		//
		// If dependencies are registered without field or function's input arguments as
		// consumers then those dependencies are being ignored before the server ran,
		// so you can bind many dependecies and use them in different controllers.
		sess.Start,
		time.Now(),
	)
	visitApp.Handle(new(VisitController))

	return app
}

func main() {
	app := newApp()

	// 1. open the browser (no in private mode)
	// 2. navigate to http://localhost:8080
	// 3. refresh the page some times
	// 4. close the browser
	// 5. re-open the browser and re-play 2.
	app.Run(iris.Addr(":8080"), iris.WithoutVersionChecker)
}

File: mvc/session-controller/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestMVCSession(t *testing.T) {
	e := httptest.New(t, newApp(), httptest.URL("http://example.com"))

	e1 := e.GET("/").Expect().Status(httptest.StatusOK)
	e1.Cookies().NotEmpty()
	e1.Body().Contains("1 visit")

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Contains("2 visit")

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Contains("3 visit")
}

Singleton

File: mvc/singleton/main.go

package main

import (
	"fmt"
	"sync/atomic"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&globalVisitorsController{visits: 0})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

type globalVisitorsController struct {
	// When a singleton controller is used then concurent safe access is up to the developers, because
	// all clients share the same controller instance instead.
	// Note that any controller's methods
	// are per-client, but the struct's field can be shared across multiple clients if the structure
	// does not have any dynamic struct field dependencies that depend on the iris.Context
	// and ALL field's values are NOT zero, at this case we use uint64 which it's no zero (even if we didn't set it
	// manually ease-of-understand reasons) because it's a value of &{0}.
	// All the above declares a Singleton, note that you don't have to write a single line of code to do this, Iris is smart enough.
	//
	// see `Get`.
	visits uint64
}

func (c *globalVisitorsController) Get() string {
	count := atomic.AddUint64(&c.visits, 1)
	return fmt.Sprintf("Total visitors: %d", count)
}

Websocket

File: mvc/websocket/main.go

package main

import (
	"sync/atomic"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

func main() {
	app := iris.New()
	// load templates.
	app.RegisterView(iris.HTML("./views", ".html"))

	// render the ./views/index.html.
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	mvc.Configure(app.Party("/websocket"), configureMVC)
	// Or mvc.New(app.Party(...)).Configure(configureMVC)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func configureMVC(m *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	// http://localhost:8080/websocket/iris-ws.js
	m.Router.Any("/iris-ws.js", websocket.ClientHandler())

	// This will bind the result of ws.Upgrade which is a websocket.Connection
	// to the controller(s) served by the `m.Handle`.
	m.Register(ws.Upgrade)
	m.Handle(new(websocketController))
}

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

type websocketController struct {
	// Note that you could use an anonymous field as well, it doesn't matter, binder will find it.
	//
	// This is the current websocket connection, each client has its own instance of the *websocketController.
	Conn websocket.Connection
}

func (c *websocketController) onLeave(roomName string) {
	// visits--
	newCount := decrement()
	// This will call the "visit" event on all clients, except the current one,
	// (it can't because it's left but for any case use this type of design)
	c.Conn.To(websocket.Broadcast).Emit("visit", newCount)
}

func (c *websocketController) update() {
	// visits++
	newCount := increment()

	// This will call the "visit" event on all clients, including the current
	// with the 'newCount' variable.
	//
	// There are many ways that u can do it and faster, for example u can just send a new visitor
	// and client can increment itself, but here we are just "showcasing" the websocket controller.
	c.Conn.To(websocket.All).Emit("visit", newCount)
}

func (c *websocketController) Get( /* websocket.Connection could be lived here as well, it doesn't matter */ ) {
	c.Conn.OnLeave(c.onLeave)
	c.Conn.On("visit", c.update)

	// call it after all event callbacks registration.
	c.Conn.Wait()
}

File: mvc/websocket/views/index.html

<html>

<head>
    <title>Online visitors MVC example</title>
    <style>
        body {
            margin: 0;
            font-family: -apple-system, "San Francisco", "Helvetica Neue", "Noto", "Roboto", "Calibri Light", sans-serif;
            color: #212121;
            font-size: 1.0em;
            line-height: 1.6;
        }

        .container {
            max-width: 750px;
            margin: auto;
            padding: 15px;
        }

        #online_visitors {
            font-weight: bold;
            font-size: 18px;
        }
    </style>
</head>

<body>
    <div class="container">
        <span id="online_visitors">1 online visitor</span>
    </div>

    <script src="/websocket/iris-ws.js"></script>

    <script type="text/javascript">
        (function () {
            var socket = new Ws("ws://localhost:8080/websocket");

            socket.OnConnect(function(){
                // update the rest of connected clients, including "myself" when "my" connection is 100% ready.
                socket.Emit("visit");
            });


            socket.On("visit", function (newCount) {
                console.log("visit websocket event with newCount of: ", newCount);

                var text = "1 online visitor";
                if (newCount > 1) {
                    text = newCount + " online visitors";
                }
                document.getElementById("online_visitors").innerHTML = text;
            });

            socket.OnDisconnect(function () {
                document.getElementById("online_visitors").innerHTML = "you've been disconnected";
            });

        })();
    </script>

</body>

</html>

ORM

Xorm

File: orm/xorm/main.go

// Package main shows how an orm can be used within your web app
// it just inserts a column and select the first.
package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

/*
	go get -u github.com/mattn/go-sqlite3
	go get -u github.com/go-xorm/xorm

	If you're on win64 and you can't install go-sqlite3:
		1. Download: https://sourceforge.net/projects/mingw-w64/files/latest/download
		2. Select "x86_x64" and "posix"
		3. Add C:\Program Files\mingw-w64\x86_64-7.1.0-posix-seh-rt_v5-rev1\mingw64\bin
		to your PATH env variable.

	Docs: http://xorm.io/docs/
*/

// User is our user table structure.
type User struct {
	ID        int64  // auto-increment by-default by xorm
	Version   string `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string    `xorm:"varchar(200)"`
	Languages string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func main() {
	app := iris.New()

	orm, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	err = orm.Sync2(new(User))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}

	app.Get("/insert", func(ctx iris.Context) {
		user := &User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		orm.Insert(user)

		ctx.Writef("user inserted: %#v", user)
	})

	app.Get("/get", func(ctx iris.Context) {
		user := User{ID: 1}
		if ok, _ := orm.Get(&user); ok {
			ctx.Writef("user found: %#v", user)
		}
	})

	// http://localhost:8080/insert
	// http://localhost:8080/get
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

Overview

Views

File: overview/views/user/create_verification.html

<html>
    <head><title>Create verification</title></head>
    <body>
        <h1> Create Verification </h1>
        <table style="width:550px">
        <tr>
            <th>Username</th>
            <th>Firstname</th>
            <th>Lastname</th>
            <th>City</th>
            <th>Age</th>
        </tr>
        <tr>
            <td>{{ .Username }}</td>
            <td>{{ .Firstname }}</td>
            <td>{{ .Lastname }}</td>
            <td>{{ .City }}</td>
            <td>{{ .Age }}</td>
        </tr>
        </table>
    </body>
</html>

File: overview/views/user/profile.html

<html>
    <head><title>Profile page</title></head>
    <body>
        <h1> Profile </h1>
        <b> {{ .Username }} </b>
    </body>
</html>

Routing

Basic

File: routing/basic/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// registers a custom handler for 404 not found http (error) status code,
	// fires when route not found or manually by ctx.StatusCode(iris.StatusNotFound).
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	// GET -> HTTP Method
	// / -> Path
	// func(ctx iris.Context) -> The route's handler.
	//
	// Third receiver should contains the route's handler(s), they are executed by order.
	app.Handle("GET", "/", func(ctx iris.Context) {
		// navigate to the middle of $GOPATH/src/github.com/kataras/iris/context/context.go
		// to overview all context's method (there a lot of them, read that and you will learn how iris works too)
		ctx.HTML("Hello from " + ctx.Path()) // Hello from /
	})

	app.Get("/home", func(ctx iris.Context) {
		ctx.Writef(`Same as app.Handle("GET", "/", [...])`)
	})

	app.Get("/donate", donateHandler, donateFinishHandler)

	// Pssst, don't forget dynamic-path example for more "magic"!
	app.Get("/api/users/{userid:int min(1)}", func(ctx iris.Context) {
		userID, err := ctx.Params().GetInt("userid")

		if err != nil {
			ctx.Writef("error while trying to parse userid parameter," +
				"this will never happen if :int is being used because if it's not integer it will fire Not Found automatically.")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		ctx.JSON(map[string]interface{}{
			// you can pass any custom structured go value of course.
			"user_id": userID,
		})
	})
	// app.Post("/", func(ctx iris.Context){}) -> for POST http method.
	// app.Put("/", func(ctx iris.Context){})-> for "PUT" http method.
	// app.Delete("/", func(ctx iris.Context){})-> for "DELETE" http method.
	// app.Options("/", func(ctx iris.Context){})-> for "OPTIONS" http method.
	// app.Trace("/", func(ctx iris.Context){})-> for "TRACE" http method.
	// app.Head("/", func(ctx iris.Context){})-> for "HEAD" http method.
	// app.Connect("/", func(ctx iris.Context){})-> for "CONNECT" http method.
	// app.Patch("/", func(ctx iris.Context){})-> for "PATCH" http method.
	// app.Any("/", func(ctx iris.Context){}) for all http methods.

	// More than one route can contain the same path with a different http mapped method.
	// You can catch any route creation errors with:
	// route, err := app.Get(...)
	// set a name to a route: route.Name = "myroute"

	// You can also group routes by path prefix, sharing middleware(s) and done handlers.

	adminRoutes := app.Party("/admin", adminMiddleware)

	adminRoutes.Done(func(ctx iris.Context) { // executes always last if ctx.Next()
		ctx.Application().Logger().Infof("response sent to " + ctx.Path())
	})
	// adminRoutes.Layout("/views/layouts/admin.html") // set a view layout for these routes, see more at view examples.

	// GET: http://localhost:8080/admin
	adminRoutes.Get("/", func(ctx iris.Context) {
		// [...]
		ctx.StatusCode(iris.StatusOK) // default is 200 == iris.StatusOK
		ctx.HTML("<h1>Hello from admin/</h1>")

		ctx.Next() // in order to execute the party's "Done" Handler(s)
	})

	// GET: http://localhost:8080/admin/login
	adminRoutes.Get("/login", func(ctx iris.Context) {
		// [...]
	})
	// POST: http://localhost:8080/admin/login
	adminRoutes.Post("/login", func(ctx iris.Context) {
		// [...]
	})

	// subdomains, easier than ever, should add localhost or 127.0.0.1 into your hosts file,
	// etc/hosts on unix or C:/windows/system32/drivers/etc/hosts on windows.
	v1 := app.Party("v1.")
	{ // braces are optional, it's just type of style, to group the routes visually.

		// http://v1.localhost:8080
		v1.Get("/", func(ctx iris.Context) {
			ctx.HTML("Version 1 API. go to <a href='" + ctx.Path() + "/api" + "'>/api/users</a>")
		})

		usersAPI := v1.Party("/api/users")
		{
			// http://v1.localhost:8080/api/users
			usersAPI.Get("/", func(ctx iris.Context) {
				ctx.Writef("All users")
			})
			// http://v1.localhost:8080/api/users/42
			usersAPI.Get("/{userid:int}", func(ctx iris.Context) {
				ctx.Writef("user with id: %s", ctx.Params().Get("userid"))
			})
		}
	}

	// wildcard subdomains.
	wildcardSubdomain := app.Party("*.")
	{
		wildcardSubdomain.Get("/", func(ctx iris.Context) {
			ctx.Writef("Subdomain can be anything, now you're here from: %s", ctx.Subdomain())
		})
	}

	// http://localhost:8080
	// http://localhost:8080/home
	// http://localhost:8080/donate
	// http://localhost:8080/api/users/42
	// http://localhost:8080/admin
	// http://localhost:8080/admin/login
	//
	// http://localhost:8080/api/users/0
	// http://localhost:8080/api/users/blabla
	// http://localhost:8080/wontfound
	//
	// if hosts edited:
	//  http://v1.localhost:8080
	//  http://v1.localhost:8080/api/users
	//  http://v1.localhost:8080/api/users/42
	//  http://anything.localhost:8080
	app.Run(iris.Addr(":8080"))
}

func adminMiddleware(ctx iris.Context) {
	// [...]
	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}

func donateHandler(ctx iris.Context) {
	ctx.Writef("Just like an inline handler, but it can be " +
		"used by other package, anywhere in your project.")

	// let's pass a value to the next handler
	// Values is the way handlers(or middleware) are communicating between each other.
	ctx.Values().Set("donate_url", "https://github.com/kataras/iris#-people")
	ctx.Next() // in order to execute the next handler in the chain, look donate route.
}

func donateFinishHandler(ctx iris.Context) {
	// values can be any type of object so we could cast the value to a string
	// but iris provides an easy to do that, if donate_url is not defined, then it returns an empty string instead.
	donateURL := ctx.Values().GetString("donate_url")
	ctx.Application().Logger().Infof("donate_url value was: " + donateURL)
	ctx.Writef("\n\nDonate sent(?).")
}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}

// Notes:
// A path parameter name should contain only alphabetical letters, symbols, containing '_' and numbers are NOT allowed.
// If route failed to be registered, the app will panic without any warnings
// if you didn't catch the second return value(error) on .Handle/.Get....

// See "file-server/single-page-application" to see how another feature, "WrapRouter", works.

Custom Context

File: routing/custom-context/method-overriding/main.go

package main

// In this package I'll show you how to override the existing Context's functions and methods.
// You can easly navigate to the custom-context example to see how you can add new functions
// to your own context (need a custom handler).
//
// This way is far easier to understand and it's faster when you want to override existing methods:
import (
	"reflect"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Create your own custom Context, put any fields you wanna need.
type MyContext struct {
	// Optional Part 1: embed (optional but required if you don't want to override all context's methods)
	context.Context // it's the context/context.go#context struct but you don't need to know it.
}

var _ context.Context = &MyContext{} // optionally: validate on compile-time if MyContext implements context.Context.

// The only one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the handlers via this "*MyContext".
func (ctx *MyContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

// The second one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the chain of handlers via this "*MyContext".
func (ctx *MyContext) Next() {
	context.Next(ctx)
}

// Override any context's method you want...
// [...]

func (ctx *MyContext) HTML(htmlContents string) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from MyContext")

	ctx.ContentType("text/html")
	return ctx.WriteString(htmlContents)
}

func main() {
	app := iris.New()
	// app.Logger().SetLevel("debug")

	// The only one Required:
	// here is how you define how your own context will
	// be created and acquired from the iris' generic context pool.
	app.ContextPool.Attach(func() context.Context {
		return &MyContext{
			// Optional Part 3:
			Context: context.NewContext(app),
		}
	})

	// Register a view engine on .html files inside the ./view/** directory.
	app.RegisterView(iris.HTML("./view", ".html"))

	// register your route, as you normally do
	app.Handle("GET", "/", recordWhichContextJustForProofOfConcept, func(ctx context.Context) {
		// use the context's overridden HTML method.
		ctx.HTML("<h1> Hello from my custom context's HTML! </h1>")
	})

	// this will be executed by the MyContext.Context
	// if MyContext is not directly define the View function by itself.
	app.Handle("GET", "/hi/{firstname:alphabetical}", recordWhichContextJustForProofOfConcept, func(ctx context.Context) {
		firstname := ctx.Values().GetString("firstname")

		ctx.ViewData("firstname", firstname)
		ctx.Gzip(true)

		ctx.View("hi.html")
	})

	app.Run(iris.Addr(":8080"))
}

// should always print "($PATH) Handler is executing from 'MyContext'"
func recordWhichContextJustForProofOfConcept(ctx context.Context) {
	ctx.Application().Logger().Infof("(%s) Handler is executing from: '%s'", ctx.Path(), reflect.TypeOf(ctx).Elem().Name())
	ctx.Next()
}

// Look "new-implementation" to see how you can create an entirely new Context with new functions.

File: routing/custom-context/method-overriding/view/hi.html

<h1> Hi {{.firstname}} </h1>

File: routing/custom-context/new-implementation/main.go

package main

import (
	"sync"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// Owner is our application structure, it contains the methods or fields we need,
// think it as the owner of our *Context.
type Owner struct {
	// define here the fields that are global
	// and shared to all clients.
	sessionsManager *sessions.Sessions
}

// this package-level variable "application" will be used inside context to communicate with our global Application.
var owner = &Owner{
	sessionsManager: sessions.New(sessions.Config{Cookie: "mysessioncookie"}),
}

// Context is our custom context.
// Let's implement a context which will give us access
// to the client's Session with a trivial `ctx.Session()` call.
type Context struct {
	iris.Context
	session *sessions.Session
}

// Session returns the current client's session.
func (ctx *Context) Session() *sessions.Session {
	// this help us if we call `Session()` multiple times in the same handler
	if ctx.session == nil {
		// start a new session if not created before.
		ctx.session = owner.sessionsManager.Start(ctx.Context)
	}

	return ctx.session
}

// Bold will send a bold text to the client.
func (ctx *Context) Bold(text string) {
	ctx.HTML("<b>" + text + "</b>")
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	ctx.session = nil      // reset the session
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

// Handler will convert our handler of func(*Context) to an iris Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

func newApp() *iris.Application {
	app := iris.New()

	// Work as you did before, the only difference
	// is that the original context.Handler should be wrapped with our custom
	// `Handler` function.
	app.Get("/", Handler(func(ctx *Context) {
		ctx.Bold("Hello from our *Context")
	}))

	app.Post("/set", Handler(func(ctx *Context) {
		nameFieldValue := ctx.FormValue("name")
		ctx.Session().Set("name", nameFieldValue)
		ctx.Writef("set session = " + nameFieldValue)
	}))

	app.Get("/get", Handler(func(ctx *Context) {
		name := ctx.Session().GetString("name")
		ctx.Writef(name)
	}))

	return app
}

func main() {
	app := newApp()

	// GET: http://localhost:8080
	// POST: http://localhost:8080/set
	// GET: http://localhost:8080/get
	app.Run(iris.Addr(":8080"))
}

File: routing/custom-context/new-implementation/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestCustomContextNewImpl(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.URL("http://localhost:8080"))

	e.GET("/").Expect().
		Status(httptest.StatusOK).
		ContentType("text/html").
		Body().Equal("<b>Hello from our *Context</b>")

	expectedName := "iris"
	e.POST("/set").WithFormField("name", expectedName).Expect().
		Status(httptest.StatusOK).
		Body().Equal("set session = " + expectedName)

	e.GET("/get").Expect().
		Status(httptest.StatusOK).
		Body().Equal(expectedName)
}

Custom Wrapper

File: routing/custom-wrapper/main.go

package main

import (
	"net/http"
	"strings"

	"github.com/kataras/iris"
)

// In this example you'll just see one use case of .WrapRouter.
// You can use the .WrapRouter to add custom logic when or when not the router should
// be executed in order to execute the registered routes' handlers.
//
// To see how you can serve files on root "/" without a custom wrapper
// just navigate to the "file-server/single-page-application" example.
//
// This is just for the proof of concept, you can skip this tutorial if it's too much for you.
func newApp() *iris.Application {

	app := iris.New()

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./public/index.html", false)
	})

	app.Get("/profile/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("username"))
	})

	// serve files from the root "/", if we used .StaticWeb it could override
	// all the routes because of the underline need of wildcard.
	// Here we will see how you can by-pass this behavior
	// by creating a new file server handler and
	// setting up a wrapper for the router(like a "low-level" middleware)
	// in order to manually check if we want to process with the router as normally
	// or execute the file server handler instead.

	// use of the .StaticHandler
	// which is the same as StaticWeb but it doesn't
	// registers the route, it just returns the handler.
	fileServer := app.StaticHandler("./public", false, false)

	// wrap the router with a native net/http handler.
	// if url does not contain any "." (i.e: .css, .js...)
	// (depends on the app , you may need to add more file-server exceptions),
	// then the handler will execute the router that is responsible for the
	// registered routes (look "/" and "/profile/{username}")
	// if not then it will serve the files based on the root "/" path.
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		// Note that if path has suffix of "index.html" it will auto-permant redirect to the "/",
		// so our first handler will be executed instead.

		if !strings.Contains(path, ".") { // if it's not a resource then continue to the router as normally.
			router(w, r)
			return
		}
		// acquire and release a context in order to use it to execute
		// our file server
		// remember: we use net/http.Handler because here we are in the "low-level", before the router itself.
		ctx := app.ContextPool.Acquire(w, r)
		fileServer(ctx)
		app.ContextPool.Release(ctx)
	})

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080
	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	// http://localhost:8080/profile/anyusername
	app.Run(iris.Addr(":8080"))

	// Note: In this example we just saw one use case,
	// you may want to .WrapRouter or .Downgrade in order to bypass the iris' default router, i.e:
	// you can use that method to setup custom proxies too.
	//
	// If you just want to serve static files on other path than root
	// you can just use the StaticWeb, i.e:
	// 					     .StaticWeb("/static", "./public")
	// ________________________________requestPath, systemPath
}

File: routing/custom-wrapper/main_test.go

package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
)

type resource string

func (r resource) String() string {
	return string(r)
}

func (r resource) strip(strip string) string {
	s := r.String()
	return strings.TrimPrefix(s, strip)
}

func (r resource) loadFromBase(dir string) string {
	filename := r.String()

	if filename == "/" {
		filename = "/index.html"
	}

	fullpath := filepath.Join(dir, filename)

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(fullpath + " failed with error: " + err.Error())
	}

	return string(b)
}

var urls = []resource{
	"/",
	"/index.html",
	"/app.js",
	"/css/main.css",
}

func TestCustomWrapper(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	for _, u := range urls {
		url := u.String()
		contents := u.loadFromBase("./public")

		e.GET(url).Expect().
			Status(httptest.StatusOK).
			Body().Equal(contents)
	}
}

File: routing/custom-wrapper/public/app.js

window.alert("app.js loaded from \"/");

File: routing/custom-wrapper/public/css/main.css

body {
    background-color: black;
}

File: routing/custom-wrapper/public/index.html

<html>

<head>
    <title>{{ .Page.Title }}</title>
</head>

<body>
    <h1> Hello from index.html </h1>


    <script src="/app.js">  </script>
</body>

</html>

Dynamic Path

File: routing/dynamic-path/main.go

package main

import (
	"regexp"
	"strconv"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// At the previous example "routing/basic",
	// we've seen static routes, group of routes, subdomains, wildcard subdomains, a small example of parameterized path
	// with a single known paramete and custom http errors, now it's time to see wildcard parameters and macros.

	// iris, like net/http std package registers route's handlers
	// by a Handler, the iris' type of handler is just a func(ctx iris.Context)
	// where context comes from github.com/kataras/iris/context.
	// Until go 1.9 you will have to import that package too, after go 1.9 this will be not be necessary.
	//
	// iris has the easiest and the most powerful routing process you have ever meet.
	//
	// At the same time,
	// iris has its own interpeter(yes like a programming language)
	// for route's path syntax and their dynamic path parameters parsing and evaluation,
	// We call them "macros" for shortcut.
	// How? It calculates its needs and if not any special regexp needed then it just
	// registers the route with the low-level underline  path syntax,
	// otherwise it pre-compiles the regexp and adds the necessary middleware(s).
	//
	// Standard macro types for parameters:
	//  +------------------------+
	//  | {param:string}         |
	//  +------------------------+
	// string type
	// anything
	//
	//  +------------------------+
	//  | {param:int}            |
	//  +------------------------+
	// int type
	// only numbers (0-9)
	//
	// +------------------------+
	// | {param:long}           |
	// +------------------------+
	// int64 type
	// only numbers (0-9)
	//
	// +------------------------+
	// | {param:boolean}        |
	// +------------------------+
	// bool type
	// only "1" or "t" or "T" or "TRUE" or "true" or "True"
	// or "0" or "f" or "F" or "FALSE" or "false" or "False"
	//
	//  +------------------------+
	//  | {param:alphabetical}   |
	//  +------------------------+
	// alphabetical/letter type
	// letters only (upper or lowercase)
	//
	//  +------------------------+
	//  | {param:file}           |
	//  +------------------------+
	// file type
	// letters (upper or lowercase)
	// numbers (0-9)
	// underscore (_)
	// dash (-)
	// point (.)
	// no spaces ! or other character
	//
	//  +------------------------+
	//  | {param:path}           |
	//  +------------------------+
	// path type
	// anything, should be the last part, more than one path segment,
	// i.e: /path1/path2/path3 , ctx.Params().Get("param") == "/path1/path2/path3"
	//
	// if type is missing then parameter's type is defaulted to string, so
	// {param} == {param:string}.
	//
	// If a function not found on that type then the `string` macro type's functions are being used.
	//
	//
	// Besides the fact that iris provides the basic types and some default "macro funcs"
	// you are able to register your own too!.
	//
	// Register a named path parameter function:
	// app.Macros().Int.RegisterFunc("min", func(argument int) func(paramValue string) bool {
	//  [...]
	//  return true/false -> true means valid.
	// })
	//
	// at the func(argument ...) you can have any standard type, it will be validated before the server starts
	// so don't care about performance here, the only thing it runs at serve time is the returning func(paramValue string) bool.
	//
	// {param:string equal(iris)} , "iris" will be the argument here:
	// app.Macros().String.RegisterFunc("equal", func(argument string) func(paramValue string) bool {
	// 	return func(paramValue string) bool { return argument == paramValue }
	// })

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

	// Another example using a custom regexp and any custom logic.
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, err := regexp.Compile(latLonExpr)
	if err != nil {
		panic(err)
	}

	app.Macros().String.RegisterFunc("coordinate", func() func(paramName string) (ok bool) {
		// MatchString is a type of func(string) bool, so we can return that as it's.
		return latLonRegex.MatchString
	})

	app.Get("/coordinates/{lat:string coordinate() else 502}/{lon:string coordinate() else 502}", func(ctx iris.Context) {
		ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
	})

	//

	// http://localhost:8080/game/a-zA-Z/level/0-9
	// remember, alphabetical is lowercase or uppercase letters only.
	app.Get("/game/{name:alphabetical}/level/{level:int}", func(ctx iris.Context) {
		ctx.Writef("name: %s | level: %s", ctx.Params().Get("name"), ctx.Params().Get("level"))
	})

	app.Get("/lowercase/static", func(ctx iris.Context) {
		ctx.Writef("static and dynamic paths are not conflicted anymore!")
	})

	// let's use a trivial custom regexp that validates a single path parameter
	// which its value is only lowercase letters.

	// http://localhost:8080/lowercase/anylowercase
	app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
		ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %s", ctx.Params().Get("name"))
	})

	// http://localhost:8080/single_file/app.js
	app.Get("/single_file/{myfile:file}", func(ctx iris.Context) {
		ctx.Writef("file type validates if the parameter value has a form of a file name, got: %s", ctx.Params().Get("myfile"))
	})

	// http://localhost:8080/myfiles/any/directory/here/
	// this is the only macro type that accepts any number of path segments.
	app.Get("/myfiles/{directory:path}", func(ctx iris.Context) {
		ctx.Writef("path type accepts any number of path segments, path after /myfiles/ is: %s", ctx.Params().Get("directory"))
	}) // for wildcard path (any number of path segments) without validation you can use:
	// /myfiles/*

	// "{param}"'s performance is exactly the same of ":param"'s.

	// alternatives -> ":param" for single path parameter and "*" for wildcard path parameter.
	// Note these:
	// if  "/mypath/*" then the parameter name is "*".
	// if  "/mypath/{myparam:path}" then the parameter has two names, one is the "*" and the other is the user-defined "myparam".

	// WARNING:
	// A path parameter name should contain only alphabetical letters. Symbols like  '_' and numbers are NOT allowed.
	// Last, do not confuse `ctx.Params()` with `ctx.Values()`.
	// Path parameter's values goes to `ctx.Params()` and context's local storage
	// that can be used to communicate between handlers and middleware(s) goes to
	// `ctx.Values()`.
	app.Run(iris.Addr(":8080"))
}

File: routing/dynamic-path/root-wildcard/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// this works as expected now,
	// will handle all GET requests
	// except:
	// /                     -> because of app.Get("/", ...)
	// /other/anything/here  -> because of app.Get("/other/{paramother:path}", ...)
	// /other2/anything/here -> because of app.Get("/other2/{paramothersecond:path}", ...)
	// /other2/static2        -> because of app.Get("/other2/static", ...)
	//
	// It isn't conflicts with the rest of the routes, without routing performance cost!
	//
	// i.e /something/here/that/cannot/be/found/by/other/registered/routes/order/not/matters
	app.Get("/{p:path}", h)
	// app.Get("/static/{p:path}", staticWildcardH)

	// this will handle only GET /
	app.Get("/", staticPath)

	// this will handle all GET requests starting with "/other/"
	//
	// i.e /other/more/than/one/path/parts
	app.Get("/other/{paramother:path}", other)

	// this will handle all GET requests starting with "/other2/"
	// except /other2/static (because of the next static route)
	//
	// i.e /other2/more/than/one/path/parts
	app.Get("/other2/{paramothersecond:path}", other2)

	// this will handle only GET "/other2/static"
	app.Get("/other2/static2", staticPathOther2)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func h(ctx iris.Context) {
	param := ctx.Params().Get("p")
	ctx.WriteString(param)
}

func staticWildcardH(ctx iris.Context) {
	param := ctx.Params().Get("p")
	ctx.WriteString("from staticWildcardH: param=" + param)
}

func other(ctx iris.Context) {
	param := ctx.Params().Get("paramother")
	ctx.Writef("from other: %s", param)
}

func other2(ctx iris.Context) {
	param := ctx.Params().Get("paramothersecond")
	ctx.Writef("from other2: %s", param)
}

func staticPath(ctx iris.Context) {
	ctx.Writef("from the static path(/): %s", ctx.Path())
}

func staticPathOther2(ctx iris.Context) {
	ctx.Writef("from the static path(/other2/static2): %s", ctx.Path())
}

HTTP Errors

File: routing/http-errors/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML("Message: <b>" + ctx.Values().GetString("message") + "</b>")
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(`Click <a href="/my500">here</a> to fire the 500 status code`)
	})

	app.Get("/my500", func(ctx iris.Context) {
		ctx.Values().Set("message", "this is the error message")
		ctx.StatusCode(500)
	})

	app.Get("/u/{firstname:alphabetical}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("firstname"))
	})

	app.Run(iris.Addr(":8080"))
}

Overview

File: routing/overview/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// GET: http://localhost:8080
	app.Get("/", info)

	// GET: http://localhost:8080/profile/anyusername
	//
	// Want to use a custom regex expression instead?
	// Easy: app.Get("/profile/{username:string regexp(^[a-zA-Z ]+$)}")
	app.Get("/profile/{username:string}", info)

	// If parameter type is missing then it's string which accepts anything,
	// i.e: /{paramname} it's exactly the same as /{paramname:string}.
	// The below is exactly the same as
	// {username:string}
	//
	// GET: http://localhost:8080/profile/anyusername/backups/any/number/of/paths/here
	app.Get("/profile/{username}/backups/{filepath:path}", info)

	// Favicon

	// GET: http://localhost:8080/favicon.ico
	app.Favicon("./public/images/favicon.ico")

	// Static assets

	// GET: http://localhost:8080/assets/css/bootstrap.min.css
	//	    maps to ./public/assets/css/bootstrap.min.css file at system location.
	// GET: http://localhost:8080/assets/js/react.min.js
	//      maps to ./public/assets/js/react.min.js file at system location.
	app.StaticWeb("/assets", "./public/assets")

	/* OR

	// GET: http://localhost:8080/js/react.min.js
	// 		maps to ./public/assets/js/react.min.js file at system location.
	app.StaticWeb("/js", "./public/assets/js")

	// GET: http://localhost:8080/css/bootstrap.min.css
	// 		maps to ./public/assets/css/bootstrap.min.css file at system location.
	app.StaticWeb("/css", "./public/assets/css")

	*/

	// Grouping

	usersRoutes := app.Party("/users")
	// GET: http://localhost:8080/users/help
	usersRoutes.Get("/help", func(ctx iris.Context) {
		ctx.Writef("GET / -- fetch all users\n")
		ctx.Writef("GET /$ID -- fetch a user by id\n")
		ctx.Writef("POST / -- create new user\n")
		ctx.Writef("PUT /$ID -- update an existing user\n")
		ctx.Writef("DELETE /$ID -- delete an existing user\n")
	})

	// GET: http://localhost:8080/users
	usersRoutes.Get("/", func(ctx iris.Context) {
		ctx.Writef("get all users")
	})

	// GET: http://localhost:8080/users/42
	// **/users/42 and /users/help works after iris version 7.0.5**
	usersRoutes.Get("/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("get user by id: %d", id)
	})

	// POST: http://localhost:8080/users
	usersRoutes.Post("/", func(ctx iris.Context) {
		username, password := ctx.PostValue("username"), ctx.PostValue("password")
		ctx.Writef("create user for username= %s and password= %s", username, password)
	})

	// PUT: http://localhost:8080/users
	usersRoutes.Put("/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id") // or .Get to get its string represatantion.
		username := ctx.PostValue("username")
		ctx.Writef("update user for id= %d and new username= %s", id, username)
	})

	// DELETE: http://localhost:8080/users/42
	usersRoutes.Delete("/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("delete user by id: %d", id)
	})

	// Subdomains, depends on the host, you have to edit the hosts or nginx/caddy's configuration if you use them.
	//
	// See more subdomains examples at _examples/subdomains folder.
	adminRoutes := app.Party("admin.")

	// GET: http://admin.localhost:8080
	adminRoutes.Get("/", info)
	// GET: http://admin.localhost:8080/settings
	adminRoutes.Get("/settings", info)

	// Wildcard/dynamic subdomain
	dynamicSubdomainRoutes := app.Party("*.")

	// GET: http://any_thing_here.localhost:8080
	dynamicSubdomainRoutes.Get("/", info)

	app.Delete("/something", func(ctx iris.Context) {
		name := ctx.URLParam("name")
		ctx.Writef(name)
	})

	// GET: http://localhost:8080/
	// GET: http://localhost:8080/profile/anyusername
	// GET: http://localhost:8080/profile/anyusername/backups/any/number/of/paths/here

	// GET: http://localhost:8080/users/help
	// GET: http://localhost:8080/users
	// GET: http://localhost:8080/users/42
	// POST: http://localhost:8080/users
	// PUT: http://localhost:8080/users
	// DELETE: http://localhost:8080/users/42
	// DELETE: http://localhost:8080/something?name=iris

	// GET: http://admin.localhost:8080
	// GET: http://admin.localhost:8080/settings
	// GET: http://any_thing_here.localhost:8080
	app.Run(iris.Addr(":8080"))
}

func info(ctx iris.Context) {
	method := ctx.Method()       // the http method requested a server's resource.
	subdomain := ctx.Subdomain() // the subdomain, if any.

	// the request path (without scheme and host).
	path := ctx.Path()
	// how to get all parameters, if we don't know
	// the names:
	paramsLen := ctx.Params().Len()

	ctx.Params().Visit(func(name string, value string) {
		ctx.Writef("%s = %s\n", name, value)
	})
	ctx.Writef("\nInfo\n\n")
	ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s\nParameters length: %d", method, subdomain, path, paramsLen)
}

Reverse

File: routing/reverse/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func main() {
	app := iris.New()
	// need for manually reverse routing when needed outside of view engine.
	// you normally don't need it because of the {{ urlpath "routename" "path" "values" "here"}}
	rv := router.NewRoutePathReverser(app)

	myroute := app.Get("/anything/{anythingparameter:path}", func(ctx iris.Context) {
		paramValue := ctx.Params().Get("anythingparameter")
		ctx.Writef("The path after /anything is: %s", paramValue)
	})

	myroute.Name = "myroute"

	// useful for links, although iris' view engine has the {{ urlpath "routename" "path values"}} already.
	app.Get("/reverse_myroute", func(ctx iris.Context) {
		myrouteRequestPath := rv.Path(myroute.Name, "any/path")
		ctx.HTML("Should be <b>/anything/any/path</b>: " + myrouteRequestPath)
	})

	// execute a route, similar to redirect but without redirect :)
	app.Get("/execute_myroute", func(ctx iris.Context) {
		ctx.Exec("GET", "/anything/any/path") // like it was called by the client.
	})

	// http://localhost:8080/reverse_myroute
	// http://localhost:8080/execute_myroute
	// http://localhost:8080/anything/any/path/here
	//
	// See view/template_html_4 example for more reverse routing examples
	// using the reverse router component and the {{url}} and {{urlpath}} template functions.
	app.Run(iris.Addr(":8080"))

}

Route State

File: routing/route-state/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	none := app.None("/invisible/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s with method: %s", ctx.Params().Get("username"), ctx.Method())

		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/change", func(ctx iris.Context) {

		if none.IsOnline() {
			none.Method = iris.MethodNone
		} else {
			none.Method = iris.MethodGet
		}

		// refresh re-builds the router at serve-time in order to be notified for its new routes.
		app.RefreshRouter()
	})

	app.Get("/execute", func(ctx iris.Context) {
		// same as navigating to "http://localhost:8080/invisible/iris" when /change has being invoked and route state changed
		// from "offline" to "online"
		ctx.Values().Set("from", "/execute") // values and session can be shared when calling Exec from a "foreign" context.
		ctx.Exec("GET", "/invisible/iris")
	})

	app.Run(iris.Addr(":8080"))
}

Writing A Middleware

File: routing/writing-a-middleware/globally/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	// register the "before" handler as the first handler which will be executed
	// on all domain's routes.
	// Or use the `UseGlobal` to register a middleware which will fire across subdomains.
	// app.Use(before)
	// register the "after" handler as the last handler which will be executed
	// after all domain's routes' handler(s).
	//
	// Or use the `DoneGlobal` to append handlers that will be fired globally.
	// app.Done(after)

	// register our routes.
	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

	// Order of those calls doesn't matter, `UseGlobal` and `DoneGlobal`
	// are applied to existing routes and future routes.
	//
	// Remember: the `Use` and `Done` are applied to the current party's and its children,
	// so if we used the `app.Use/Don`e before the routes registration
	// it would work like UseGlobal/DoneGlobal in this case, because the `app` is the root party.
	//
	// See `app.Party/PartyFunc` for more.
	app.UseGlobal(before)
	app.DoneGlobal(after)

	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next()
}

func after(ctx iris.Context) {
	println("After the indexHandler or contactHandler")
}

func indexHandler(ctx iris.Context) {
	println("Inside indexHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after" handler registered via `DoneGlobal`.
}

func contactHandler(ctx iris.Context) {
	println("Inside contactHandler")

	// write something to the client as a response.
	ctx.HTML("<h1>Contact</h1>")

	ctx.Next() // execute the "after" handler registered via `DoneGlobal`.
}

File: routing/writing-a-middleware/per-route/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	// or app.Use(before) and app.Done(after).
	app.Get("/", before, mainHandler, after)

	// Use registers a middleware(prepend handlers) to all party's, and its children that will be registered
	// after.
	//
	// (`app` is the root children so those use and done handlers will be registered everywhere)
	app.Use(func(ctx iris.Context) {
		println(`before the party's routes and its children,
but this is not applied to the '/' route
because it's registered before the middleware, order matters.`)
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {
		println("this is executed always last, if the previous handler calls the `ctx.Next()`, it's the reverse of `.Use`")
		message := ctx.Values().GetString("message")
		println("message: " + message)
	})

	app.Get("/home", func(ctx iris.Context) {
		ctx.HTML("<h1> Home </h1>")
		ctx.Values().Set("message", "this is the home message, ip: "+ctx.RemoteAddr())
		ctx.Next() // call the done handlers.
	})

	child := app.Party("/child")
	child.Get("/", func(ctx iris.Context) {
		ctx.Writef(`this is the localhost:8080/child route.
All Use and Done handlers that are registered to the parent party,
are applied here as well.`)
		ctx.Next() // call the done handlers.
	})

	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}

Sessions

Database

File: sessions/database/badger/main.go

package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/badger"
)

func main() {
	db, err := badger.New("./data")
	if err != nil {
		panic(err)
	}

	// close and unlock the database when control+C/cmd+C pressed
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	defer db.Close() // close and unlock the database if application errored.

	sess := sessions.New(sessions.Config{
		Cookie:       "sessionscookieid",
		Expires:      45 * time.Minute, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	})

	//
	// IMPORTANT:
	//
	sess.UseDatabase(db)

	// the rest of the code stays the same.
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		//set session values
		s.Set("name", "iris")

		//test if setted here
		ctx.Writef("All ok session value of the 'name' is: %s", s.GetString("name"))
	})

	app.Get("/set/{key}/{value}", func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		s := sess.Start(ctx)
		// set session values
		s.Set(key, value)

		// test if setted here
		ctx.Writef("All ok session value of the '%s' is: %s", key, s.GetString(key))
	})

	app.Get("/get", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The 'name' on the /set was: %s", name)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		// delete a specific key
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// removes all entries
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy", func(ctx iris.Context) {
		//destroy, removes the entire session data and cookie
		sess.Destroy(ctx)
	})

	app.Get("/update", func(ctx iris.Context) {
		// updates expire date with a new date
		sess.ShiftExpiration(ctx)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

File: sessions/database/boltdb/main.go

package main

import (
	"os"
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
)

func main() {
	db, err := boltdb.New("./sessions.db", os.FileMode(0750))
	if err != nil {
		panic(err)
	}

	// close and unlobkc the database when control+C/cmd+C pressed
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	defer db.Close() // close and unlock the database if application errored.

	sess := sessions.New(sessions.Config{
		Cookie:       "sessionscookieid",
		Expires:      45 * time.Minute, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	})

	//
	// IMPORTANT:
	//
	sess.UseDatabase(db)

	// the rest of the code stays the same.
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		//set session values
		s.Set("name", "iris")

		//test if setted here
		ctx.Writef("All ok session value of the 'name' is: %s", s.GetString("name"))
	})

	app.Get("/set/{key}/{value}", func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		s := sess.Start(ctx)
		// set session values
		s.Set(key, value)

		// test if setted here
		ctx.Writef("All ok session value of the '%s' is: %s", key, s.GetString(key))
	})

	app.Get("/get", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The 'name' on the /set was: %s", name)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		// delete a specific key
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// removes all entries
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy", func(ctx iris.Context) {
		//destroy, removes the entire session data and cookie
		sess.Destroy(ctx)
	})

	app.Get("/update", func(ctx iris.Context) {
		// updates expire date with a new date
		sess.ShiftExpiration(ctx)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

File: sessions/database/redis/main.go

package main

import (
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

// tested with redis version 3.0.503.
// for windows see: https://github.com/ServiceStack/redis-windows
func main() {
	// replace with your running redis' server settings:
	db := redis.New(service.Config{
		Network:     "tcp",
		Addr:        "127.0.0.1:6379",
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: time.Duration(5) * time.Minute,
		Prefix:      ""}) // optionally configure the bridge between your redis server

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	defer db.Close() // close the database connection if application errored.

	sess := sessions.New(sessions.Config{
		Cookie:       "sessionscookieid",
		Expires:      45 * time.Minute, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	},
	)

	//
	// IMPORTANT:
	//
	sess.UseDatabase(db)

	// the rest of the code stays the same.
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		//set session values
		s.Set("name", "iris")

		//test if setted here
		ctx.Writef("All ok session value of the 'name' is: %s", s.GetString("name"))
	})

	app.Get("/set/{key}/{value}", func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		s := sess.Start(ctx)
		// set session values
		s.Set(key, value)

		// test if setted here
		ctx.Writef("All ok session value of the '%s' is: %s", key, s.GetString(key))
	})

	app.Get("/set/int/{key}/{value}", func(ctx iris.Context) {
		key := ctx.Params().Get("key")
		value, _ := ctx.Params().GetInt("value")
		s := sess.Start(ctx)
		// set session values
		s.Set(key, value)
		valueSet := s.Get(key)
		// test if setted here
		ctx.Writef("All ok session value of the '%s' is: %v", key, valueSet)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		key := ctx.Params().Get("key")
		value := sess.Start(ctx).Get(key)

		ctx.Writef("The '%s' on the /set was: %v", key, value)
	})

	app.Get("/get", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The 'name' on the /set was: %s", name)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		// delete a specific key
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// removes all entries
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy", func(ctx iris.Context) {
		//destroy, removes the entire session data and cookie
		sess.Destroy(ctx)
	})

	app.Get("/update", func(ctx iris.Context) {
		// updates expire date with a new date
		sess.ShiftExpiration(ctx)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

Flash Messages

File: sessions/flash-messages/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
)

func main() {
	app := iris.New()
	sess := sessions.New(sessions.Config{Cookie: "myappsessionid", AllowReclaim: true})

	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		s.SetFlash("name", "iris")
		ctx.Writef("Message setted, is available for the next request")
	})

	app.Get("/get", func(ctx iris.Context) {
		s := sess.Start(ctx)
		name := s.GetFlashString("name")
		if name == "" {
			ctx.Writef("Empty name!!")
			return
		}
		ctx.Writef("Hello %s", name)
	})

	app.Get("/test", func(ctx iris.Context) {
		s := sess.Start(ctx)
		name := s.GetFlashString("name")
		if name == "" {
			ctx.Writef("Empty name!!")
			return
		}

		ctx.Writef("Ok you are coming from /set ,the value of the name is %s", name)
		ctx.Writef(", and again from the same context: %s", name)
	})

	app.Run(iris.Addr(":8080"))
}

Overview

File: sessions/overview/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
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
	// ...

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

Securecookie

File: sessions/securecookie/main.go

package main

// developers can use any library to add a custom cookie encoder/decoder.
// At this example we use the gorilla's securecookie package:
// $ go get github.com/gorilla/securecookie
// $ go run main.go

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

		//set session values
		s := mySessions.Start(ctx)
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
		// delete a specific key
		s := mySessions.Start(ctx)
		s.Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// removes all entries
		mySessions.Start(ctx).Clear()
	})

	app.Get("/update", func(ctx iris.Context) {
		// updates expire date with a new date
		mySessions.ShiftExpiration(ctx)
	})

	app.Get("/destroy", func(ctx iris.Context) {
		//destroy, removes the entire session data and cookie
		mySessions.Destroy(ctx)
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

File: sessions/securecookie/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func TestSessionsEncodeDecode(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.URL("http://example.com"))

	es := e.GET("/set").Expect()
	es.Status(iris.StatusOK)
	es.Cookies().NotEmpty()
	es.Body().Equal("All ok session setted to: iris")

	e.GET("/get").Expect().Status(iris.StatusOK).Body().Equal("The name on the /set was: iris")
	// delete and re-get
	e.GET("/delete").Expect().Status(iris.StatusOK)
	e.GET("/get").Expect().Status(iris.StatusOK).Body().Equal("The name on the /set was: ")
	// set, clear and re-get
	e.GET("/set").Expect().Body().Equal("All ok session setted to: iris")
	e.GET("/clear").Expect().Status(iris.StatusOK)
	e.GET("/get").Expect().Status(iris.StatusOK).Body().Equal("The name on the /set was: ")
}

Standalone

File: sessions/standalone/main.go

package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
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
		// All it does is that it removes the cookie for both `Request` and `ResponseWriter` while `Destroy`
		// or add a new cookie to `Request` while `Start`.
		//
		// Defaults to false.
		AllowReclaim: true,
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx iris.Context) {
		//set session values.
		s := sess.Start(ctx)
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
		// get a specific value, as string,
		// if not found then it returns just an empty string.
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		// delete a specific key
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// removes all entries.
		sess.Start(ctx).Clear()
	})

	app.Get("/update", func(ctx iris.Context) {
		// updates expire date.
		sess.ShiftExpiration(ctx)
	})

	app.Get("/destroy", func(ctx iris.Context) {
		//destroy, removes the entire session data and cookie
		sess.Destroy(ctx)
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
		// change will affect the underline array of the session's value "businessEdit", but now it will not.
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
			panic("Report this as a bug, immutable data cannot be changed from the caller without re-SetImmutable")
		}

		ctx.Writef("[]businessModel[0].Name remains: %s", firstModel.Name)

		// the name should remains "Edward"
	})

	app.Run(iris.Addr(":8080"))
}

Structuring

Bootstrap

File: structuring/bootstrap/bootstrap/bootstrapper.go

package bootstrap

import (
	"time"

	"github.com/gorilla/securecookie"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
}

// SetupSessions initializes the sessions, optionally.
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupWebsockets prepares the websocket server.
func (b *Bootstrapper) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(onConnection)

	b.Get(endpoint, ws.Handler())
	b.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}

File: structuring/bootstrap/main.go

package main

import (
	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
	"github.com/kataras/iris/_examples/structuring/bootstrap/middleware/identity"
	"github.com/kataras/iris/_examples/structuring/bootstrap/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}

File: structuring/bootstrap/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

// go test -v
func TestApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)

	// test our routes
	e.GET("/").Expect().Status(httptest.StatusOK)
	e.GET("/follower/42").Expect().Status(httptest.StatusOK).
		Body().Equal("from /follower/{id:long} with ID: 42")
	e.GET("/following/52").Expect().Status(httptest.StatusOK).
		Body().Equal("from /following/{id:long} with ID: 52")
	e.GET("/like/64").Expect().Status(httptest.StatusOK).
		Body().Equal("from /like/{id:long} with ID: 64")

	// test not found
	e.GET("/notfound").Expect().Status(httptest.StatusNotFound)
	expectedErr := map[string]interface{}{
		"app":     app.AppName,
		"status":  httptest.StatusNotFound,
		"message": "",
	}
	e.GET("/anotfoundwithjson").WithQuery("json", nil).
		Expect().Status(httptest.StatusNotFound).JSON().Equal(expectedErr)
}

File: structuring/bootstrap/middleware/identity/identity.go

package identity

import (
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())

		ctx.Header("Server", "Iris: https://iris-go.com")

		// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h)
}

File: structuring/bootstrap/routes/follower.go

package routes

import (
	"github.com/kataras/iris"
)

// GetFollowerHandler handles the GET: /follower/{id}
func GetFollowerHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/following.go

package routes

import (
	"github.com/kataras/iris"
)

// GetFollowingHandler handles the GET: /following/{id}
func GetFollowingHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/index.go

package routes

import (
	"github.com/kataras/iris"
)

// GetIndexHandler handles the GET: /
func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}

File: structuring/bootstrap/routes/like.go

package routes

import (
	"github.com/kataras/iris"
)

// GetLikeHandler handles the GET: /like/{id}
func GetLikeHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/routes.go

package routes

import (
	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", GetIndexHandler)
	b.Get("/follower/{id:long}", GetFollowerHandler)
	b.Get("/following/{id:long}", GetFollowingHandler)
	b.Get("/like/{id:long}", GetLikeHandler)
}

File: structuring/bootstrap/views/index.html

<h1>Welcome!!</h1>

File: structuring/bootstrap/views/shared/error.html

<h1 class="text-danger">Error.</h1>
<h2 class="text-danger">An error occurred while processing your request.</h2>

<h3>{{.Err.status}}</h3>
<h4>{{.Err.message}}</h4>

File: structuring/bootstrap/views/shared/layout.html

<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico" />
    <title>{{.Title}} - {{.AppName}}</title>

</head>

<body>
    <div>
        <!-- Render the current template here -->
        {{ yield }}
        <hr />
        <footer>
            <p>&copy; 2017 - {{.AppOwner}}</p>
        </footer>
    </div>
</body>

</html>

Login Mvc Single Responsibility Package

File: structuring/login-mvc-single-responsibility-package/main.go

package main

import (
	"time"

	"github.com/kataras/iris/_examples/structuring/login-mvc-single-responsibility-package/user"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/layout.html"))

	app.StaticWeb("/public", "./public")

	mvc.Configure(app, configureMVC)

	// http://localhost:8080/user/register
	// http://localhost:8080/user/login
	// http://localhost:8080/user/me
	// http://localhost:8080/user/logout
	// http://localhost:8080/user/1
	app.Run(iris.Addr(":8080"), configure)
}

func configureMVC(app *mvc.Application) {
	manager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})

	userApp := app.Party("/user")
	userApp.Register(
		user.NewDataSource(),
		manager.Start,
	)
	userApp.Handle(new(user.Controller))
}

func configure(app *iris.Application) {
	app.Configure(
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}

File: structuring/login-mvc-single-responsibility-package/public/css/site.css

/* Bordered form */
form {
    border: 3px solid #f1f1f1;
}

/* Full-width inputs */
input[type=text], input[type=password] {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    box-sizing: border-box;
}

/* Set a style for all buttons */
button {
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    width: 100%;
}

/* Add a hover effect for buttons */
button:hover {
    opacity: 0.8;
}

/* Extra style for the cancel button (red) */
.cancelbtn {
    width: auto;
    padding: 10px 18px;
    background-color: #f44336;
}

/* Center the container */

/* Add padding to containers */
.container {
    padding: 16px;
}

/* The "Forgot password" text */
span.psw {
    float: right;
    padding-top: 16px;
}

/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
    span.psw {
        display: block;
        float: none;
    }
    .cancelbtn {
        width: 100%;
    }
}

File: structuring/login-mvc-single-responsibility-package/user/auth.go

package user

import (
	"errors"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

const sessionIDKey = "UserID"

// paths
var (
	PathLogin  = mvc.Response{Path: "/user/login"}
	PathLogout = mvc.Response{Path: "/user/logout"}
)

// AuthController is the user authentication controller, a custom shared controller.
type AuthController struct {
	// context is auto-binded if struct depends on this,
	// in this controller we don't we do everything with mvc-style,
	// and that's neither the 30% of its features.
	// Ctx iris.Context

	Source  *DataSource
	Session *sessions.Session

	// the whole controller is request-scoped because we already depend on Session, so
	// this will be new for each new incoming request, BeginRequest sets that based on the session.
	UserID int64
}

// BeginRequest saves login state to the context, the user id.
func (c *AuthController) BeginRequest(ctx iris.Context) {
	c.UserID, _ = c.Session.GetInt64(sessionIDKey)
}

// EndRequest is here just to complete the BaseController
// in order to be tell iris to call the `BeginRequest` before the main method.
func (c *AuthController) EndRequest(ctx iris.Context) {}

func (c *AuthController) fireError(err error) mvc.View {
	return mvc.View{
		Code: iris.StatusBadRequest,
		Name: "shared/error.html",
		Data: iris.Map{"Title": "User Error", "Message": strings.ToUpper(err.Error())},
	}
}

func (c *AuthController) redirectTo(id int64) mvc.Response {
	return mvc.Response{Path: "/user/" + strconv.Itoa(int(id))}
}

func (c *AuthController) createOrUpdate(firstname, username, password string) (user Model, err error) {
	username = strings.Trim(username, " ")
	if username == "" || password == "" || firstname == "" {
		return user, errors.New("empty firstname, username or/and password")
	}

	userToInsert := Model{
		Firstname: firstname,
		Username:  username,
		password:  password,
	} // password is hashed by the Source.

	newUser, err := c.Source.InsertOrUpdate(userToInsert)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (c *AuthController) isLoggedIn() bool {
	// we don't search by session, we have the user id
	// already by the `BeginRequest` middleware.
	return c.UserID > 0
}

func (c *AuthController) verify(username, password string) (user Model, err error) {
	if username == "" || password == "" {
		return user, errors.New("please fill both username and password fields")
	}

	u, found := c.Source.GetByUsername(username)
	if !found {
		// if user found with that username not found at all.
		return user, errors.New("user with that username does not exist")
	}

	if ok, err := ValidatePassword(password, u.HashedPassword); err != nil || !ok {
		// if user found but an error occurred or the password is not valid.
		return user, errors.New("please try to login with valid credentials")
	}

	return u, nil
}

// if logged in then destroy the session
// and redirect to the login page
// otherwise redirect to the registration page.
func (c *AuthController) logout() mvc.Response {
	if c.isLoggedIn() {
		c.Session.Destroy()
	}
	return PathLogin
}

File: structuring/login-mvc-single-responsibility-package/user/controller.go

package user

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var (
	// About Code: iris.StatusSeeOther ->
	// When redirecting from POST to GET request you -should- use this HTTP status code,
	// however there're some (complicated) alternatives if you
	// search online or even the HTTP RFC.
	// "See Other" RFC 7231
	pathMyProfile = mvc.Response{Path: "/user/me", Code: iris.StatusSeeOther}
	pathRegister  = mvc.Response{Path: "/user/register"}
)

// Controller is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// GET				/user/{id:long} | long is a new param type, it's the int64.
// All HTTP Methods /user/logout
type Controller struct {
	AuthController
}

type formValue func(string) string

// BeforeActivation called once before the server start
// and before the controller's registration, here you can add
// dependencies, to this controller and only, that the main caller may skip.
func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	// bind the context's `FormValue` as well in order to be
	// acceptable on the controller or its methods' input arguments (NEW feature as well).
	b.Dependencies().Add(func(ctx iris.Context) formValue { return ctx.FormValue })
}

type page struct {
	Title string
}

// GetRegister handles GET:/user/register.
// mvc.Result can accept any struct which contains a `Dispatch(ctx iris.Context)` method.
// Both mvc.Response and mvc.View are mvc.Result.
func (c *Controller) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		return c.logout()
	}

	// You could just use it as a variable to win some time in serve-time,
	// this is an exersise for you :)
	return mvc.View{
		Name: pathRegister.Path + ".html",
		Data: page{"User Registration"},
	}
}

// PostRegister handles POST:/user/register.
func (c *Controller) PostRegister(form formValue) mvc.Result {
	// we can either use the `c.Ctx.ReadForm` or read values one by one.
	var (
		firstname = form("firstname")
		username  = form("username")
		password  = form("password")
	)

	user, err := c.createOrUpdate(firstname, username, password)
	if err != nil {
		return c.fireError(err)
	}

	// setting a session value was never easier.
	c.Session.Set(sessionIDKey, user.ID)
	// succeed, nothing more to do here, just redirect to the /user/me.
	return pathMyProfile
}

// with these static views,
// you can use variables-- that are initialized before server start
// so you can win some time on serving.
// You can do it else where as well but I let them as pracise for you,
// essentially you can understand by just looking below.
var userLoginView = mvc.View{
	Name: PathLogin.Path + ".html",
	Data: page{"User Login"},
}

// GetLogin handles GET:/user/login.
func (c *Controller) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		return c.logout()
	}
	return userLoginView
}

// PostLogin handles POST:/user/login.
func (c *Controller) PostLogin(form formValue) mvc.Result {
	var (
		username = form("username")
		password = form("password")
	)

	user, err := c.verify(username, password)
	if err != nil {
		return c.fireError(err)
	}

	c.Session.Set(sessionIDKey, user.ID)
	return pathMyProfile
}

// AnyLogout handles any method on path /user/logout.
func (c *Controller) AnyLogout() {
	c.logout()
}

// GetMe handles GET:/user/me.
func (c *Controller) GetMe() mvc.Result {
	id, err := c.Session.GetInt64(sessionIDKey)
	if err != nil || id <= 0 {
		// when not already logged in, redirect to login.
		return PathLogin
	}

	u, found := c.Source.GetByID(id)
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout him and redirect to the register page.
		return c.logout()
	}

	// set the model and render the view template.
	return mvc.View{
		Name: pathMyProfile.Path + ".html",
		Data: iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

func (c *Controller) renderNotFound(id int64) mvc.View {
	return mvc.View{
		Code: iris.StatusNotFound,
		Name: "user/notfound.html",
		Data: iris.Map{
			"Title": "User Not Found",
			"ID":    id,
		},
	}
}

// Dispatch completes the `mvc.Result` interface
// in order to be able to return a type of `Model`
// as mvc.Result.
// If this function didn't exist then
// we should explicit set the output result to that Model or to an interface{}.
func (u Model) Dispatch(ctx iris.Context) {
	ctx.JSON(u)
}

// GetBy handles GET:/user/{id:long},
// i.e http://localhost:8080/user/1
func (c *Controller) GetBy(userID int64) mvc.Result {
	// we have /user/{id}
	// fetch and render user json.
	user, found := c.Source.GetByID(userID)
	if !found {
		// not user found with that ID.
		return c.renderNotFound(userID)
	}

	// Q: how the hell Model can be return as mvc.Result?
	// A: I told you before on some comments and the docs,
	// any struct that has a `Dispatch(ctx iris.Context)`
	// can be returned as an mvc.Result(see ~20 lines above),
	// therefore we are able to combine many type of results in the same method.
	// For example, here, we return either an mvc.View to render a not found custom template
	// either a user which returns the Model as JSON via its Dispatch.
	//
	// We could also return just a struct value that is not an mvc.Result,
	// if the output result of the `GetBy` was that struct's type or an interface{}
	// and iris would render that with JSON as well, but here we can't do that without complete the `Dispatch`
	// function, because we may return an mvc.View which is an mvc.Result.
	return user
}

File: structuring/login-mvc-single-responsibility-package/user/datasource.go

package user

import (
	"errors"
	"sync"
	"time"
)

// IDGenerator would be our user ID generator
// but here we keep the order of users by their IDs
// so we will use numbers that can be easly written
// to the browser to get results back from the REST API.
// var IDGenerator = func() string {
// 	return uuid.NewV4().String()
// }

// DataSource is our data store example.
type DataSource struct {
	Users map[int64]Model
	mu    sync.RWMutex
}

// NewDataSource returns a new user data source.
func NewDataSource() *DataSource {
	return &DataSource{
		Users: make(map[int64]Model),
	}
}

// GetBy receives a query function
// which is fired for every single user model inside
// our imaginary database.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known boolean value
// and the last known user model
// to help callers to reduce the loc.
//
// But be carefully, the caller should always check for the "found"
// because it may be false but the user model has actually real data inside it.
//
// It's actually a simple but very clever prototype function
// I'm think of and using everywhere since then,
// hope you find it very useful too.
func (d *DataSource) GetBy(query func(Model) bool) (user Model, found bool) {
	d.mu.RLock()
	for _, user = range d.Users {
		found = query(user)
		if found {
			break
		}
	}
	d.mu.RUnlock()
	return
}

// GetByID returns a user model based on its ID.
func (d *DataSource) GetByID(id int64) (Model, bool) {
	return d.GetBy(func(u Model) bool {
		return u.ID == id
	})
}

// GetByUsername returns a user model based on the Username.
func (d *DataSource) GetByUsername(username string) (Model, bool) {
	return d.GetBy(func(u Model) bool {
		return u.Username == username
	})
}

func (d *DataSource) getLastID() (lastID int64) {
	d.mu.RLock()
	for id := range d.Users {
		if id > lastID {
			lastID = id
		}
	}
	d.mu.RUnlock()

	return lastID
}

// InsertOrUpdate adds or updates a user to the (memory) storage.
func (d *DataSource) InsertOrUpdate(user Model) (Model, error) {
	// no matter what we will update the password hash
	// for both update and insert actions.
	hashedPassword, err := GeneratePassword(user.password)
	if err != nil {
		return user, err
	}
	user.HashedPassword = hashedPassword

	// update
	if id := user.ID; id > 0 {
		_, found := d.GetByID(id)
		if !found {
			return user, errors.New("ID should be zero or a valid one that maps to an existing User")
		}
		d.mu.Lock()
		d.Users[id] = user
		d.mu.Unlock()
		return user, nil
	}

	// insert
	id := d.getLastID() + 1
	user.ID = id
	d.mu.Lock()
	user.CreatedAt = time.Now()
	d.Users[id] = user
	d.mu.Unlock()

	return user, nil
}

File: structuring/login-mvc-single-responsibility-package/user/model.go

package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Model is our User example model.
type Model struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Username  string `json:"username"`
	// password is the client-given password
	// which will not be stored anywhere in the server.
	// It's here only for actions like registration and update password,
	// because we caccept a Model instance
	// inside the `DataSource#InsertOrUpdate` function.
	password       string
	HashedPassword []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

File: structuring/login-mvc-single-responsibility-package/views/shared/error.html

<h1>Error.</h1>
<h2>An error occurred while processing your request.</h2>

<h3>{{.Message}}</h3>

File: structuring/login-mvc-single-responsibility-package/views/shared/layout.html

<html>

<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" type="text/css" href="/public/css/site.css" />
</head>

<body>
    {{ yield }}
</body>

</html>

File: structuring/login-mvc-single-responsibility-package/views/user/login.html

<form action="/user/login" method="POST">
    <div class="container">
        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="username" required>

        <label><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="password" required>

        <button type="submit">Login</button>
    </div>
</form>

File: structuring/login-mvc-single-responsibility-package/views/user/me.html

<p>
    Welcome back <strong>{{.User.Firstname}}</strong>!
</p>

File: structuring/login-mvc-single-responsibility-package/views/user/notfound.html

<p>
    User with ID <strong>{{.ID}}</strong> does not exist.
</p>

File: structuring/login-mvc-single-responsibility-package/views/user/register.html

<form action="/user/register" method="POST">
    <div class="container">
        <label><b>Firstname</b></label>
        <input type="text" placeholder="Enter Firstname" name="firstname" required>

        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="username" required>

        <label><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="password" required>

        <button type="submit">Register</button>
    </div>
</form>

Subdomains

Multi

File: subdomains/multi/hosts

# Copyright (c) 1993-2009 Microsoft Corp.

#

# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.

#

# This file contains the mappings of IP addresses to host names. Each

# entry should be kept on an individual line. The IP address should

# be placed in the first column followed by the corresponding host name.

# The IP address and the host name should be separated by at least one

# space.

#

# Additionally, comments (such as these) may be inserted on individual

# lines or following the machine name denoted by a '#' symbol.

#

# For example:

#

#      102.54.94.97     rhino.acme.com          # source server

#       38.25.63.10     x.acme.com              # x client host



# localhost name resolution is handled within DNS itself.

127.0.0.1       localhost

::1             localhost

#-iris-For development machine, you have to configure your dns also for online, search google how to do it if you don't know



127.0.0.1 domain.local

127.0.0.1 system.domain.local

127.0.0.1 dashboard.domain.local



#-END iris-

File: subdomains/multi/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	/*
	 * Setup static files
	 */

	app.StaticWeb("/assets", "./public/assets")
	app.StaticWeb("/upload_resources", "./public/upload_resources")

	dashboard := app.Party("dashboard.")
	{
		dashboard.Get("/", func(ctx iris.Context) {
			ctx.Writef("HEY FROM dashboard")
		})
	}
	system := app.Party("system.")
	{
		system.Get("/", func(ctx iris.Context) {
			ctx.Writef("HEY FROM system")
		})
	}

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("HEY FROM frontend /")
	})
	// http://domain.local:80
	// http://dashboard.local
	// http://system.local
	// Make sure you prepend the "http" in your browser
	// because .local is a virtual domain we think to show case you
	// that you can declare any syntactical correct name as a subdomain in iris.
	app.Run(iris.Addr("domain.local:80")) // for beginners: look ../hosts file
}

Redirect

File: subdomains/redirect/hosts

127.0.0.1	mydomain.com
127.0.0.1	www.mydomain.com

# Windows: Drive:/Windows/system32/drivers/etc/hosts, on Linux: /etc/hosts

File: subdomains/redirect/main.go

// Package main shows how to register a simple 'www' subdomain,
// using the `app.WWW` method, which will register a router wrapper which will
// redirect all 'mydomain.com' requests to 'www.mydomain.com'.
// Check the 'hosts' file to see how to test the 'mydomain.com' on your local machine.
package main

import "github.com/kataras/iris"

const addr = "mydomain.com:80"

func main() {
	app := newApp()

	// http(s)://mydomain.com, will be redirect to http(s)://www.mydomain.com.
	// The `www` variable is the `app.Subdomain("www")`.
	//
	// app.WWW() wraps the router so it can redirect all incoming requests
	// that comes from 'http(s)://mydomain.com/%path%' (www is missing)
	// to `http(s)://www.mydomain.com/%path%`.
	//
	// Try:
	// http://mydomain.com             -> http://www.mydomain.com
	// http://mydomain.com/users       -> http://www.mydomain.com/users
	// http://mydomain.com/users/login -> http://www.mydomain.com/users/login
	app.Run(iris.Addr(addr))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("This will never be executed.")
	})

	www := app.Subdomain("www") // <- same as app.Party("www.")
	www.Get("/", index)

	// www is an `iris.Party`, use it like you already know, like grouping routes.
	www.PartyFunc("/users", func(p iris.Party) { // <- same as www.Party("/users").Get(...)
		p.Get("/", usersIndex)
		p.Get("/login", getLogin)
	})

	// redirects mydomain.com/%anypath% to www.mydomain.com/%anypath%.
	// First argument is the 'from' and second is the 'to/target'.
	app.SubdomainRedirect(app, www)

	// SubdomainRedirect works for multi-level subdomains as well:
	// subsub := www.Subdomain("subsub") // subsub.www.mydomain.com
	// subsub.Get("/", func(ctx iris.Context) { ctx.Writef("subdomain is: " + ctx.Subdomain()) })
	// app.SubdomainRedirect(subsub, www)
	//
	// If you need to redirect any subdomain to 'www' then:
	// app.SubdomainRedirect(app.WildcardSubdomain(), www)
	// If you need to redirect from a subdomain to the root domain then:
	// app.SubdomainRedirect(app.Subdomain("mysubdomain"), app)
	//
	// Note that app.Party("mysubdomain.") and app.Subdomain("mysubdomain")
	// is the same exactly thing, the difference is that the second can omit the last dot('.').

	return app
}

func index(ctx iris.Context) {
	ctx.Writef("This is the www.mydomain.com endpoint.")
}

func usersIndex(ctx iris.Context) {
	ctx.Writef("This is the www.mydomain.com/users endpoint.")
}

func getLogin(ctx iris.Context) {
	ctx.Writef("This is the www.mydomain.com/users/login endpoint.")
}

File: subdomains/redirect/main_test.go

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestSubdomainRedirectWWW(t *testing.T) {
	app := newApp()
	root := strings.TrimSuffix(addr, ":80")

	e := httptest.New(t, app)

	tests := []struct {
		path     string
		response string
	}{
		{"/", fmt.Sprintf("This is the www.%s endpoint.", root)},
		{"/users", fmt.Sprintf("This is the www.%s/users endpoint.", root)},
		{"/users/login", fmt.Sprintf("This is the www.%s/users/login endpoint.", root)},
	}

	for _, test := range tests {
		e.GET(test.path).Expect().Status(httptest.StatusOK).Body().Equal(test.response)
	}

}

Single

File: subdomains/single/hosts

# Copyright (c) 1993-2009 Microsoft Corp.
#
# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.
#
# This file contains the mappings of IP addresses to host names. Each
# entry should be kept on an individual line. The IP address should
# be placed in the first column followed by the corresponding host name.
# The IP address and the host name should be separated by at least one
# space.
#
# Additionally, comments (such as these) may be inserted on individual
# lines or following the machine name denoted by a '#' symbol.
#
# For example:
#
#      102.54.94.97     rhino.acme.com          # source server
#       38.25.63.10     x.acme.com              # x client host

# localhost name resolution is handled within DNS itself.
127.0.0.1       localhost
::1             localhost

#-iris-For development machine, you have to configure your dns also for online, search google how to do it if you don't know
127.0.0.1		mydomain.com
127.0.0.1		admin.mydomain.com

#-END iris-

File: subdomains/single/main.go

// Package main register static subdomains, simple as parties, check ./hosts if you use windows
package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// subdomains works with all available routers, like other features too.

	// no order, you can register subdomains at the end also.
	admin := app.Party("admin.")
	{
		// admin.mydomain.com
		admin.Get("/", func(c iris.Context) {
			c.Writef("INDEX FROM admin.mydomain.com")
		})
		// admin.mydomain.com/hey
		admin.Get("/hey", func(c iris.Context) {
			c.Writef("HEY FROM admin.mydomain.com/hey")
		})
		// admin.mydomain.com/hey2
		admin.Get("/hey2", func(c iris.Context) {
			c.Writef("HEY SECOND FROM admin.mydomain.com/hey")
		})
	}

	// mydomain.com/
	app.Get("/", func(c iris.Context) {
		c.Writef("INDEX FROM no-subdomain hey")
	})

	// mydomain.com/hey
	app.Get("/hey", func(c iris.Context) {
		c.Writef("HEY FROM no-subdomain hey")
	})

	// http://admin.mydomain.com
	// http://admin.mydomain.com/hey
	// http://admin.mydomain.com/hey2
	// http://mydomain.com
	// http://mydomain.com/hey
	app.Run(iris.Addr("mydomain.com:80")) // for beginners: look ../hosts file
}

Wildcard

File: subdomains/wildcard/hosts

# Copyright (c) 1993-2009 Microsoft Corp.
#
# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.
#
# This file contains the mappings of IP addresses to host names. Each
# entry should be kept on an individual line. The IP address should
# be placed in the first column followed by the corresponding host name.
# The IP address and the host name should be separated by at least one
# space.
#
# Additionally, comments (such as these) may be inserted on individual
# lines or following the machine name denoted by a '#' symbol.
#
# For example:
#
#      102.54.94.97     rhino.acme.com          # source server
#       38.25.63.10     x.acme.com              # x client host

# localhost name resolution is handled within DNS itself.
127.0.0.1       localhost
::1             localhost

#-iris-For development machine, you have to configure your dns also for online, search google how to do it if you don't know
127.0.0.1		mydomain.com
127.0.0.1		username1.mydomain.com
127.0.0.1		username2.mydomain.com
127.0.0.1		username3.mydomain.com
127.0.0.1		username4.mydomain.com
127.0.0.1		username5.mydomain.com

#-END iris-

File: subdomains/wildcard/main.go

// Package main an example on how to catch dynamic subdomains - wildcard.
// On the first example (subdomains_1) we saw how to create routes for static subdomains, subdomains you know that you will have.
// Here we will see an example how to catch unknown subdomains, dynamic subdomains, like username.mydomain.com:8080.
package main

import (
	"github.com/kataras/iris"
)

// register a dynamic-wildcard subdomain to your server machine(dns/...) first, check ./hosts if you use windows.
// run this file and try to redirect: http://username1.mydomain.com:8080/ , http://username2.mydomain.com:8080/ , http://username1.mydomain.com/something, http://username1.mydomain.com/something/sadsadsa

func main() {
	app := iris.New()

	/* Keep note that you can use both type of subdomains (named and wildcard(*.) )
	   admin.mydomain.com,  and for other the Party(*.) but this is not this example's purpose

	admin := app.Party("admin.")
	{
		// admin.mydomain.com
		admin.Get("/", func(ctx iris.Context) {
			ctx.Writef("INDEX FROM admin.mydomain.com")
		})
		// admin.mydomain.com/hey
		admin.Get("/hey", func(ctx iris.Context) {
			ctx.Writef("HEY FROM admin.mydomain.com/hey")
		})
		// admin.mydomain.com/hey2
		admin.Get("/hey2", func(ctx iris.Context) {
			ctx.Writef("HEY SECOND FROM admin.mydomain.com/hey")
		})
	}*/

	// no order, you can register subdomains at the end also.
	dynamicSubdomains := app.Party("*.")
	{
		dynamicSubdomains.Get("/", dynamicSubdomainHandler)

		dynamicSubdomains.Get("/something", dynamicSubdomainHandler)

		dynamicSubdomains.Get("/something/{paramfirst}", dynamicSubdomainHandlerWithParam)
	}

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from mydomain.com path: %s", ctx.Path())
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.Writef("Hello from mydomain.com path: %s", ctx.Path())
	})

	// http://mydomain.com:8080
	// http://username1.mydomain.com:8080
	// http://username2.mydomain.com:8080/something
	// http://username3.mydomain.com:8080/something/yourname
	app.Run(iris.Addr("mydomain.com:8080")) // for beginners: look ../hosts file
}

func dynamicSubdomainHandler(ctx iris.Context) {
	username := ctx.Subdomain()
	ctx.Writef("Hello from dynamic subdomain path: %s, here you can handle the route for dynamic subdomains, handle the user: %s", ctx.Path(), username)
	// if  http://username4.mydomain.com:8080/ prints:
	// Hello from dynamic subdomain path: /, here you can handle the route for dynamic subdomains, handle the user: username4
}

func dynamicSubdomainHandlerWithParam(ctx iris.Context) {
	username := ctx.Subdomain()
	ctx.Writef("Hello from dynamic subdomain path: %s, here you can handle the route for dynamic subdomains, handle the user: %s", ctx.Path(), username)
	ctx.Writef("The paramfirst is: %s", ctx.Params().Get("paramfirst"))
}

WWW

File: subdomains/www/hosts

# Copyright (c) 1993-2009 Microsoft Corp.
#
# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.
#
# This file contains the mappings of IP addresses to host names. Each
# entry should be kept on an individual line. The IP address should
# be placed in the first column followed by the corresponding host name.
# The IP address and the host name should be separated by at least one
# space.
#
# Additionally, comments (such as these) may be inserted on individual
# lines or following the machine name denoted by a '#' symbol.
#
# For example:
#
#      102.54.94.97     rhino.acme.com          # source server
#       38.25.63.10     x.acme.com              # x client host

# localhost name resolution is handled within DNS itself.
127.0.0.1       localhost
::1             localhost

#-iris-For development machine, you have to configure your dns also for online, search google how to do it if you don't know
127.0.0.1		mydomain.com
127.0.0.1		www.mydomain.com
#-END iris-

File: subdomains/www/main.go

package main

import (
	"github.com/kataras/iris"
)

func newApp() *iris.Application {
	app := iris.New()

	app.Get("/", info)
	app.Get("/about", info)
	app.Get("/contact", info)

	app.PartyFunc("/api/users", func(r iris.Party) {
		r.Get("/", info)
		r.Get("/{id:int}", info)

		r.Post("/", info)

		r.Put("/{id:int}", info)
	}) /* <- same as:
	 usersAPI := app.Party("/api/users")
	 {  // those brackets are just syntactic-sugar things.
		// This method is rarely used but you can make use of it when you want
	    // scoped variables to that code block only.
		usersAPI.Get/Post...
	 }
	 usersAPI.Get/Post...
	*/

	www := app.Party("www.")
	{
		// Just to show how you can get all routes and copy them to another
		// party or subdomain:
		// Get all routes that are registered so far, including all "Parties" and subdomains:
		currentRoutes := app.GetRoutes()
		// Register them to the www subdomain/vhost as well:
		for _, r := range currentRoutes {
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}

		// http://www.mydomain.com/hi
		www.Get("/hi", func(ctx iris.Context) {
			ctx.Writef("hi from www.mydomain.com")
		})
	}
	// See also the "subdomains/redirect" to register redirect router wrappers between subdomains,
	// i.e mydomain.com to www.mydomain.com (like facebook does for SEO reasons(;)).

	return app
}

func main() {
	app := newApp()
	// http://mydomain.com
	// http://mydomain.com/about
	// http://imydomain.com/contact
	// http://mydomain.com/api/users
	// http://mydomain.com/api/users/42

	// http://www.mydomain.com
	// http://www.mydomain.com/hi
	// http://www.mydomain.com/about
	// http://www.mydomain.com/contact
	// http://www.mydomain.com/api/users
	// http://www.mydomain.com/api/users/42
	if err := app.Run(iris.Addr("mydomain.com:80")); err != nil {
		panic(err)
	}
}

func info(ctx iris.Context) {
	method := ctx.Method()
	subdomain := ctx.Subdomain()
	path := ctx.Path()

	ctx.Writef("\nInfo\n\n")
	ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s", method, subdomain, path)
}

File: subdomains/www/main_test.go

package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

type testRoute struct {
	path      string
	method    string
	subdomain string
}

func (r testRoute) response() string {
	msg := fmt.Sprintf("\nInfo\n\nMethod: %s\nSubdomain: %s\nPath: %s", r.method, r.subdomain, r.path)
	return msg
}

func TestSubdomainWWW(t *testing.T) {
	app := newApp()

	tests := []testRoute{
		// host
		{"/", "GET", ""},
		{"/about", "GET", ""},
		{"/contact", "GET", ""},
		{"/api/users", "GET", ""},
		{"/api/users/42", "GET", ""},
		{"/api/users", "POST", ""},
		{"/api/users/42", "PUT", ""},
		// www sub domain
		{"/", "GET", "www"},
		{"/about", "GET", "www"},
		{"/contact", "GET", "www"},
		{"/api/users", "GET", "www"},
		{"/api/users/42", "GET", "www"},
		{"/api/users", "POST", "www"},
		{"/api/users/42", "PUT", "www"},
	}

	host := "localhost:1111"
	e := httptest.New(t, app, httptest.URL("http://"+host), httptest.Debug(false))

	for _, test := range tests {

		req := e.Request(test.method, test.path)
		if subdomain := test.subdomain; subdomain != "" {
			req.WithURL("http://" + subdomain + "." + host)
		}

		req.Expect().
			Status(httptest.StatusOK).
			Body().Equal(test.response())
	}

}

Testing

HTTPtest

File: testing/httptest/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		Users: map[string]string{"myusername": "mypassword"},
	}

	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })

	// to party

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}

	return app
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter it will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

File: testing/httptest/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

// $ go test -v
func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	// redirects to /admin without basic auth
	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	// without basic auth
	e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)

	// with valid basic auth
	e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
	e.GET("/admin/profile").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/profile myusername:mypassword")
	e.GET("/admin/settings").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/settings myusername:mypassword")

	// with invalid basic auth
	e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
		Expect().Status(httptest.StatusUnauthorized)

}

Tutorial

Caddy

File: tutorial/caddy/Caddyfile

example.com {

	header / Server "Iris"

	proxy / example.com:9091 # localhost:9091

}



api.example.com {

	header / Server "Iris"

	proxy / api.example.com:9092 # localhost:9092

}

File: tutorial/caddy/server1/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	templates := iris.HTML("./views", ".html").Layout("shared/layout.html")
	app.RegisterView(templates)

	mvc.New(app).Handle(new(Controller))

	// http://localhost:9091
	app.Run(iris.Addr(":9091"))
}

// Layout contains all the binding properties for the shared/layout.html
type Layout struct {
	Title string
}

// Controller is our example controller, request-scoped, each request has its own instance.
type Controller struct {
	Layout Layout
}

// BeginRequest is the first method fired when client requests from this Controller's root path.
func (c *Controller) BeginRequest(ctx iris.Context) {
	c.Layout.Title = "Home Page"
}

// EndRequest is the last method fired.
// It's here just to complete the BaseController
// in order to be tell iris to call the `BeginRequest` before the main method.
func (c *Controller) EndRequest(ctx iris.Context) {}

// Get handles GET http://localhost:9091
func (c *Controller) Get() mvc.View {
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Layout":  c.Layout,
			"Message": "Welcome to my website!",
		},
	}
}

File: tutorial/caddy/server1/views/index.html

<div>
    {{.Message}}
</div>

File: tutorial/caddy/server1/views/shared/layout.html

<html>

<head>
    <title>{{.Layout.Title}}</title>
</head>

<body>
    {{ yield }}
</body>

</html>

File: tutorial/caddy/server2/main.go

package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type postValue func(string) string

func main() {
	app := iris.New()

	mvc.New(app.Party("/user")).Register(
		func(ctx iris.Context) postValue {
			return ctx.PostValue
		}).Handle(new(UserController))

	// GET http://localhost:9092/user
	// GET http://localhost:9092/user/42
	// POST http://localhost:9092/user
	// PUT http://localhost:9092/user/42
	// DELETE http://localhost:9092/user/42
	// GET http://localhost:9092/user/followers/42
	app.Run(iris.Addr(":9092"))
}

// UserController is our user example controller.
type UserController struct{}

// Get handles GET /user
func (c *UserController) Get() string {
	return "Select all users"
}

// User is our test User model, nothing tremendous here.
type User struct{ ID int64 }

// GetBy handles GET /user/42, equal to .Get("/user/{id:long}")
func (c *UserController) GetBy(id int64) User {
	// Select User by ID == $id.
	return User{id}
}

// Post handles POST /user
func (c *UserController) Post(post postValue) string {
	username := post("username")
	return "Create by user with username: " + username
}

// PutBy handles PUT /user/42
func (c *UserController) PutBy(id int) string {
	// Update user by ID == $id
	return "User updated"
}

// DeleteBy handles DELETE /user/42
func (c *UserController) DeleteBy(id int) bool {
	// Delete user by ID == %id
	//
	// when boolean then true = iris.StatusOK, false = iris.StatusNotFound
	return true
}

// GetFollowersBy handles GET /user/followers/42
func (c *UserController) GetFollowersBy(id int) []User {
	// Select all followers by user ID == $id
	return []User{ /* ... */ }
}

DropzoneJS

File: tutorial/dropzonejs/no_files.png
File: tutorial/dropzonejs/src/main.go

package main

import (
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kataras/iris"

	"github.com/nfnt/resize"
)

// $ go get -u github.com/nfnt/resize

const uploadsDir = "./public/uploads/"

type uploadedFile struct {
	// {name: "", size: } are the dropzone's only requirements.
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type uploadedFiles struct {
	dir   string
	items []uploadedFile
	mu    sync.RWMutex // slices are safe but RWMutex is a good practise for you.
}

func scanUploads(dir string) *uploadedFiles {

	f := new(uploadedFiles)

	lindex := dir[len(dir)-1]
	if lindex != os.PathSeparator && lindex != '/' {
		dir += string(os.PathSeparator)
	}

	// create directories if necessary
	// and if, then return empty uploaded files; skipping the scan.
	if err := os.MkdirAll(dir, os.FileMode(0666)); err != nil {
		return f
	}

	// otherwise scan the given "dir" for files.
	f.scan(dir)
	return f
}

func (f *uploadedFiles) scan(dir string) {
	f.dir = dir
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		// if it's directory or a thumbnail we saved earlier, skip it.
		if info.IsDir() || strings.HasPrefix(info.Name(), "thumbnail_") {
			return nil
		}

		f.add(info.Name(), info.Size())
		return nil
	})
}

func (f *uploadedFiles) add(name string, size int64) uploadedFile {
	uf := uploadedFile{
		Name: name,
		Size: size,
	}

	f.mu.Lock()
	f.items = append(f.items, uf)
	f.mu.Unlock()

	return uf
}

func (f *uploadedFiles) createThumbnail(uf uploadedFile) {
	file, err := os.Open(path.Join(f.dir, uf.Name))
	if err != nil {
		return
	}
	defer file.Close()

	name := strings.ToLower(uf.Name)

	out, err := os.OpenFile(f.dir+"thumbnail_"+uf.Name,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer out.Close()

	if strings.HasSuffix(name, ".jpg") {
		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			return
		}

		// write new image to file
		resized := resize.Thumbnail(180, 180, img, resize.Lanczos3)
		jpeg.Encode(out, resized,
			&jpeg.Options{Quality: jpeg.DefaultQuality})

	} else if strings.HasSuffix(name, ".png") {
		img, err := png.Decode(file)
		if err != nil {
			return
		}

		// write new image to file
		resized := resize.Thumbnail(180, 180, img, resize.Lanczos3) // slower but better res
		png.Encode(out, resized)
	}
	// and so on... you got the point, this code can be simplify, as a practise.

}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.StaticWeb("/public", "./public")

	app.Get("/", func(ctx iris.Context) {
		ctx.View("upload.html")
	})

	files := scanUploads(uploadsDir)

	app.Get("/uploads", func(ctx iris.Context) {
		ctx.JSON(files.items)
	})

	app.Post("/upload", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		// Get the file from the dropzone request
		file, info, err := ctx.FormFile("file")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
			return
		}

		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile(uploadsDir+fname,
			os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
			return
		}
		defer out.Close()

		io.Copy(out, file)

		// optionally, add that file to the list in order to be visible when refresh.
		uploadedFile := files.add(fname, info.Size)
		go files.createThumbnail(uploadedFile)
	})

	// start the server at http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

File: tutorial/dropzonejs/src/public/css/dropzone.css

/*
 * The MIT License
 * Copyright (c) 2012 Matias Meno <m@tias.me>
 */
@-webkit-keyframes passing-through {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30%, 70% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); }
  100% {
    opacity: 0;
    -webkit-transform: translateY(-40px);
    -moz-transform: translateY(-40px);
    -ms-transform: translateY(-40px);
    -o-transform: translateY(-40px);
    transform: translateY(-40px); } }
@-moz-keyframes passing-through {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30%, 70% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); }
  100% {
    opacity: 0;
    -webkit-transform: translateY(-40px);
    -moz-transform: translateY(-40px);
    -ms-transform: translateY(-40px);
    -o-transform: translateY(-40px);
    transform: translateY(-40px); } }
@keyframes passing-through {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30%, 70% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); }
  100% {
    opacity: 0;
    -webkit-transform: translateY(-40px);
    -moz-transform: translateY(-40px);
    -ms-transform: translateY(-40px);
    -o-transform: translateY(-40px);
    transform: translateY(-40px); } }
@-webkit-keyframes slide-in {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); } }
@-moz-keyframes slide-in {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); } }
@keyframes slide-in {
  0% {
    opacity: 0;
    -webkit-transform: translateY(40px);
    -moz-transform: translateY(40px);
    -ms-transform: translateY(40px);
    -o-transform: translateY(40px);
    transform: translateY(40px); }
  30% {
    opacity: 1;
    -webkit-transform: translateY(0px);
    -moz-transform: translateY(0px);
    -ms-transform: translateY(0px);
    -o-transform: translateY(0px);
    transform: translateY(0px); } }
@-webkit-keyframes pulse {
  0% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); }
  10% {
    -webkit-transform: scale(1.1);
    -moz-transform: scale(1.1);
    -ms-transform: scale(1.1);
    -o-transform: scale(1.1);
    transform: scale(1.1); }
  20% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); } }
@-moz-keyframes pulse {
  0% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); }
  10% {
    -webkit-transform: scale(1.1);
    -moz-transform: scale(1.1);
    -ms-transform: scale(1.1);
    -o-transform: scale(1.1);
    transform: scale(1.1); }
  20% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); } }
@keyframes pulse {
  0% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); }
  10% {
    -webkit-transform: scale(1.1);
    -moz-transform: scale(1.1);
    -ms-transform: scale(1.1);
    -o-transform: scale(1.1);
    transform: scale(1.1); }
  20% {
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    -o-transform: scale(1);
    transform: scale(1); } }
.dropzone, .dropzone * {
  box-sizing: border-box; }

.dropzone {
  min-height: 150px;
  border: 2px solid rgba(0, 0, 0, 0.3);
  background: white;
  padding: 20px 20px; }
  .dropzone.dz-clickable {
    cursor: pointer; }
    .dropzone.dz-clickable * {
      cursor: default; }
    .dropzone.dz-clickable .dz-message, .dropzone.dz-clickable .dz-message * {
      cursor: pointer; }
  .dropzone.dz-started .dz-message {
    display: none; }
  .dropzone.dz-drag-hover {
    border-style: solid; }
    .dropzone.dz-drag-hover .dz-message {
      opacity: 0.5; }
  .dropzone .dz-message {
    text-align: center;
    margin: 2em 0; }
  .dropzone .dz-preview {
    position: relative;
    display: inline-block;
    vertical-align: top;
    margin: 16px;
    min-height: 100px; }
    .dropzone .dz-preview:hover {
      z-index: 1000; }
      .dropzone .dz-preview:hover .dz-details {
        opacity: 1; }
    .dropzone .dz-preview.dz-file-preview .dz-image {
      border-radius: 20px;
      background: #999;
      background: linear-gradient(to bottom, #eee, #ddd); }
    .dropzone .dz-preview.dz-file-preview .dz-details {
      opacity: 1; }
    .dropzone .dz-preview.dz-image-preview {
      background: white; }
      .dropzone .dz-preview.dz-image-preview .dz-details {
        -webkit-transition: opacity 0.2s linear;
        -moz-transition: opacity 0.2s linear;
        -ms-transition: opacity 0.2s linear;
        -o-transition: opacity 0.2s linear;
        transition: opacity 0.2s linear; }
    .dropzone .dz-preview .dz-remove {
      font-size: 14px;
      text-align: center;
      display: block;
      cursor: pointer;
      border: none; }
      .dropzone .dz-preview .dz-remove:hover {
        text-decoration: underline; }
    .dropzone .dz-preview:hover .dz-details {
      opacity: 1; }
    .dropzone .dz-preview .dz-details {
      z-index: 20;
      position: absolute;
      top: 0;
      left: 0;
      opacity: 0;
      font-size: 13px;
      min-width: 100%;
      max-width: 100%;
      padding: 2em 1em;
      text-align: center;
      color: rgba(0, 0, 0, 0.9);
      line-height: 150%; }
      .dropzone .dz-preview .dz-details .dz-size {
        margin-bottom: 1em;
        font-size: 16px; }
      .dropzone .dz-preview .dz-details .dz-filename {
        white-space: nowrap; }
        .dropzone .dz-preview .dz-details .dz-filename:hover span {
          border: 1px solid rgba(200, 200, 200, 0.8);
          background-color: rgba(255, 255, 255, 0.8); }
        .dropzone .dz-preview .dz-details .dz-filename:not(:hover) {
          overflow: hidden;
          text-overflow: ellipsis; }
          .dropzone .dz-preview .dz-details .dz-filename:not(:hover) span {
            border: 1px solid transparent; }
      .dropzone .dz-preview .dz-details .dz-filename span, .dropzone .dz-preview .dz-details .dz-size span {
        background-color: rgba(255, 255, 255, 0.4);
        padding: 0 0.4em;
        border-radius: 3px; }
    .dropzone .dz-preview:hover .dz-image img {
      -webkit-transform: scale(1.05, 1.05);
      -moz-transform: scale(1.05, 1.05);
      -ms-transform: scale(1.05, 1.05);
      -o-transform: scale(1.05, 1.05);
      transform: scale(1.05, 1.05);
      -webkit-filter: blur(8px);
      filter: blur(8px); }
    .dropzone .dz-preview .dz-image {
      border-radius: 20px;
      overflow: hidden;
      width: 120px;
      height: 120px;
      position: relative;
      display: block;
      z-index: 10; }
      .dropzone .dz-preview .dz-image img {
        display: block; }
    .dropzone .dz-preview.dz-success .dz-success-mark {
      -webkit-animation: passing-through 3s cubic-bezier(0.77, 0, 0.175, 1);
      -moz-animation: passing-through 3s cubic-bezier(0.77, 0, 0.175, 1);
      -ms-animation: passing-through 3s cubic-bezier(0.77, 0, 0.175, 1);
      -o-animation: passing-through 3s cubic-bezier(0.77, 0, 0.175, 1);
      animation: passing-through 3s cubic-bezier(0.77, 0, 0.175, 1); }
    .dropzone .dz-preview.dz-error .dz-error-mark {
      opacity: 1;
      -webkit-animation: slide-in 3s cubic-bezier(0.77, 0, 0.175, 1);
      -moz-animation: slide-in 3s cubic-bezier(0.77, 0, 0.175, 1);
      -ms-animation: slide-in 3s cubic-bezier(0.77, 0, 0.175, 1);
      -o-animation: slide-in 3s cubic-bezier(0.77, 0, 0.175, 1);
      animation: slide-in 3s cubic-bezier(0.77, 0, 0.175, 1); }
    .dropzone .dz-preview .dz-success-mark, .dropzone .dz-preview .dz-error-mark {
      pointer-events: none;
      opacity: 0;
      z-index: 500;
      position: absolute;
      display: block;
      top: 50%;
      left: 50%;
      margin-left: -27px;
      margin-top: -27px; }
      .dropzone .dz-preview .dz-success-mark svg, .dropzone .dz-preview .dz-error-mark svg {
        display: block;
        width: 54px;
        height: 54px; }
    .dropzone .dz-preview.dz-processing .dz-progress {
      opacity: 1;
      -webkit-transition: all 0.2s linear;
      -moz-transition: all 0.2s linear;
      -ms-transition: all 0.2s linear;
      -o-transition: all 0.2s linear;
      transition: all 0.2s linear; }
    .dropzone .dz-preview.dz-complete .dz-progress {
      opacity: 0;
      -webkit-transition: opacity 0.4s ease-in;
      -moz-transition: opacity 0.4s ease-in;
      -ms-transition: opacity 0.4s ease-in;
      -o-transition: opacity 0.4s ease-in;
      transition: opacity 0.4s ease-in; }
    .dropzone .dz-preview:not(.dz-processing) .dz-progress {
      -webkit-animation: pulse 6s ease infinite;
      -moz-animation: pulse 6s ease infinite;
      -ms-animation: pulse 6s ease infinite;
      -o-animation: pulse 6s ease infinite;
      animation: pulse 6s ease infinite; }
    .dropzone .dz-preview .dz-progress {
      opacity: 1;
      z-index: 1000;
      pointer-events: none;
      position: absolute;
      height: 16px;
      left: 50%;
      top: 50%;
      margin-top: -8px;
      width: 80px;
      margin-left: -40px;
      background: rgba(255, 255, 255, 0.9);
      -webkit-transform: scale(1);
      border-radius: 8px;
      overflow: hidden; }
      .dropzone .dz-preview .dz-progress .dz-upload {
        background: #333;
        background: linear-gradient(to bottom, #666, #444);
        position: absolute;
        top: 0;
        left: 0;
        bottom: 0;
        width: 0;
        -webkit-transition: width 300ms ease-in-out;
        -moz-transition: width 300ms ease-in-out;
        -ms-transition: width 300ms ease-in-out;
        -o-transition: width 300ms ease-in-out;
        transition: width 300ms ease-in-out; }
    .dropzone .dz-preview.dz-error .dz-error-message {
      display: block; }
    .dropzone .dz-preview.dz-error:hover .dz-error-message {
      opacity: 1;
      pointer-events: auto; }
    .dropzone .dz-preview .dz-error-message {
      pointer-events: none;
      z-index: 1000;
      position: absolute;
      display: block;
      display: none;
      opacity: 0;
      -webkit-transition: opacity 0.3s ease;
      -moz-transition: opacity 0.3s ease;
      -ms-transition: opacity 0.3s ease;
      -o-transition: opacity 0.3s ease;
      transition: opacity 0.3s ease;
      border-radius: 8px;
      font-size: 13px;
      top: 130px;
      left: -10px;
      width: 140px;
      background: #be2626;
      background: linear-gradient(to bottom, #be2626, #a92222);
      padding: 0.5em 1.2em;
      color: white; }
      .dropzone .dz-preview .dz-error-message:after {
        content: '';
        position: absolute;
        top: -6px;
        left: 64px;
        width: 0;
        height: 0;
        border-left: 6px solid transparent;
        border-right: 6px solid transparent;
        border-bottom: 6px solid #be2626; }

File: tutorial/dropzonejs/src/views/upload.html

<html>

<head>
    <title>DropzoneJS Uploader</title>

    <!-- 1 -->
    <link href="/public/css/dropzone.css" type="text/css" rel="stylesheet" />

    <!-- 2 -->
    <script src="/public/js/dropzone.js"></script>
    <!-- 4 -->
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <!-- 5 -->
    <script>
        Dropzone.options.myDropzone = {
            paramName: "file", // The name that will be used to transfer the file
            init: function () {
                thisDropzone = this;
                // 6
                $.get('/uploads', function (data) {

                    if (data == null) {
                        return;
                    }
                    // 7
                    $.each(data, function (key, value) {
                        var mockFile = { name: value.name, size: value.size };

                        thisDropzone.emit("addedfile", mockFile);
                        thisDropzone.options.thumbnail.call(thisDropzone, mockFile, '/public/uploads/thumbnail_' + value.name);
                        // thisDropzone.createThumbnailFromUrl(mockFile, '/public/uploads/' + value.name); <- doesn't work...
                        // Make sure that there is no progress bar, etc...
                        thisDropzone.emit("complete", mockFile);
                    });

                });
            }
        };
    </script>
</head>

<body>

    <!-- 3 -->
    <form action="/upload" method="POST" class="dropzone" id="my-dropzone">
        <div class="fallback">
            <input name="file" type="file" multiple />
            <input type="submit" value="Upload" />
        </div>
    </form>
</body>

</html>

File: tutorial/dropzonejs/with_files.png

How to build a file upload form using DropzoneJS and Gowritten by https://twitter.com/@kataras

How to display existing files on server using DropzoneJS and Gowritten by https://twitter.com/@kataras

Online Visitors

File: tutorial/online-visitors/main.go

package main

import (
	"sync/atomic"

	"github.com/kataras/iris"

	"github.com/kataras/iris/websocket"
)

func main() {
	// init the web application instance
	// app := iris.New()
	app := iris.Default()

	// load templates
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))
	// setup the websocket server
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(HandleWebsocketConnection)

	app.Get("/my_endpoint", ws.Handler())
	app.Any("/iris-ws.js", websocket.ClientHandler())

	// register static assets request path and system directory
	app.StaticWeb("/js", "./static/assets/js")

	h := func(ctx iris.Context) {
		ctx.ViewData("", page{PageID: "index page"})
		ctx.View("index.html")
	}

	h2 := func(ctx iris.Context) {
		ctx.ViewData("", page{PageID: "other page"})
		ctx.View("other.html")
	}

	// Open some browser tabs/or windows
	// and navigate to
	// http://localhost:8080/ and http://localhost:8080/other multiple times.
	// Each page has its own online-visitors counter.
	app.Get("/", h)
	app.Get("/other", h2)
	app.Run(iris.Addr(":8080"))
}

type page struct {
	PageID string
}

type pageView struct {
	source string
	count  uint64
}

func (v *pageView) increment() {
	atomic.AddUint64(&v.count, 1)
}

func (v *pageView) decrement() {
	atomic.AddUint64(&v.count, ^uint64(0))
}

func (v *pageView) getCount() uint64 {
	return atomic.LoadUint64(&v.count)
}

type (
	pageViews []pageView
)

func (v *pageViews) Add(source string) {
	args := *v
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.source == source {
			kv.increment()
			return
		}
	}

	c := cap(args)
	if c > n {
		args = args[:n+1]
		kv := &args[n]
		kv.source = source
		kv.count = 1
		*v = args
		return
	}

	kv := pageView{}
	kv.source = source
	kv.count = 1
	*v = append(args, kv)
}

func (v *pageViews) Get(source string) *pageView {
	args := *v
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if kv.source == source {
			return kv
		}
	}
	return nil
}

func (v *pageViews) Reset() {
	*v = (*v)[:0]
}

var v pageViews

// HandleWebsocketConnection handles the online viewers per example(gist source)
func HandleWebsocketConnection(c websocket.Connection) {

	c.On("watch", func(pageSource string) {
		v.Add(pageSource)
		// join the socket to a room linked with the page source
		c.Join(pageSource)

		viewsCount := v.Get(pageSource).getCount()
		if viewsCount == 0 {
			viewsCount++ // count should be always > 0 here
		}
		c.To(pageSource).Emit("watch", viewsCount)
	})

	c.OnLeave(func(roomName string) {
		if roomName != c.ID() { // if the roomName  it's not the connection iself
			// the roomName here is the source, this is the only room(except the connection's ID room) which we join the users to.
			pageV := v.Get(roomName)
			if pageV == nil {
				return // for any case that this room is not a pageView source
			}
			// decrement -1 the specific counter for this page source.
			pageV.decrement()
			// 1. open 30 tabs.
			// 2. close the browser.
			// 3. re-open the browser
			// 4. should be  v.getCount() = 1
			// in order to achieve the previous flow we should decrement exactly when the user disconnects
			// but emit the result a little after, on a goroutine
			// getting all connections within this room and emit the online views one by one.
			// note:
			// we can also add a time.Sleep(2-3 seconds) inside the goroutine at the future if we don't need 'real-time' updates.
			go func(currentConnID string) {
				for _, conn := range c.Server().GetConnectionsByRoom(roomName) {
					if conn.ID() != currentConnID {
						conn.Emit("watch", pageV.getCount())
					}

				}
			}(c.ID())
		}

	})
}

File: tutorial/online-visitors/static/assets/js/visitors.js

(function() {
  var socket = new Ws("ws://localhost:8080/my_endpoint");

  socket.OnConnect(function () {
      socket.Emit("watch", PAGE_SOURCE);
  });


  socket.On("watch", function (onlineViews) {
      var text = "1 online view";
      if (onlineViews > 1) {
          text = onlineViews + " online views";
      }
      document.getElementById("online_views").innerHTML = text;
  });

  socket.OnDisconnect(function () {
    document.getElementById("online_views").innerHTML = "you've been disconnected";
  });

})();

File: tutorial/online-visitors/templates/index.html

<html>

<head>
    <title>Online visitors example</title>
    <style>
        body {
            margin: 0;
            font-family: -apple-system, "San Francisco", "Helvetica Neue", "Noto", "Roboto", "Calibri Light", sans-serif;
            color: #212121;
            font-size: 1.0em;
            line-height: 1.6;
        }

        .container {
            max-width: 750px;
            margin: auto;
            padding: 15px;
        }

        #online_views {
            font-weight: bold;
            font-size: 18px;
        }
    </style>
</head>

<body>
    <div class="container">
        <span id="online_views">1 online view</span>
    </div>

    <script type="text/javascript">
      /* take the page source from our passed struct  on .Render */
      var PAGE_SOURCE = {{ .PageID }}
    </script>

    <script src="/iris-ws.js"></script>

    <script src="/js/visitors.js"></script>

</body>

</html>

File: tutorial/online-visitors/templates/other.html

<html>

<head>
    <title>Different page, different results</title>
    <style>
        #online_views {
            font-weight: bold;
            font-size: 18px;
        }
    </style>
</head>

<body>

    <span id="online_views">1 online view</span>


    <script type="text/javascript">
      /* take the page source from our passed struct  on .Render */
      var PAGE_SOURCE = {{ .PageID }}
    </script>

    <script src="/iris-ws.js"></script>

    <script src="/js/visitors.js"></script>

</body>

</html>

Url Shortener
File: tutorial/url-shortener/factory.go

package main

import (
	"net/url"

	"github.com/satori/go.uuid"
)

// Generator the type to generate keys(short urls)
type Generator func() string

// DefaultGenerator is the defautl url generator
var DefaultGenerator = func() string {
	id, _ := uuid.NewV4()
	return id.String()
}

// Factory is responsible to generate keys(short urls)
type Factory struct {
	store     Store
	generator Generator
}

// NewFactory receives a generator and a store and returns a new url Factory.
func NewFactory(generator Generator, store Store) *Factory {
	return &Factory{
		store:     store,
		generator: generator,
	}
}

// Gen generates the key.
func (f *Factory) Gen(uri string) (key string, err error) {
	// we don't return the parsed url because #hash are converted to uri-compatible
	// and we don't want to encode/decode all the time, there is no need for that,
	// we save the url as the user expects if the uri validation passed.
	_, err = url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}

	key = f.generator()
	// Make sure that the key is unique
	for {
		if v := f.store.Get(key); v == "" {
			break
		}
		key = f.generator()
	}

	return key, nil
}

File: tutorial/url-shortener/main.go

// Package main shows how you can create a simple URL Shortener.
//
// Article: https://medium.com/@kataras/a-url-shortener-service-using-go-iris-and-bolt-4182f0b00ae7
//
// $ go get github.com/coreos/bbolt
// $ go get github.com/satori/go.uuid
// $ cd $GOPATH/src/github.com/kataras/iris/_examples/tutorial/url-shortener
// $ go build
// $ ./url-shortener
package main

import (
	"html/template"

	"github.com/kataras/iris"
)

func main() {
	// assign a variable to the DB so we can use its features later.
	db := NewDB("shortener.db")
	// Pass that db to our app, in order to be able to test the whole app with a different database later on.
	app := newApp(db)

	// release the "db" connection when server goes off.
	iris.RegisterOnInterrupt(db.Close)

	app.Run(iris.Addr(":8080"))
}

func newApp(db *DB) *iris.Application {
	app := iris.Default() // or app := iris.New()

	// create our factory, which is the manager for the object creation.
	// between our web app and the db.
	factory := NewFactory(DefaultGenerator, db)

	// serve the "./templates" directory's "*.html" files with the HTML std view engine.
	tmpl := iris.HTML("./templates", ".html").Reload(true)
	// register any template func(s) here.
	//
	// Look ./templates/index.html#L16
	tmpl.AddFunc("IsPositive", func(n int) bool {
		if n > 0 {
			return true
		}
		return false
	})

	app.RegisterView(tmpl)

	// Serve static files (css)
	app.StaticWeb("/static", "./resources")

	indexHandler := func(ctx iris.Context) {
		ctx.ViewData("URL_COUNT", db.Len())
		ctx.View("index.html")
	}
	app.Get("/", indexHandler)

	// find and execute a short url by its key
	// used on http://localhost:8080/u/dsaoj41u321dsa
	execShortURL := func(ctx iris.Context, key string) {
		if key == "" {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		value := db.Get(key)
		if value == "" {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.Writef("Short URL for key: '%s' not found", key)
			return
		}

		ctx.Redirect(value, iris.StatusTemporaryRedirect)
	}
	app.Get("/u/{shortkey}", func(ctx iris.Context) {
		execShortURL(ctx, ctx.Params().Get("shortkey"))
	})

	app.Post("/shorten", func(ctx iris.Context) {
		formValue := ctx.FormValue("url")
		if formValue == "" {
			ctx.ViewData("FORM_RESULT", "You need to a enter a URL")
			ctx.StatusCode(iris.StatusLengthRequired)
		} else {
			key, err := factory.Gen(formValue)
			if err != nil {
				ctx.ViewData("FORM_RESULT", "Invalid URL")
				ctx.StatusCode(iris.StatusBadRequest)
			} else {
				if err = db.Set(key, formValue); err != nil {
					ctx.ViewData("FORM_RESULT", "Internal error while saving the URL")
					app.Logger().Infof("while saving URL: " + err.Error())
					ctx.StatusCode(iris.StatusInternalServerError)
				} else {
					ctx.StatusCode(iris.StatusOK)
					shortenURL := "http://" + app.ConfigurationReadOnly().GetVHost() + "/u/" + key
					ctx.ViewData("FORM_RESULT",
						template.HTML("<pre><a target='_new' href='"+shortenURL+"'>"+shortenURL+" </a></pre>"))
				}

			}
		}

		indexHandler(ctx) // no redirect, we need the FORM_RESULT.
	})

	app.Post("/clear_cache", func(ctx iris.Context) {
		db.Clear()
		ctx.Redirect("/")
	})

	return app
}

File: tutorial/url-shortener/main_test.go

package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/kataras/iris/httptest"
)

// TestURLShortener tests the simple tasks of our url shortener application.
// Note that it's a pure test.
// The rest possible checks is up to you, take it as as an exercise!
func TestURLShortener(t *testing.T) {
	// temp db file
	f, err := ioutil.TempFile("", "shortener")
	if err != nil {
		t.Fatalf("creating temp file for database failed: %v", err)
	}

	db := NewDB(f.Name())
	app := newApp(db)

	e := httptest.New(t, app)
	originalURL := "https://google.com"

	// save
	e.POST("/shorten").
		WithFormField("url", originalURL).Expect().
		Status(httptest.StatusOK).Body().Contains("<pre><a target='_new' href=")

	keys := db.GetByValue(originalURL)
	if got := len(keys); got != 1 {
		t.Fatalf("expected to have 1 key but saved %d short urls", got)
	}

	// get
	e.GET("/u/" + keys[0]).Expect().
		Status(httptest.StatusTemporaryRedirect).Header("Location").Equal(originalURL)

	// save the same again, it should add a new key
	e.POST("/shorten").
		WithFormField("url", originalURL).Expect().
		Status(httptest.StatusOK).Body().Contains("<pre><a target='_new' href=")

	keys2 := db.GetByValue(originalURL)
	if got := len(keys2); got != 1 {
		t.Fatalf("expected to have 1 keys even if we save the same original url but saved %d short urls", got)
	} // the key is the same, so only the first one matters.

	if keys[0] != keys2[0] {
		t.Fatalf("expected keys to be equal if the original url is the same, but got %s = %s ", keys[0], keys2[0])
	}

	// clear db
	e.POST("/clear_cache").Expect().Status(httptest.StatusOK)
	if got := db.Len(); got != 0 {
		t.Fatalf("expected database to have 0 registered objects after /clear_cache but has %d", got)
	}

	// give it some time to release the db connection
	db.Close()
	time.Sleep(1 * time.Second)
	// close the file
	if err := f.Close(); err != nil {
		t.Fatalf("unable to close the file: %s", f.Name())
	}

	// and remove the file
	if err := os.Remove(f.Name()); err != nil {
		t.Fatalf("unable to remove the file from %s", f.Name())
	}

	time.Sleep(1 * time.Second)

}

File: tutorial/url-shortener/resources/css/style.css

body{
    background-color:silver;
}

File: tutorial/url-shortener/store.go

package main

import (
	"bytes"

	"github.com/coreos/bbolt"
)

// Panic panics, change it if you don't want to panic on critical INITIALIZE-ONLY-ERRORS
var Panic = func(v interface{}) {
	panic(v)
}

// Store is the store interface for urls.
// Note: no Del functionality.
type Store interface {
	Set(key string, value string) error // error if something went wrong
	Get(key string) string              // empty value if not found
	Len() int                           // should return the number of all the records/tables/buckets
	Close()                             // release the store or ignore
}

var (
	tableURLs = []byte("urls")
)

// DB representation of a Store.
// Only one table/bucket which contains the urls, so it's not a fully Database,
// it works only with single bucket because that all we need.
type DB struct {
	db *bolt.DB
}

var _ Store = &DB{}

// openDatabase open a new database connection
// and returns its instance.
func openDatabase(stumb string) *bolt.DB {
	// Open the data(base) file in the current working directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(stumb, 0600, nil)
	if err != nil {
		Panic(err)
	}

	// create the buckets here
	var tables = [...][]byte{
		tableURLs,
	}

	db.Update(func(tx *bolt.Tx) (err error) {
		for _, table := range tables {
			_, err = tx.CreateBucketIfNotExists(table)
			if err != nil {
				Panic(err)
			}
		}

		return
	})

	return db
}

// NewDB returns a new DB instance, its connection is opened.
// DB implements the Store.
func NewDB(stumb string) *DB {
	return &DB{
		db: openDatabase(stumb),
	}
}

// Set sets a shorten url and its key
// Note: Caller is responsible to generate a key.
func (d *DB) Set(key string, value string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(tableURLs)
		// Generate ID for the url
		// Note: we could use that instead of a random string key
		// but we want to simulate a real-world url shortener
		// so we skip that.
		// id, _ := b.NextSequence()
		if err != nil {
			return err
		}

		k := []byte(key)
		valueB := []byte(value)
		c := b.Cursor()

		found := false
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(valueB, v) {
				found = true
				break
			}
		}
		// if value already exists don't re-put it.
		if found {
			return nil
		}

		return b.Put(k, []byte(value))
	})
}

// Clear clears all the database entries for the table urls.
func (d *DB) Clear() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(tableURLs)
	})
}

// Get returns a url by its key.
//
// Returns an empty string if not found.
func (d *DB) Get(key string) (value string) {
	keyB := []byte(key)
	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(keyB, k) {
				value = string(v)
				break
			}
		}

		return nil
	})

	return
}

// GetByValue returns all keys for a specific (original) url value.
func (d *DB) GetByValue(value string) (keys []string) {
	valueB := []byte(value)
	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		// first for the bucket's table "urls"
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Equal(valueB, v) {
				keys = append(keys, string(k))
			}
		}

		return nil
	})

	return
}

// Len returns all the "shorted" urls length
func (d *DB) Len() (num int) {
	d.db.View(func(tx *bolt.Tx) error {

		// Assume bucket exists and has keys
		b := tx.Bucket(tableURLs)
		if b == nil {
			return nil
		}

		b.ForEach(func([]byte, []byte) error {
			num++
			return nil
		})
		return nil
	})
	return
}

// Close shutdowns the data(base) connection.
func (d *DB) Close() {
	if err := d.db.Close(); err != nil {
		Panic(err)
	}
}

File: tutorial/url-shortener/templates/index.html

<html>

<head>
    <meta charset="utf-8">
    <title>Golang URL Shortener</title>
    <link rel="stylesheet" href="/static/css/style.css" />
</head>

<body>
    <h2>Golang URL Shortener</h2>
    <h3>{{ .FORM_RESULT}}</h3>
    <form action="/shorten" method="POST">
        <input type="text" name="url" style="width: 35em;" />
        <input type="submit" value="Shorten!" />
    </form>
    {{ if IsPositive .URL_COUNT }}
        <p>{{ .URL_COUNT }} URLs shortened</p>
    {{ end }}

    <form action="/clear_cache" method="POST">
        <input type="submit" value="Clear DB" />
    </form>
</body>

</html>

Vuejs Todo Mvc

File: tutorial/vuejs-todo-mvc/screen.png
File: tutorial/vuejs-todo-mvc/src/todo/item.go

package todo

type Item struct {
	SessionID string `json:"-"`
	ID        int64  `json:"id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

File: tutorial/vuejs-todo-mvc/src/todo/service.go

package todo

import (
	"sync"
)

type Service interface {
	Get(owner string) []Item
	Save(owner string, newItems []Item) error
}

type MemoryService struct {
	// key = session id, value the list of todo items that this session id has.
	items map[string][]Item
	// protected by locker for concurrent access.
	mu sync.RWMutex
}

func NewMemoryService() *MemoryService {
	return &MemoryService{
		items: make(map[string][]Item, 0),
	}
}

func (s *MemoryService) Get(sessionOwner string) []Item {
	s.mu.RLock()
	items := s.items[sessionOwner]
	s.mu.RUnlock()

	return items
}

func (s *MemoryService) Save(sessionOwner string, newItems []Item) error {
	var prevID int64
	for i := range newItems {
		if newItems[i].ID == 0 {
			newItems[i].ID = prevID
			prevID++
		}
	}

	s.mu.Lock()
	s.items[sessionOwner] = newItems
	s.mu.Unlock()
	return nil
}

File: tutorial/vuejs-todo-mvc/src/web/controllers/todo_controller.go

package controllers

import (
	"github.com/kataras/iris/_examples/tutorial/vuejs-todo-mvc/src/todo"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

// TodoController is our TODO app's web controller.
type TodoController struct {
	Service todo.Service

	Session *sessions.Session
}

// BeforeActivation called once before the server ran, and before
// the routes and dependencies binded.
// You can bind custom things to the controller, add new methods, add middleware,
// add dependencies to the struct or the method(s) and more.
func (c *TodoController) BeforeActivation(b mvc.BeforeActivation) {
	// this could be binded to a controller's function input argument
	// if any, or struct field if any:
	b.Dependencies().Add(func(ctx iris.Context) (items []todo.Item) {
		ctx.ReadJSON(&items)
		return
	})
}

// Get handles the GET: /todos route.
func (c *TodoController) Get() []todo.Item {
	return c.Service.Get(c.Session.ID())
}

// PostItemResponse the response data that will be returned as json
// after a post save action of all todo items.
type PostItemResponse struct {
	Success bool `json:"success"`
}

var emptyResponse = PostItemResponse{Success: false}

// Post handles the POST: /todos route.
func (c *TodoController) Post(newItems []todo.Item) PostItemResponse {
	if err := c.Service.Save(c.Session.ID(), newItems); err != nil {
		return emptyResponse
	}

	return PostItemResponse{Success: true}
}

func (c *TodoController) GetSync(conn websocket.Connection) {
	// join to the session in order to send "saved"
	// events only to a single user, that means
	// that if user has opened more than one browser window/tab
	// of the same session then the changes will be reflected to one another.
	conn.Join(c.Session.ID())
	conn.On("save", func() { // "save" event from client.
		conn.To(c.Session.ID()).Emit("saved", nil) // fire a "saved" event to the rest of the clients w.
	})

	conn.Wait()
}

File: tutorial/vuejs-todo-mvc/src/web/main.go

package main

import (
	"github.com/kataras/iris/_examples/tutorial/vuejs-todo-mvc/src/todo"
	"github.com/kataras/iris/_examples/tutorial/vuejs-todo-mvc/src/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"

	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	// serve our app in public, public folder
	// contains the client-side vue.js application,
	// no need for any server-side template here,
	// actually if you're going to just use vue without any
	// back-end services, you can just stop afer this line and start the server.
	app.StaticWeb("/", "./public")

	// configure the http sessions.
	sess := sessions.New(sessions.Config{
		Cookie: "iris_session",
	})

	// configure the websocket server.
	ws := websocket.New(websocket.Config{})

	// create a sub router and register the client-side library for the iris websockets,
	// you could skip it but iris websockets supports socket.io-like API.
	todosRouter := app.Party("/todos")
	// http://localhost:8080/todos/iris-ws.js
	// serve the javascript client library to communicate with
	// the iris high level websocket event system.
	todosRouter.Any("/iris-ws.js", websocket.ClientHandler())

	// create our mvc application targeted to /todos relative sub path.
	todosApp := mvc.New(todosRouter)

	// any dependencies bindings here...
	todosApp.Register(
		todo.NewMemoryService(),
		sess.Start,
		ws.Upgrade,
	)

	// controllers registration here...
	todosApp.Handle(new(controllers.TodoController))

	// start the web server at http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutVersionChecker)
}

File: tutorial/vuejs-todo-mvc/src/web/public/css/index

index.css is not here to reduce the disk space for the examples.
https://unpkg.com/todomvc-app-css@2.0.4/index.css is used instead.

File: tutorial/vuejs-todo-mvc/src/web/public/index.html

<!doctype html>
<html data-framework="vue">

<head>
  <meta charset="utf-8">
  <title>Iris + Vue.js ? TodoMVC</title>
  <link rel="stylesheet" href="https://unpkg.com/todomvc-app-css@2.0.4/index.css">
  <!-- this needs to be loaded before guide's inline scripts -->
  <script src="https://vuejs.org/js/vue.js"></script>
  <!-- $http -->
  <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
  <!-- -->
  <script src="https://unpkg.com/director@1.2.8/build/director.js"></script>
  <!-- websocket sync between multiple tabs -->
  <script src="/todos/iris-ws.js"></script>
  <!-- -->
  <style>
    [v-cloak] {
      display: none;
    }
  </style>
</head>

<body>
  <section class="todoapp">
    <header class="header">
      <h1>todos</h1>
      <input class="new-todo" autofocus autocomplete="off" placeholder="What needs to be done?" v-model="newTodo" @keyup.enter="addTodo">
    </header>
    <section class="main" v-show="todos.length" v-cloak>
      <input class="toggle-all" type="checkbox" v-model="allDone">
      <ul class="todo-list">
        <li v-for="todo in filteredTodos" class="todo" :key="todo.id" :class="{ completed: todo.completed, editing: todo == editedTodo }">
          <div class="view">
             <!-- v-model="todo.completed" -->
            <input class="toggle" type="checkbox" @click="completeTodo(todo)">
            <label @dblclick="editTodo(todo)">{{ todo.title }}</label>
            <button class="destroy" @click="removeTodo(todo)"></button>
          </div>
          <input class="edit" type="text" v-model="todo.title" v-todo-focus="todo == editedTodo" @blur="doneEdit(todo)" @keyup.enter="doneEdit(todo)"
            @keyup.esc="cancelEdit(todo)">
        </li>
      </ul>
    </section>
    <footer class="footer" v-show="todos.length" v-cloak>
      <span class="todo-count">
        <strong>{{ remaining }}</strong> {{ remaining | pluralize }} left
      </span>
      <ul class="filters">
        <li>
          <a href="#/all" :class="{ selected: visibility == 'all' }">All</a>
        </li>
        <li>
          <a href="#/active" :class="{ selected: visibility == 'active' }">Active</a>
        </li>
        <li>
          <a href="#/completed" :class="{ selected: visibility == 'completed' }">Completed</a>
        </li>
      </ul>
      <button class="clear-completed" @click="removeCompleted" v-show="todos.length > remaining">
        Clear completed
      </button>
    </footer>
  </section>
  <footer class="info">
    <p>Double-click to edit a todo</p>
  </footer>

  <script src="/js/app.js"></script>
</body>

</html>

File: tutorial/vuejs-todo-mvc/src/web/public/js/app.js

// Full spec-compliant TodoMVC with Iris
// and hash-based routing in ~200 effective lines of JavaScript.

var socket = new Ws("ws://localhost:8080/todos/sync");

socket.On("saved", function () {
  // console.log("receive: on saved");
  fetchTodos(function (items) {
    app.todos = items
  });
});


function fetchTodos(onComplete) {
  axios.get("/todos").then(response => {
    if (response.data === null) {
      return;
    }

    onComplete(response.data);
  });
}

var todoStorage = {
  fetch: function () {
    var todos = [];
    fetchTodos(function (items) {
      for (var i = 0; i < items.length; i++) {
        todos.push(items[i]);
      }
    });
    return todos;
  },
  save: function (todos) {
    axios.post("/todos", JSON.stringify(todos)).then(response => {
      if (!response.data.success) {
        window.alert("saving had a failure");
        return;
      }
      // console.log("send: save");
      socket.Emit("save")
    });
  }
}

// visibility filters
var filters = {
  all: function (todos) {
    return todos
  },
  active: function (todos) {
    return todos.filter(function (todo) {
      return !todo.completed
    })
  },
  completed: function (todos) {
    return todos.filter(function (todo) {
      return todo.completed
    })
  }
}

// app Vue instance
var app = new Vue({
  // app initial state
  data: {
    todos: todoStorage.fetch(),
    newTodo: '',
    editedTodo: null,
    visibility: 'all'
  },

  // we will not use the "watch" as it works with the fields like "hasChanges"
  // and callbacks to make it true but let's keep things very simple as it's just a small getting started.
  // // watch todos change for persistence
  // watch: {
  //   todos: {
  //     handler: function (todos) {
  //       if (app.hasChanges) {
  //         todoStorage.save(todos);
  //         app.hasChanges = false;
  //       }

  //     },
  //     deep: true
  //   }
  // },

  // computed properties
  // http://vuejs.org/guide/computed.html
  computed: {
    filteredTodos: function () {
      return filters[this.visibility](this.todos)
    },
    remaining: function () {
      return filters.active(this.todos).length
    },
    allDone: {
      get: function () {
        return this.remaining === 0
      },
      set: function (value) {
        this.todos.forEach(function (todo) {
          todo.completed = value
        })
        this.notifyChange();
      }
    }
  },

  filters: {
    pluralize: function (n) {
      return n === 1 ? 'item' : 'items'
    }
  },

  // methods that implement data logic.
  // note there's no DOM manipulation here at all.
  methods: {
    notifyChange: function () {
      todoStorage.save(this.todos)
    },
    addTodo: function () {
      var value = this.newTodo && this.newTodo.trim()
      if (!value) {
        return
      }
      this.todos.push({
        id: this.todos.length + 1, // just for the client-side.
        title: value,
        completed: false
      })
      this.newTodo = ''
      this.notifyChange();
    },

    completeTodo: function (todo) {
      if (todo.completed) {
        todo.completed = false;
      } else {
        todo.completed = true;
      }
      this.notifyChange();
    },
    removeTodo: function (todo) {
      this.todos.splice(this.todos.indexOf(todo), 1)
      this.notifyChange();
    },

    editTodo: function (todo) {
      this.beforeEditCache = todo.title
      this.editedTodo = todo
    },

    doneEdit: function (todo) {
      if (!this.editedTodo) {
        return
      }
      this.editedTodo = null
      todo.title = todo.title.trim();
      if (!todo.title) {
        this.removeTodo(todo);
      }
      this.notifyChange();
    },

    cancelEdit: function (todo) {
      this.editedTodo = null
      todo.title = this.beforeEditCache
    },

    removeCompleted: function () {
      this.todos = filters.active(this.todos);
      this.notifyChange();
    }
  },

  // a custom directive to wait for the DOM to be updated
  // before focusing on the input field.
  // http://vuejs.org/guide/custom-directive.html
  directives: {
    'todo-focus': function (el, binding) {
      if (binding.value) {
        el.focus()
      }
    }
  }
})

// handle routing
function onHashChange() {
  var visibility = window.location.hash.replace(/#\/?/, '')
  if (filters[visibility]) {
    app.visibility = visibility
  } else {
    window.location.hash = ''
    app.visibility = 'all'
  }
}

window.addEventListener('hashchange', onHashChange)
onHashChange()

// mount
app.$mount('.todoapp')

File: tutorial/vuejs-todo-mvc/src/web/public/js/lib/vue

vue.js is not here to reduce the disk space for the examples.
Instead https://vuejs.org/js/vue.js is used instead.

View

Context View Data

File: view/context-view-data/main.go

package main

import (
	"time"

	"github.com/kataras/iris"
)

const (
	DefaultTitle  = "My Awesome Site"
	DefaultLayout = "layouts/layout.html"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout

	// set the view engine target to ./templates folder
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Use(func(ctx iris.Context) {
		// set the title, current time and a layout in order to be used if and when the next handler(s) calls the .Render function
		ctx.ViewData("Title", DefaultTitle)
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.ViewLayout(DefaultLayout)

		ctx.Next()
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("BodyMessage", "a sample text here... setted by the route handler")
		if err := ctx.View("index.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	})

	app.Get("/about", func(ctx iris.Context) {
		ctx.ViewData("Title", "My About Page")
		ctx.ViewData("BodyMessage", "about text here... setted by the route handler")

		// same file, just to keep things simple.
		if err := ctx.View("index.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	})

	// http://localhost:8080
	// http://localhost:8080/about
	app.Run(iris.Addr(":8080"))
}

// Notes: ViewData("", myCustomStruct{}) will set this myCustomStruct value as a root binding data,
// so any View("other", "otherValue") will probably fail.
// To clear the binding data: ctx.Set(ctx.Application().ConfigurationReadOnly().GetViewDataContextKey(), nil)

File: view/context-view-data/templates/index.html

<h1>
	Title: {{.Title}}
</h1>
<h3>{{.BodyMessage}} </h3>

<hr/>

Current time: {{.CurrentTime}}

File: view/context-view-data/templates/layouts/layout.html

<html>
<head>
<title>My WebsiteLayout</title>

</head>
<body>
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

Embedding Templates Into App
File: view/embedding-templates-into-app/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Layout("layouts/layout.html")
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	// $ go get -u github.com/shuLhan/go-bindata/...
	// $ go-bindata ./templates/...
	// $ go build
	// $ ./embedding-templates-into-app
	// html files are not used, you can delete the folder and run the example.
	tmpl.Binary(Asset, AssetNames) // <-- IMPORTANT

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// remove the layout for a specific route
	app.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := app.Party("/my").Layout("layouts/mylayout.html")
	{ // both of these will use the layouts/mylayout.html as their layout.
		my.Get("/", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
		my.Get("/other", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
	}

	// http://localhost:8080
	// http://localhost:8080/nolayout
	// http://localhost:8080/my
	// http://localhost:8080/my/other
	app.Run(iris.Addr(":8080"))
}

// Note for new Gophers:
// `go build` is used instead of `go run main.go` as the example comments says
// otherwise you will get compile errors, this is a Go thing;
// because you have multiple files in the `package main`.

File: view/embedding-templates-into-app/templates/layouts/layout.html

<html>
<head>
<title>Layout</title>

</head>
<body>
	<h1>This is the global layout</h1>
	<br />
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

File: view/embedding-templates-into-app/templates/layouts/mylayout.html

<html>
<head>
<title>my Layout</title>

</head>
<body>
	<h1>This is the layout for the /my/ and /my/other routes only</h1>
	<br />
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

File: view/embedding-templates-into-app/templates/page1.html

<div style="background-color: black; color: blue">

	<h1>Page 1 {{ greet "iris developer"}}</h1>

	{{ render "partials/page1_partial1.html"}}

</div>

File: view/embedding-templates-into-app/templates/partials/page1_partial1.html

<div style="background-color: white; color: red">
	<h1>Page 1's Partial 1</h1>
</div>

Overview

File: view/overview/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	// - standard html  | iris.HTML(...)
	// - django         | iris.Django(...)
	// - pug(jade)      | iris.Pug(...)
	// - handlebars     | iris.Handlebars(...)
	// - amber          | iris.Amber(...)
	// with default template funcs:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	app.RegisterView(iris.HTML("./templates", ".html"))
	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("Name", "iris") // the .Name inside the ./templates/hi.html
		ctx.Gzip(true)               // enable gzip for big files
		ctx.View("hi.html")          // render the template with the file name relative to the './templates'

	})

	// http://localhost:8080/
	app.Run(iris.Addr(":8080"))
}

/*
Note:

In case you're wondering, the code behind the view engines derives from the "github.com/kataras/iris/view" package,
access to the engines' variables can be granded by "github.com/kataras/iris" package too.

    iris.HTML(...) is a shortcut of view.HTML(...)
    iris.Django(...)     >> >>      view.Django(...)
    iris.Pug(...)        >> >>      view.Pug(...)
    iris.Handlebars(...) >> >>      view.Handlebars(...)
    iris.Amber(...)      >> >>      view.Amber(...)
*/

File: view/overview/templates/hi.html

<html>

<head>
	<title>Hi iris</title>
</head>

<body>
	<h1>Hi {{.Name}} </h1>
</body>

</html>

Template Html 0

File: view/template_html_0/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New() // defaults to these

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
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
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

/*
Note:

In case you're wondering, the code behind the view engines derives from the "github.com/kataras/iris/view" package,
access to the engines' variables can be granded by "github.com/kataras/iris" package too.

    iris.HTML(...) is a shortcut of view.HTML(...)
    iris.Django(...)     >> >>      view.Django(...)
    iris.Pug(...)        >> >>      view.Pug(...)
    iris.Handlebars(...) >> >>      view.Handlebars(...)
    iris.Amber(...)      >> >>      view.Amber(...)
*/

File: view/template_html_0/templates/hi.html

<html>
<head>
<title>{{.Title}}</title>
</head>
<body>
	<h1>Hi {{.Name}} </h1>
</body>
</html>

Template Html 1

File: view/template_html_1/main.go

package main

import (
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))
	// TIP: append .Reload(true) to reload the templates on each request.

	app.Get("/", func(ctx iris.Context) {
		ctx.Gzip(true)
		ctx.ViewData("", mypage{"My Page title", "Hello world!"})
		ctx.View("mypage.html")
		// Note that: you can pass "layout" : "otherLayout.html" to bypass the config's Layout property
		// or view.NoLayout to disable layout on this render action.
		// third is an optional parameter
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

File: view/template_html_1/templates/layout.html

<html>
<head>
<title>My Layout</title>

</head>
<body>
	<h1>[layout] Body content is below...</h1>
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

File: view/template_html_1/templates/mypage.html

<h1>
	Title: {{.Title}}
</h1>
<h3>Message: {{.Message}} </h3>

Template Html 2

File: view/template_html_2/main.go

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Layout("layouts/layout.html")
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// remove the layout for a specific route
	app.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := app.Party("/my").Layout("layouts/mylayout.html")
	{ // both of these will use the layouts/mylayout.html as their layout.
		my.Get("/", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
		my.Get("/other", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
	}

	// http://localhost:8080
	// http://localhost:8080/nolayout
	// http://localhost:8080/my
	// http://localhost:8080/my/other
	app.Run(iris.Addr(":8080"))
}

File: view/template_html_2/templates/layouts/layout.html

<html>
<head>
<title>Layout</title>

</head>
<body>
	<h1>This is the global layout</h1>
	<br />
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

File: view/template_html_2/templates/layouts/mylayout.html

<html>
<head>
<title>my Layout</title>

</head>
<body>
	<h1>This is the layout for the /my/ and /my/other routes only</h1>
	<br />
	<!-- Render the current template here -->
	{{ yield }}
</body>
</html>

File: view/template_html_2/templates/page1.html

<div style="background-color: black; color: blue">

	<h1>Page 1 {{ greet "iris developer"}}</h1>

	{{ render "partials/page1_partial1.html"}}

</div>

File: view/template_html_2/templates/partials/page1_partial1.html

<div style="background-color: white; color: red">
	<h1>Page 1's Partial 1</h1>
</div>

Template Html 3

File: view/template_html_3/main.go

// Package main an example on how to naming your routes & use the custom 'url path' HTML Template Engine, same for other template engines.
package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	mypathRoute := app.Get("/mypath", writePathHandler)
	mypathRoute.Name = "my-page1"

	mypath2Route := app.Get("/mypath2/{paramfirst}/{paramsecond}", writePathHandler)
	mypath2Route.Name = "my-page2"

	mypath3Route := app.Get("/mypath3/{paramfirst}/statichere/{paramsecond}", writePathHandler)
	mypath3Route.Name = "my-page3"

	mypath4Route := app.Get("/mypath4/{paramfirst}/statichere/{paramsecond}/{otherparam}/{something:path}", writePathHandler)
	// same as: app.Get("/mypath4/:paramfirst/statichere/:paramsecond/:otherparam/*something", writePathHandler)
	mypath4Route.Name = "my-page4"

	// same with Handle/Func
	mypath5Route := app.Handle("GET", "/mypath5/{paramfirst}/statichere/{paramsecond}/{otherparam}/anything/{something:path}", writePathHandler)
	mypath5Route.Name = "my-page5"

	mypath6Route := app.Get("/mypath6/{paramfirst}/{paramsecond}/statichere/{paramThirdAfterStatic}", writePathHandler)
	mypath6Route.Name = "my-page6"

	app.Get("/", func(ctx iris.Context) {
		// for /mypath6...
		paramsAsArray := []string{"theParam1", "theParam2", "paramThirdAfterStatic"}
		ctx.ViewData("ParamsAsArray", paramsAsArray)
		if err := ctx.View("page.html"); err != nil {
			panic(err)
		}
	})

	app.Get("/redirect/{namedRoute}", func(ctx iris.Context) {
		routeName := ctx.Params().Get("namedRoute")
		r := app.GetRoute(routeName)
		if r == nil {
			ctx.StatusCode(404)
			ctx.Writef("Route with name %s not found", routeName)
			return
		}

		println("The path of " + routeName + "is: " + r.Path)
		// if routeName == "my-page1"
		// prints: The path of of my-page1 is: /mypath
		// if it's a path which takes named parameters
		// then use "r.ResolvePath(paramValuesHere)"
		ctx.Redirect(r.Path)
		// http://localhost:8080/redirect/my-page1 will redirect to -> http://localhost:8080/mypath
	})

	// http://localhost:8080
	// http://localhost:8080/redirect/my-page1
	app.Run(iris.Addr(":8080"))

}

func writePathHandler(ctx iris.Context) {
	ctx.Writef("Hello from %s.", ctx.Path())
}

File: view/template_html_3/templates/page.html

<a href="{{urlpath "my-page1"}}">/mypath</a>
<br />
<br />

<a href="{{urlpath "my-page2" "theParam1" "theParam2"}}">/mypath2/{paramfirst}/{paramsecond}</a>
<br />
<br />

<a href="{{urlpath "my-page3" "theParam1" "theParam2AfterStatic"}}">/mypath3/{paramfirst}/statichere/{paramsecond}</a>
<br />
<br />

<a href="{{urlpath "my-page4" "theParam1" "theparam2AfterStatic"  "otherParam"  "matchAnything"}}">
  /mypath4/{paramfirst}/statichere/{paramsecond}/{otherparam}/{something:path}</a>
<br />
<br />

<a href="{{urlpath "my-page5" "theParam1" "theParam2Afterstatichere" "otherParam"  "matchAnythingAfterStatic"}}">
  /mypath5/{paramfirst}/statichere/{paramsecond}/{otherparam}/anything/{anything:path}</a>
<br />
<br />

<a href={{urlpath "my-page6" .ParamsAsArray }}>
  /mypath6/{paramfirst}/{paramsecond}/statichere/{paramThirdAfterStatic}
</a>

Template Html 4

File: view/template_html_4/hosts

# Copyright (c) 1993-2009 Microsoft Corp.

#

# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.

#

# This file contains the mappings of IP addresses to host names. Each

# entry should be kept on an individual line. The IP address should

# be placed in the first column followed by the corresponding host name.

# The IP address and the host name should be separated by at least one

# space.

#

# Additionally, comments (such as these) may be inserted on individual

# lines or following the machine name denoted by a '#' symbol.

#

# For example:

#

#      102.54.94.97     rhino.acme.com          # source server

#       38.25.63.10     x.acme.com              # x client host



# localhost name resolution is handled within DNS itself.

127.0.0.1       localhost

::1             localhost

#-iris-For development machine, you have to configure your dns also for online, search google how to do it if you don't know



127.0.0.1		username1.127.0.0.1

127.0.0.1		username2.127.0.0.1

127.0.0.1		username3.127.0.0.1

127.0.0.1		username4.127.0.0.1

127.0.0.1		username5.127.0.0.1

# note that you can always use custom subdomains

#-END iris-



# Windows: Drive:/Windows/system32/drivers/etc/hosts, on Linux: /etc/hosts

File: view/template_html_4/main.go

// Package main an example on how to naming your routes & use the custom 'url' HTML Template Engine, same for other template engines.
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

const (
	host = "127.0.0.1:8080"
)

func main() {
	app := iris.New()

	// create a custom path reverser, iris let you define your own host and scheme
	// which is useful when you have nginx or caddy in front of iris.
	rv := router.NewRoutePathReverser(app, router.WithHost(host), router.WithScheme("http"))
	// locate and define our templates as usual.
	templates := iris.HTML("./templates", ".html")
	// add a custom func of "url" and pass the rv.URL as its template function body,
	// so {{url "routename" "paramsOrSubdomainAsFirstArgument"}} will work inside our templates.
	templates.AddFunc("url", rv.URL)

	app.RegisterView(templates)

	// wildcard subdomain, will catch username1.... username2.... username3... username4.... username5...
	// that our below links are providing via page.html's first argument which is the subdomain.

	subdomain := app.Party("*.")

	mypathRoute := subdomain.Get("/mypath", emptyHandler)
	mypathRoute.Name = "my-page1"

	mypath2Route := subdomain.Get("/mypath2/{paramfirst}/{paramsecond}", emptyHandler)
	mypath2Route.Name = "my-page2"

	mypath3Route := subdomain.Get("/mypath3/{paramfirst}/statichere/{paramsecond}", emptyHandler)
	mypath3Route.Name = "my-page3"

	mypath4Route := subdomain.Get("/mypath4/{paramfirst}/statichere/{paramsecond}/{otherparam}/{something:path}", emptyHandler)
	mypath4Route.Name = "my-page4"

	mypath5Route := subdomain.Handle("GET", "/mypath5/{paramfirst}/statichere/{paramsecond}/{otherparam}/anything/{something:path}", emptyHandler)
	mypath5Route.Name = "my-page5"

	mypath6Route := subdomain.Get("/mypath6/{paramfirst}/{paramsecond}/staticParam/{paramThirdAfterStatic}", emptyHandler)
	mypath6Route.Name = "my-page6"

	app.Get("/", func(ctx iris.Context) {
		// for username5./mypath6...
		paramsAsArray := []string{"username5", "theParam1", "theParam2", "paramThirdAfterStatic"}
		ctx.ViewData("ParamsAsArray", paramsAsArray)
		if err := ctx.View("page.html"); err != nil {
			panic(err)
		}
	})

	// simple path so you can test it without host mapping and subdomains,
	// at view it make uses of {{urlpath ...}}
	// in order to showcase you that you can use it
	// if you don't want the entire scheme and host to be part of the url.
	app.Get("/mypath7/{paramfirst}/{paramsecond}/static/{paramthird}", emptyHandler).Name = "my-page7"

	// http://127.0.0.1:8080
	app.Run(iris.Addr(host))
}

func emptyHandler(ctx iris.Context) {
	ctx.Writef("Hello from subdomain: %s , you're in path:  %s", ctx.Subdomain(), ctx.Path())
}

// Note:
// If you got an empty string on {{ url }} or {{ urlpath }} it means that
// args length are not aligned with the route's parameters length
// or the route didn't found by the passed name.

File: view/template_html_4/templates/page.html

<!-- the only difference between normal named routes and dynamic subdomains named routes is that the first argument of  url
is the subdomain part instead of named parameter-->

<a href="{{url "my-page1" "username1"}}">username1.127.0.0.1:8080/mypath</a>
<br />
<br />

<a href="{{url  "my-page2" "username2" "theParam1" "theParam2"}}">
    username2.127.0.0.1:8080/mypath2/{paramfirst}/{paramsecond}
</a>
<br />
<br />

<a href="{{url "my-page3" "username3" "theParam1" "theParam2AfterStatic"}}">
    username3.127.0.0.1:8080/mypath3/{paramfirst}/statichere/{paramsecond}
</a>
<br />
<br />

<a href="{{url "my-page4" "username4" "theParam1" "theparam2AfterStatic" "otherParam" "matchAnything"}}">
    username4.127.0.0.1:8080/mypath4/{paramfirst}/statichere/{paramsecond}/{otherParam}/{something:path}
</a>
<br />
<br />

<a href="{{url "my-page5" "username5" "theParam1" "theparam2AfterStatic" "otherParam" "matchAnything"}}">
    username5.127.0.0.1:8080/mypath5/{paramfirst}/statichere/{paramsecond}/{otherparam}/anything/{something:path}
</a>
<br/>
<br/>

<a href="{{url "my-page6" .ParamsAsArray }}">
    username5.127.0.0.1:8080/mypath6/{paramfirst}/{paramsecond}/staticParam/{paramThirdAfterStatic}
</a>
<br/>
<br/>

<a href="{{urlpath "my-page7" "theParam1" "theParam2" "theParam3" }}">
    mypath7/{paramfirst}/{paramsecond}/static/{paramthird}
</a>
<br/>
<br/>

Template Pug 0

File: view/template_pug_0/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Reload(true)                             // reload templates on each request (development mode)
	tmpl.AddFunc("greet", func(s string) string { // add your template func here.
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/", index)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.ViewData("pageTitle", "My Index Page")
	ctx.ViewData("youAreUsingJade", true)
	// Q: why need extension .pug?
	// A: Because you can register more than one view engine per Iris application.
	ctx.View("index.pug")

}

File: view/template_pug_0/templates/index.pug

mixin withGo
  | Generating Go html/template output.

doctype html
html(lang="en")
  head
    title= .pageTitle
    script(type='text/javascript').
      if (foo) {
         bar(1 + 5)
      }
  body
    h1 Jade - template engine
    #container.col
      if .youAreUsingJade
        p {{ greet "iris user" }} <!-- execute template funcs -->
        p You are amazing!
      else
        p Get on it!
      p.
        Jade is #[a(terse)] and simple
        templating language with a
        #[strong focus] on performance
        and powerful features.
      + withGo

Template Pug 1

File: view/template_pug_1/main.go

// Package main shows an example of pug actions based on https://github.com/Joker/jade/tree/master/example/actions
package main

import "github.com/kataras/iris"

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

func main() {
	app := iris.New()

	tmpl := iris.Pug("./templates", ".pug")
	app.RegisterView(tmpl)

	app.Get("/", index)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	job1 := Job{Employer: "Monash B", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of HE"}

	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	ctx.View("index.pug", person)
}

File: view/template_pug_1/templates/index.pug

doctype html
html(lang="en")
	head
		meta(charset="utf-8")
		title Title
	body
		p ads
		ul
			li The name is {{.Name}}.
			li The age is {{.Age}}.

		each _,_ in .Emails
			div An email is {{.}}

		| {{ with .Jobs }}
			each _,_ in .
				div
				 An employer is {{.Employer}}
				 and the role is {{.Role}}
		| {{ end }}

Template Pug 2

File: view/template_pug_2/main.go

package main

import (
	"html/template"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	tmpl := iris.Pug("./templates", ".pug")
	tmpl.Reload(true)                                            // reload templates on each request (development mode)
	tmpl.AddFunc("bold", func(s string) (template.HTML, error) { // add your template func here.
		return template.HTML("<b>" + s + "</b>"), nil
	})

	app.RegisterView(tmpl)

	app.Get("/", index)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.View("index.pug")
}

File: view/template_pug_2/templates/footer.pug

#footer
  p Copyright (c) foobar

File: view/template_pug_2/templates/header.pug

head
  title My Site
  <!-- script(src='/javascripts/jquery.js')
  script(src='/javascripts/app.js') -->

File: view/template_pug_2/templates/index.pug

doctype html
html
  include templates/header.pug
  body
    h1 My Site
    p {{ bold "Welcome to my super lame site."}}
    include templates/footer.pug

Template Pug 3

File: view/template_pug_3/main.go

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	tmpl := iris.Pug("./templates", ".pug")

	app.RegisterView(tmpl)

	app.Get("/", index)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.View("index.pug")
}

File: view/template_pug_3/templates/index.pug

extends templates/layout.pug

block title
  title Article Title

block content
  h1 My Article

File: view/template_pug_3/templates/layout.pug

doctype html
html
  head
    block title
      title Default title
  body
    block content

Write To

File: view/write-to/main.go

package main

import (
	"os"

	"github.com/kataras/iris"
)

type mailData struct {
	Title    string
	Body     string
	RefTitle string
	RefLink  string
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./views", ".html"))

	// you need to call `app.Build` manually before using the `app.View` func,
	// so templates are built in that state.
	app.Build()

	// Or a string-buffered writer to use its body to send an e-mail
	// for sending e-mails you can use the https://github.com/kataras/go-mailer
	// or any other third-party package you like.
	//
	// The template's parsed result will be written to that writer.
	writer := os.Stdout
	err := app.View(writer, "email/simple.html", "shared/email.html", mailData{
		Title:    "This is my e-mail title",
		Body:     "This is my e-mail body",
		RefTitle: "Iris web framework",
		RefLink:  "https://iris-go.com",
	})

	if err != nil {
		app.Logger().Errorf("error from app.View: %v", err)
	}

	app.Run(iris.Addr(":8080"))
}

File: view/write-to/views/email/simple.html

{{.Body}}

File: view/write-to/views/shared/email.html

<h1>{{.Title}}</h1>
<p class="body">
    {{yield}}
</p>

<a href="{{.RefLink}}" target="_new">{{.RefTitle}}</a>

Websocket

Chat

File: websocket/chat/main.go

package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("websockets.html", false) // second parameter: enable gzip?
	})

	setupWebsocket(app)

	// x2
	// http://localhost:8080
	// http://localhost:8080
	// write something, press submit, see the result.
	app.Run(iris.Addr(":8080"))
}

func setupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/echo", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner with:
		// c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}

File: websocket/chat/websockets.html

<!-- the message's input -->
<input id="input" type="text" />

<!-- when clicked then an iris websocket event will be sent to the server, at this example we registered the 'chat' -->
<button onclick="send()">Send</button>

<!-- the messages will be shown here -->
<pre id="output"></pre>
<!-- import the iris client-side library for browser-->
<script src="/iris-ws.js"></script>

<script>
    var scheme = document.location.protocol == "https:" ? "wss" : "ws";
    var port = document.location.port ? (":" + document.location.port) : "";
    // see app.Get("/echo", ws.Handler()) on main.go
    var wsURL = scheme + "://" + document.location.hostname + port+"/echo";

    var input = document.getElementById("input");
    var output = document.getElementById("output");

    // Ws comes from the auto-served '/iris-ws.js'
    var socket = new Ws(wsURL)
    socket.OnConnect(function () {
        output.innerHTML += "Status: Connected\n";
    });

    socket.OnDisconnect(function () {
        output.innerHTML += "Status: Disconnected\n";
    });

    // read events from the server
    socket.On("chat", function (msg) {
        addMessage(msg);
    });

    function send() {
        addMessage("Me: " + input.value); // write ourselves
        socket.Emit("chat", input.value);// send chat event data to the websocket server
        input.value = ""; // clear the input
    }

    function addMessage(msg) {
        output.innerHTML += msg + "\n";
    }
</script>

Connectionlist

File: websocket/connectionlist/main.go

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html")) // select the html engine to serve templates

	ws := websocket.New(websocket.Config{})

	// register the server on an endpoint.
	// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
	app.Get("/my_endpoint", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})

	app.StaticWeb("/js", "./static/js") // serve our custom javascript code

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("", clientPage{"Client Page", "localhost:8080"})
		ctx.View("client.html")
	})

	Conn := make(map[websocket.Connection]bool)
	var myChatRoom = "room1"
	var mutex = new(sync.Mutex)

	ws.OnConnection(func(c websocket.Connection) {
		c.Join(myChatRoom)
		mutex.Lock()
		Conn[c] = true
		mutex.Unlock()
		c.On("chat", func(message string) {
			if message == "leave" {
				c.Leave(myChatRoom)
				c.To(myChatRoom).Emit("chat", "Client with ID: "+c.ID()+" left from the room and cannot send or receive message to/from this room.")
				c.Emit("chat", "You have left from the room: "+myChatRoom+" you cannot send or receive any messages from others inside that room.")
				return
			}
		})
		c.OnDisconnect(func() {
			mutex.Lock()
			delete(Conn, c)
			mutex.Unlock()
			fmt.Printf("\nConnection with ID: %s has been disconnected!\n", c.ID())
		})
	})

	var delay = 1 * time.Second
	go func() {
		i := 0
		for {
			mutex.Lock()
			broadcast(Conn, fmt.Sprintf("aaaa %d\n", i))
			mutex.Unlock()
			time.Sleep(delay)
			i++
		}
	}()

	go func() {
		i := 0
		for range time.Tick(1 * time.Second) { //another way to get clock signal
			mutex.Lock()
			broadcast(Conn, fmt.Sprintf("aaaa2 %d\n", i))
			mutex.Unlock()
			time.Sleep(delay)
			i++
		}
	}()

	app.Run(iris.Addr(":8080"))
}

func broadcast(Conn map[websocket.Connection]bool, message string) {
	for k := range Conn {
		k.To("room1").Emit("chat", message)
	}
}

File: websocket/connectionlist/static/js/chat.js

var messageTxt;
var messages;

$(function () {

	messageTxt = $("#messageTxt");
	messages = $("#messages");


	w = new Ws("ws://" + HOST + "/my_endpoint");
	w.OnConnect(function () {
		console.log("Websocket connection established");
	});

	w.OnDisconnect(function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	});

	w.On("chat", function (message) {
		appendMessage($("<div>" + message + "</div>"));
	});

	$("#sendBtn").click(function () {
		w.Emit("chat", messageTxt.val().toString());
		messageTxt.val("");
	});

})


function appendMessage(messageDiv) {
    var theDiv = messages[0];
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages);
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}

File: websocket/connectionlist/templates/client.html

<html>

<head>
<title>{{ .Title}}</title>
</head>

<body>
	<div id="messages"
		style="border-width: 1px; border-style: solid; height: 400px; width: 375px;">

	</div>
	<input type="text" id="messageTxt" />
	<button type="button" id="sendBtn">Send</button>
	<script type="text/javascript">
		var HOST = {{.Host}}
	</script>
	<script src="js/vendor/jquery-2.2.3.min.js" type="text/javascript"></script>
	<!-- This is auto-serving by the iris, you don't need to have this file in your disk-->
	<script src="/iris-ws.js" type="text/javascript"></script>
	<!-- -->
	<script src="js/chat.js" type="text/javascript"></script>
</body>

</html>

Custom Go Client

File: websocket/custom-go-client/main.go

package main

// Run first `go run main.go server`
// and `go run main.go client` as many times as you want.
// Originally written by: github.com/antlaw to describe an old issue.
import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"

	xwebsocket "golang.org/x/net/websocket"
)

// WS is the current websocket connection
var WS *xwebsocket.Conn

func main() {
	if len(os.Args) == 2 && strings.ToLower(os.Args[1]) == "server" {
		ServerLoop()
	} else if len(os.Args) == 2 && strings.ToLower(os.Args[1]) == "client" {
		ClientLoop()
	} else {
		fmt.Println("wsserver [server|client]")
	}
}

/////////////////////////////////////////////////////////////////////////
// client side
func sendUntilErr(sendInterval int) {
	i := 1
	for {
		time.Sleep(time.Duration(sendInterval) * time.Second)
		err := SendMessage("2", "all", "objectupdate", "2.UsrSchedule_v1_1")
		if err != nil {
			fmt.Println("failed to send join message", err.Error())
			return
		}
		fmt.Println("objectupdate", i)
		i++
	}
}

func recvUntilErr() {
	var msg = make([]byte, 2048)
	var n int
	var err error
	i := 1
	for {
		if n, err = WS.Read(msg); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%v Received: %s.%v\n", time.Now(), string(msg[:n]), i)
		i++
	}

}

//ConnectWebSocket connect a websocket to host
func ConnectWebSocket() error {
	var origin = "http://localhost/"
	var url = "ws://localhost:8080/socket"
	var err error
	WS, err = xwebsocket.Dial(url, "", origin)
	return err
}

// CloseWebSocket closes the current websocket connection
func CloseWebSocket() error {
	if WS != nil {
		return WS.Close()
	}
	return nil
}

// SendMessage broadcast a message to server
func SendMessage(serverID, to, method, message string) error {
	buffer := []byte(message)
	return SendtBytes(serverID, to, method, buffer)
}

// SendtBytes broadcast a message to server
func SendtBytes(serverID, to, method string, message []byte) error {
	// look https://github.com/kataras/iris/blob/master/websocket/message.go , client.go and client.js
	// to understand the buffer line:
	buffer := []byte(fmt.Sprintf("iris-websocket-message:%v;0;%v;%v;", method, serverID, to))
	buffer = append(buffer, message...)
	_, err := WS.Write(buffer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// ClientLoop connects to websocket server, the keep send and recv dataS
func ClientLoop() {
	for {
		time.Sleep(time.Second)
		err := ConnectWebSocket()
		if err != nil {
			fmt.Println("failed to connect websocket", err.Error())
			continue
		}
		// time.Sleep(time.Second)
		err = SendMessage("2", "all", "join", "dummy2")
		go sendUntilErr(1)
		recvUntilErr()
		err = CloseWebSocket()
		if err != nil {
			fmt.Println("failed to close websocket", err.Error())
		}
	}

}

/////////////////////////////////////////////////////////////////////////
// server side

// OnConnect handles incoming websocket connection
func OnConnect(c websocket.Connection) {
	fmt.Println("socket.OnConnect()")
	c.On("join", func(message string) { OnJoin(message, c) })
	c.On("objectupdate", func(message string) { OnObjectUpdated(message, c) })
	// ok works too c.EmitMessage([]byte("dsadsa"))
	c.OnDisconnect(func() { OnDisconnect(c) })

}

// ServerLoop listen and serve websocket requests
func ServerLoop() {
	app := iris.New()

	ws := websocket.New(websocket.Config{})

	// register the server on an endpoint.
	// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
	app.Get("/socket", ws.Handler())

	ws.OnConnection(OnConnect)
	app.Run(iris.Addr(":8080"))
}

// OnJoin handles Join broadcast group request
func OnJoin(message string, c websocket.Connection) {
	t := time.Now()
	c.Join("server2")
	fmt.Println("OnJoin() time taken:", time.Since(t))
}

// OnObjectUpdated broadcasts to all client an incoming message
func OnObjectUpdated(message string, c websocket.Connection) {
	t := time.Now()
	s := strings.Split(message, ";")
	if len(s) != 3 {
		fmt.Println("OnObjectUpdated() invalid message format:" + message)
		return
	}
	serverID, _, objectID := s[0], s[1], s[2]
	err := c.To("server"+serverID).Emit("objectupdate", objectID)
	if err != nil {
		fmt.Println(err, "failed to broacast object")
		return
	}
	fmt.Println(fmt.Sprintf("OnObjectUpdated() message:%v, time taken: %v", message, time.Since(t)))
}

// OnDisconnect clean up things when a client is disconnected
func OnDisconnect(c websocket.Connection) {
	c.Leave("server2")
	fmt.Println("OnDisconnect(): client disconnected!")

}

Native Messages

File: websocket/native-messages/main.go

package main

import (
	"fmt"

	"github.com/kataras/iris"

	"github.com/kataras/iris/websocket"
)

/* Native messages no need to import the iris-ws.js to the ./templates.client.html
Use of: OnMessage and EmitMessage.


NOTICE: IF YOU HAVE RAN THE PREVIOUS EXAMPLES YOU HAVE TO CLEAR YOUR BROWSER's CACHE
BECAUSE chat.js is different than the CACHED. OTHERWISE YOU WILL GET Ws is undefined from the browser's console, because it will use the cached.
*/

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html")) // select the html engine to serve templates

	ws := websocket.New(websocket.Config{
	// to enable binary messages (useful for protobuf):
	// BinaryMessages: true,
	})

	// register the server on an endpoint.
	// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
	app.Get("/my_endpoint", ws.Handler())

	app.StaticWeb("/js", "./static/js") // serve our custom javascript code

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("", clientPage{"Client Page", "localhost:8080"})
		ctx.View("client.html")
	})

	ws.OnConnection(func(c websocket.Connection) {

		c.OnMessage(func(data []byte) {
			message := string(data)
			c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + message)) // broadcast to all clients except this
			c.EmitMessage([]byte("Me: " + message))                                                    // writes to itself
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})

	})

	app.Run(iris.Addr(":8080"))

}

File: websocket/native-messages/static/js/chat.js

var messageTxt;
var messages;

$(function () {

	messageTxt = $("#messageTxt");
	messages = $("#messages");


	w = new WebSocket("ws://" + HOST + "/my_endpoint");
	w.onopen = function () {
		console.log("Websocket connection enstablished");
	};

	w.onclose = function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	};
	w.onmessage = function(message){
		appendMessage($("<div>" + message.data + "</div>"));
	};


	$("#sendBtn").click(function () {
		w.send(messageTxt.val().toString());
		messageTxt.val("");
	});

})


function appendMessage(messageDiv) {
    var theDiv = messages[0];
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages);
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}

File: websocket/native-messages/templates/client.html

<html>

<head>
<title>{{ .Title}}</title>
</head>

<body>
	<div id="messages"
		style="border-width: 1px; border-style: solid; height: 400px; width: 375px;">

	</div>
	<input type="text" id="messageTxt" />
	<button type="button" id="sendBtn">Send</button>
	<script type="text/javascript">
		var HOST = {{.Host}}
	</script>
	<script src="js/vendor/jquery-2.2.3.min.js" type="text/javascript"></script>
	<script src="js/chat.js" type="text/javascript"></script>
</body>

</html>

Secure

File: websocket/secure/main.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html")) // select the html engine to serve templates

	ws := websocket.New(websocket.Config{})

	// register the server on an endpoint.
	// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
	app.Get("/my_endpoint", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})

	app.StaticWeb("/js", "./static/js")
	app.Get("/", func(ctx iris.Context) {
		// send our custom javascript source file before client really asks for that
		// using the go v1.8's HTTP/2 Push.
		// Note that you have to listen using ListenTLS in order this to work.
		if err := ctx.ResponseWriter().Push("/js/chat.js", nil); err != nil {
			ctx.Application().Logger().Warn(err.Error())
		}
		ctx.ViewData("", clientPage{"Client Page", ctx.Host()})
		ctx.View("client.html")
	})

	var myChatRoom = "room1"

	ws.OnConnection(func(c websocket.Connection) {
		// Context returns the (upgraded) iris.Context of this connection
		// avoid using it, you normally don't need it,
		// websocket has everything you need to authenticate the user BUT if it's necessary
		// then  you use it to receive user information, for example: from headers.

		// ctx := c.Context()

		// join to a room (optional)
		c.Join(myChatRoom)

		c.On("chat", func(message string) {
			if message == "leave" {
				c.Leave(myChatRoom)
				c.To(myChatRoom).Emit("chat", "Client with ID: "+c.ID()+" left from the room and cannot send or receive message to/from this room.")
				c.Emit("chat", "You have left from the room: "+myChatRoom+" you cannot send or receive any messages from others inside that room.")
				return
			}
			// to all except this connection ->
			// c.To(websocket.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message)
			// to all connected clients: c.To(websocket.All)

			// to the client itself ->
			//c.Emit("chat", "Message from myself: "+message)

			//send the message to the whole room,
			//all connections are inside this room will receive this message
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
		})

		// or create a new leave event
		// c.On("leave", func() {
		// 	c.Leave(myChatRoom)
		// })

		c.OnDisconnect(func() {
			fmt.Printf("Connection with ID: %s has been disconnected!\n", c.ID())
		})
	})

	listenTLS(app)

}

// a test listenTLS for our localhost
func listenTLS(app *iris.Application) {

	const (
		testTLSCert = `-----BEGIN CERTIFICATE-----
MIIDBTCCAe2gAwIBAgIJAOYzROngkH6NMA0GCSqGSIb3DQEBBQUAMBkxFzAVBgNV
BAMMDmxvY2FsaG9zdDo4MDgwMB4XDTE3MDIxNzAzNDM1NFoXDTI3MDIxNTAzNDM1
NFowGTEXMBUGA1UEAwwObG9jYWxob3N0OjgwODAwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQCfsiVHO14FpKsi0pvBv68oApQm2MO+dCvq87sDU4E0QJhG
KV1RCUmQVypChEqdLlUQsopcXSyKwbWoyg1/KNHYO3DHMfePb4bC1UD2HENq7Ph2
8QJTEi/CJvUB9hqke/YCoWYdjFiI3h3Hw8q5whGO5XR3R23z69vr5XxoNlcF2R+O
TdkzArd0CWTZS27vbgdnyi9v3Waydh/rl+QRtPUgEoCEqOOkMSMldXO6Z9GlUk9b
FQHwIuEnlSoVFB5ot5cqebEjJnWMLLP83KOCQekJeHZOyjeTe8W0Fy1DGu5fvFNh
xde9e/7XlFE//00vT7nBmJAUV/2CXC8U5lsjLEqdAgMBAAGjUDBOMB0GA1UdDgQW
BBQOfENuLn/t0Z4ZY1+RPWaz7RBH+TAfBgNVHSMEGDAWgBQOfENuLn/t0Z4ZY1+R
PWaz7RBH+TAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQBG7AEEuIq6
rWCE5I2t4IXz0jN7MilqEhUWDbUajl1paYf6Ikx5QhMsFx21p6WEWYIYcnWAKZe2
chAgnnGojuxdx0qjiaH4N4xWGHsWhaesnIF1xJepLlX3kJZQURvRxM4wlljlQPIb
9tqzKP131K1HDqplAtp7nWQ72m3J0ZfzH0mYIUxuaS/uQIVtgKqdilwy/VE5dRZ9
QFIb4G9TnNThXMqgTLjfNr33jVbTuv6fzKHYNbCkP3L10ydEs/ddlREmtsn9nE8Q
XCTIYXzA2kr5kWk7d3LkUiSvu3g2S1Ol1YaIKaOQyRveseCGwR4xohLT+dPUW9dL
3hDVLlwE3mB3
-----END CERTIFICATE-----

`
		testTLSKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAn7IlRzteBaSrItKbwb+vKAKUJtjDvnQr6vO7A1OBNECYRild
UQlJkFcqQoRKnS5VELKKXF0sisG1qMoNfyjR2DtwxzH3j2+GwtVA9hxDauz4dvEC
UxIvwib1AfYapHv2AqFmHYxYiN4dx8PKucIRjuV0d0dt8+vb6+V8aDZXBdkfjk3Z
MwK3dAlk2Utu724HZ8ovb91msnYf65fkEbT1IBKAhKjjpDEjJXVzumfRpVJPWxUB
8CLhJ5UqFRQeaLeXKnmxIyZ1jCyz/NyjgkHpCXh2Tso3k3vFtBctQxruX7xTYcXX
vXv+15RRP/9NL0+5wZiQFFf9glwvFOZbIyxKnQIDAQABAoIBAEzBx4ExW8PCni8i
o5LAm2PTuXniflMwa1uGwsCahmOjGI3AnAWzPRSPkNRf2a0q8+AOsMosTphy+umi
FFKmQBZ6m35i2earaE6FSbABbbYbKGGi/ccH2sSrDOBgdfXRTzF8eiSBrJw8hnvZ
87rNOLtCNnSOdJ7lItODfgRo+fLo4uQenJ8VONYwtwm1ejn8qLXq8O5zF66IYUD6
gAzqOiAWumgZL0tEmndeQ+noe4STpJZlOjiCsA12NiJaKDDeDIn5A/pXce+bYNfJ
k4yoroyq/JXBkhyuZDvX9vYp5AA+Q68h8/KmsKkifUgSGSHun5/80lYyT/f60TLX
PxT9GYECgYEA0s8qck7L29nBBTQ6IPF3GHGmqiRdfH+qhP/Jn4NtoW3XuVe4A15i
REq1L8WAiOUIBnBaD8HzbeioqJJYx1pu7x9h/GCNDhdBfwhTjnBe+JjfLqvJKnc0
HUT5wj4DVqattxKzUW8kTRBSWtVremzeffDo+EL6dnR7Bc02Ibs4WpUCgYEAwe34
Uqhie+/EFr4HjYRUNZSNgYNAJkKHVxk4qGzG5VhvjPafnHUbo+Kk/0QW7eIB+kvR
FDO8oKh9wTBrWZEcLJP4jDIKh4y8hZTo9B8EjxFONXVxZlOSYuGjheL8AiLzE7L9
C1spaKMM/MyxAXDRHpG/NeEgXM7Kn6kUGwJdNekCgYAshLNiEGHcu8+XWcAs1NFh
yB56L9PORuerzpi1pvuv65JzAaNKktQNt/krbXoHbtaTBYb/bOYLf+aeMsmsz9w9
g1MeCQXAxAiA2zFKE1D7Ds2S/ZQt8559z+MusgnicrCcyMY1nFL+M0QxCoD4CaWy
0v1f8EUUXuTcBMo5tV/hQQKBgDoBBW8jsiFDu7DZscSgOde00QZVzZAkAfsJLisi
LfNXGjZdZawUUuoX1iYLpZgNK25D0wtp1hdvjf2Ej/dAMd8bexHjvcaBT7ncqjiq
NmDcWjofIIXspTIyLwjStXGmJnJT7N/CqoYDjtTmHGND7Shpi3mAFn/r0isjFUJm
2J5RAoGALuGXxzmSRWmkIp11F/Qr3PBFWBWkrRWaH2TRLMhrU/wO8kCsSyo4PmAZ
ltOfD7InpDiCu43hcDPQ/29FUbDnmAhvMnmIQuHXGgPF/LhqEhbKPA/o/eZdQVCK
QG+tmveBBIYMed5YbWstZu/95lIHF+u8Hl+Z6xgveozfE5yqiUA=
-----END RSA PRIVATE KEY-----

	`
	)

	// create the key and cert files on the fly, and delete them when this test finished
	certFile, ferr := ioutil.TempFile("", "cert")

	if ferr != nil {
		panic(ferr)
	}

	keyFile, ferr := ioutil.TempFile("", "key")
	if ferr != nil {
		panic(ferr)
	}

	certFile.WriteString(testTLSCert)
	keyFile.WriteString(testTLSKey)

	// https://localhost
	app.Run(iris.TLS("localhost:443", certFile.Name(), keyFile.Name()))

	certFile.Close()
	time.Sleep(50 * time.Millisecond)
	os.Remove(certFile.Name())

	keyFile.Close()
	time.Sleep(50 * time.Millisecond)
	os.Remove(keyFile.Name())
}

File: websocket/secure/static/js/chat.js

var messageTxt;
var messages;

$(function () {

	messageTxt = $("#messageTxt");
	messages = $("#messages");

  /* secure wss because we ListenTLS */
	w = new Ws("wss://" + HOST + "/my_endpoint");
	w.OnConnect(function () {
		console.log("Websocket connection established");
	});

	w.OnDisconnect(function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	});

	w.On("chat", function (message) {
		appendMessage($("<div>" + message + "</div>"));
	});

	$("#sendBtn").click(function () {
		w.Emit("chat", messageTxt.val().toString());
		messageTxt.val("");
	});

})


function appendMessage(messageDiv) {
    var theDiv = messages[0];
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages);
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}

File: websocket/secure/templates/client.html

<html>

<head>
<title>{{ .Title}}</title>
</head>

<body>
	<div id="messages"
		style="border-width: 1px; border-style: solid; height: 400px; width: 375px;">

	</div>
	<input type="text" id="messageTxt" />
	<button type="button" id="sendBtn">Send</button>
	<script type="text/javascript">
		var HOST = {{.Host}}
	</script>
	<script src="/js/vendor/jquery-2.2.3.min.js"></script>
	<!-- This is auto-serving by the iris, you don't need to have this file in your disk-->
	<script src="/iris-ws.js"></script>
	<!-- -->
	<script src="/js/chat.js"></script>
</body>

</html>

Third Party Socketio

File: websocket/third-party-socketio/main.go

package main

import (
	"github.com/kataras/iris"

	"github.com/googollee/go-socket.io"
)

/*
	go get -u github.com/googollee/go-socket.io
*/

func main() {
	app := iris.New()
	server, err := socketio.NewServer(nil)
	if err != nil {
		app.Logger().Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		app.Logger().Infof("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			app.Logger().Infof("emit: %v", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			app.Logger().Infof("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		app.Logger().Errorf("error: %v", err)
	})

	// serve the socket.io endpoint.
	app.Any("/socket.io/{p:path}", iris.FromStd(server))

	// serve the index.html and the javascript libraries at
	// http://localhost:8080
	app.StaticWeb("/", "./public")

	app.Run(iris.Addr("localhost:8080"))
}

File: websocket/third-party-socketio/public/index.html

<!doctype html>
<html>
  <head>
    <title>Socket.IO chat</title>
    <style>
      * { margin: 0; padding: 0; box-sizing: border-box; }
      body { font: 13px Helvetica, Arial; }
      form { background: #000; padding: 3px; position: fixed; bottom: 0; width: 100%; }
      form input { border: 0; padding: 10px; width: 90%; margin-right: .5%; }
      form button { width: 9%; background: rgb(130, 224, 255); border: none; padding: 10px; }
      #messages { list-style-type: none; margin: 0; padding: 0; }
      #messages li { padding: 5px 10px; }
      #messages li:nth-child(odd) { background: #eee; }
    </style>
  </head>
  <body>
    <ul id="messages"></ul>
    <form action="">
      <input id="m" autocomplete="off" /><button>Send</button>
    </form>
    <script src="/socket.io-1.3.7.js"></script>
    <script src="/jquery-1.11.1.js"></script>
    <script>
      var socket = io();
      $('form').submit(function(){
        socket.emit('chat message with ack', $('#m').val(), function(data){
          $('#messages').append($('<li>').text('ACK CALLBACK: ' + data));
        });
        socket.emit('chat message', $('#m').val());
        $('#m').val('');
        return false;
      });
      socket.on('chat message', function(msg){
        $('#messages').append($('<li>').text(msg));
      });
    </script>
  </body>
</html>

