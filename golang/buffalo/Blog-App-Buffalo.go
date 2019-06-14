// --------------------actions/app.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/mikaelm1/blog_app/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_blog_app_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}
		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)
		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())
		app.Use(SetCurrentUser)

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", assetsBox)
		// users routes
		auth := app.Group("/users")
		auth.GET("/register", UsersRegisterGet)
		auth.POST("/register", UsersRegisterPost)
		auth.GET("/login", UsersLoginGet)
		auth.POST("/login", UsersLoginPost)
		auth.GET("/logout", UsersLogout)
		postGroup := app.Group("/posts")
		postGroup.GET("/index", PostsIndex)
		postGroup.GET("/create", AdminRequired(PostsCreateGet))
		postGroup.POST("/create", AdminRequired(PostsCreatePost))
		postGroup.GET("/detail/{pid}", PostsDetail)
		postGroup.GET("/edit/{pid}", AdminRequired(PostsEditGet))
		postGroup.POST("/edit/{pid}", AdminRequired(PostsEditPost))
		postGroup.GET("/delete/{pid}", AdminRequired(PostsDelete))
		commentsGroup := app.Group("/comments")
		commentsGroup.Use(LoginRequired)
		commentsGroup.POST("/create/{pid}", CommentsCreatePost)
		commentsGroup.GET("/edit/{cid}", CommentsEditGet)
		commentsGroup.POST("/edit/{cid}", CommentsEditPost)
		commentsGroup.GET("/delete/{cid}", CommentsDelete)
	}

	return app
}

// --------------------actions/render.go
package actions

import (
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public/assets")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"csrf": func() template.HTML {
				return template.HTML("<input name=\"authenticity_token\" value=\"<%= authenticity_token %>\" type=\"hidden\">")
			},
			"markdown2": func(content string) template.HTML {
				output := blackfriday.Run([]byte(content), blackfriday.WithExtensions(blackfriday.CommonExtensions))
				html := bluemonday.UGCPolicy().SanitizeBytes(output)
				return template.HTML(html)
			},
		},
	})
}

// --------------------actions/home.go
package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// actions/users.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/mikaelm1/blog_app/models"
	"github.com/pkg/errors"
)

// UserRegisterGet displays a register form
func UsersRegisterGet(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/register.html"))
}

// UsersRegisterPost adds a User to the DB. This function is mapped to the
// path POST /accounts/register
func UsersRegisterPost(c buffalo.Context) error {
	// Allocate an empty User
	user := &models.User{}
	// Bind user to the html form elements
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := user.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make user available inside the html template
		c.Set("user", user)
		// Make the errors available inside the html template
		c.Set("errors", verrs.Errors)
		// Render again the register.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("users/register.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Account created successfully.")
	// and redirect to the home page
	return c.Redirect(302, "/")
}

// UsersLoginGet displays a login form
func UsersLoginGet(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/login"))
}

// UsersLoginPost logs in a user.
func UsersLoginPost(c buffalo.Context) error {
	user := &models.User{}
	// Bind the user to the html form elements
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)
	err := user.Authorize(tx)
	if err != nil {
		c.Set("user", user)
		verrs := validate.NewErrors()
		verrs.Add("Login", "Invalid email or password.")
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("users/login"))
	}
	c.Session().Set("current_user_id", user.ID)
	c.Flash().Add("success", "Welcome back!")
	return c.Redirect(302, "/")
}

// UsersLogout clears the session and logs out the user.
func UsersLogout(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "Goodbye!")
	return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// AdminRequired requires a user to be logged in and to be an admin before accessing a route.
func AdminRequired(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		user, ok := c.Value("current_user").(*models.User)
		if ok && user.Admin {
			return next(c)
		}
		c.Flash().Add("danger", "You are not authorized to view that page.")
		return c.Redirect(302, "/")
	}
}

func LoginRequired(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		_, ok := c.Value("current_user").(*models.User)
		if ok {
			return next(c)
		}
		c.Flash().Add("danger", "You are not authorized to view that page.")
		return c.Redirect(302, "/")
	}
}

// --------------------actions/posts.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/mikaelm1/blog_app/models"
	"github.com/pkg/errors"
)

// PostsIndex default implementation.
func PostsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	posts := &models.Posts{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all Posts from the DB
	if err := q.All(posts); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("posts", posts)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("posts/index.html"))
}

// PostsCreateGet displays a form to create a post
func PostsCreateGet(c buffalo.Context) error {
	c.Set("post", &models.Post{})
	return c.Render(200, r.HTML("posts/create"))
}

// PostsCreatePost adds a new Post to the db
func PostsCreatePost(c buffalo.Context) error {
	// Allocate an empty Post
	post := &models.Post{}
	user := c.Value("current_user").(*models.User)
	// Bind post to the html form elements
	if err := c.Bind(post); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	post.AuthorID = user.ID
	verrs, err := tx.ValidateAndCreate(post)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("post", post)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("posts/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "New post added successfully.")
	// and redirect to the index page
	return c.Redirect(302, "/")
}

// PostsDetail displays a single post.
func PostsDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	post := &models.Post{}
	if err := tx.Find(post, c.Param("pid")); err != nil {
		return c.Error(404, err)
	}
	author := &models.User{}
	if err := tx.Find(author, post.AuthorID); err != nil {
		return c.Error(404, err)
	}
	c.Set("post", post)
	c.Set("author", author)
	comment := &models.Comment{}
	c.Set("comment", comment)
	comments := models.Comments{}
	if err := tx.BelongsTo(post).All(&comments); err != nil {
		return errors.WithStack(err)
	}
	for i := 0; i < len(comments); i++ {
		u := models.User{}
		if err := tx.Find(&u, comments[i].AuthorID); err != nil {
			return c.Error(404, err)
		}
		comments[i].Author = u
	}
	c.Set("comments", comments)
	return c.Render(200, r.HTML("posts/detail"))
}

// PostsEditGet displays a form to edit the post.
func PostsEditGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	post := &models.Post{}
	if err := tx.Find(post, c.Param("pid")); err != nil {
		return c.Error(404, err)
	}
	c.Set("post", post)
	return c.Render(200, r.HTML("posts/edit.html"))
}

// PostsEditPost updates a post.
func PostsEditPost(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	post := &models.Post{}
	if err := tx.Find(post, c.Param("pid")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(post); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(post)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("post", post)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("posts/edit.html"))
	}
	c.Flash().Add("success", "Post was updated successfully.")
	return c.Redirect(302, "/posts/detail/%s", post.ID)
}

// PostsDelete deletes a post from the db
func PostsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	post := &models.Post{}
	if err := tx.Find(post, c.Param("pid")); err != nil {
		return c.Error(404, err)
	}
	if err := tx.Destroy(post); err != nil {
		return errors.WithStack(err)
	}
	c.Flash().Add("success", "Post was successfully deleted.")
	return c.Redirect(302, "/posts/index")
}

// --------------------actions/comments.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/mikaelm1/blog_app/models"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// CommentsCreatePost adds the comment to the db.
func CommentsCreatePost(c buffalo.Context) error {
	comment := &models.Comment{}
	user := c.Value("current_user").(*models.User)
	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)
	comment.AuthorID = user.ID
	postID, err := uuid.FromString(c.Param("pid"))
	if err != nil {
		return errors.WithStack(err)
	}
	comment.PostID = postID
	verrs, err := tx.ValidateAndCreate(comment)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Flash().Add("danger", "There was an error adding your comment.")
		return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
	}
	c.Flash().Add("success", "Comment added successfully.")
	return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
}

// CommentsEditGet displays a form to edit a comment.
func CommentsEditGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := c.Value("current_user").(*models.User)
	comment := &models.Comment{}
	if err := tx.Find(comment, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	// make sure the comment was made by the logged in user
	if user.ID != comment.AuthorID {
		c.Flash().Add("danger", "You are not authorized to view that page.")
		return c.Redirect(302, "/posts/detail/%s", comment.PostID)
	}
	c.Set("comment", comment)
	return c.Render(200, r.HTML("comments/edit.html"))
}

// CommentsEditPost process the edit form of a comment.
func CommentsEditPost(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	comment := &models.Comment{}
	if err := tx.Find(comment, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}
	user := c.Value("current_user").(*models.User)
	// make sure the comment was made by the logged in user
	if user.ID != comment.AuthorID {
		c.Flash().Add("danger", "You are not authorized to view that page.")
		return c.Redirect(302, "/posts/detail/%s", comment.PostID)
	}
	verrs, err := tx.ValidateAndUpdate(comment)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("comment", comment)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("comments/edit.html"))
	}
	c.Flash().Add("success", "Comment updated successfully.")
	return c.Redirect(302, "/posts/detail/%s", comment.PostID)
}

// CommentsDelete deletes the comment from the db.
func CommentsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	comment := &models.Comment{}
	if err := tx.Find(comment, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	user := c.Value("current_user").(*models.User)
	// only admins and comment creators can delete the comment
	if user.ID != comment.AuthorID && user.Admin == false {
		c.Flash().Add("danger", "You are not authorized to view that page.")
		return c.Redirect(302, "/posts/detail/%s", comment.PostID)
	}
	if err := tx.Destroy(comment); err != nil {
		return errors.WithStack(err)
	}
	c.Flash().Add("success", "Comment deleted successfuly.")
	return c.Redirect(302, "/posts/detail/%s", comment.PostID)
}

// --------------------models/models.go
package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

// --------------------models/users.go
package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uuid.UUID `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Username        string    `json:"username" db:"username"`
	Email           string    `json:"email" db:"email"`
	Admin           bool      `json:"admin" db:"admin"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	Password        string    `json:"-" db:"-"`
	PasswordConfirm string    `json:"-" db:"-"`
}

// Create validates and creates a new User.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(u.Email)
	u.Admin = false
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.PasswordHash = string(pwdHash)
	return tx.ValidateAndCreate(u)
}

// Authorize checks user's password for logging in
func (u *User) Authorize(tx *pop.Connection) error {
	err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with that email address
			return errors.New("User not found.")
		}
		return errors.WithStack(err)
	}
	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return errors.New("Invalid password.")
	}
	return nil
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirm, Message: "Passwords do not match."},
		&UsernameNotTaken{Name: "Username", Field: u.Username, tx: tx},
		&EmailNotTaken{Name: "Email", Field: u.Email, tx: tx},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

type UsernameNotTaken struct {
	Name  string
	Field string
	tx    *pop.Connection
}

func (v *UsernameNotTaken) IsValid(errors *validate.Errors) {
	query := v.tx.Where("username = ?", v.Field)
	queryUser := User{}
	err := query.First(&queryUser)
	if err == nil {
		// found a user with same username
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("The username %s is not available.", v.Field))
	}
}

type EmailNotTaken struct {
	Name  string
	Field string
	tx    *pop.Connection
}

// IsValid performs the validation check for unique emails
func (v *EmailNotTaken) IsValid(errors *validate.Errors) {
	query := v.tx.Where("email = ?", v.Field)
	queryUser := User{}
	err := query.First(&queryUser)
	if err == nil {
		// found a user with the same email
		errors.Add(validators.GenerateKey(v.Name), "An account with that email already exists.")
	}
}

// --------------------models/comments.go
package models

import (
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	uuid "github.com/satori/go.uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Content   string    `json:"content" db:"content"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id"`
	PostID    uuid.UUID `json:"post_id" db:"post_id"`
	Author    User      `json:"-" db:"-"`
}

type Comments []Comment

// Validate gets run every time you call a "pop.Validate*" 
// (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (c *Comment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Content, Name: "Content"},
	), nil
}

// --------------------models/posts.go
package models

import (
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id"`
}

type Posts []Post

// Validate gets run every time you call a "pop.Validate*" 
// (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Post) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Title, Name: "Title"},
		&validators.StringIsPresent{Field: p.Content, Name: "Content"},
	), nil
}
