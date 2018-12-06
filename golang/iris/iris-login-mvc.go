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

