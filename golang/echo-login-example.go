// main.go
package main

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// 这里的 studyecho ，实际应用，不应该写在代码中，应该写在配置或环境变量中
var cookieStore = sessions.NewCookieStore([]byte("studyecho"))

func init() {
	rand.Seed(time.Now().UnixNano())

	os.Mkdir("log", 0755)
}

func main() {
	// 创建 echo 实例
	e := echo.New()

	// 配置日志
	configLogger(e)

	// 注册静态文件路由
	e.Static("img", "img")
	e.File("/favicon.ico", "img/favicon.ico")

	// 设置中间件
	setMiddleware(e)

	// 注册路由
	RegisterRoutes(e)

	// 启动服务
	e.Logger.Fatal(e.Start(":2020"))
}

func configLogger(e *echo.Echo) {
	// 定义日志级别
	e.Logger.SetLevel(log.INFO)

	// 记录业务日志
	echoLog, err := os.OpenFile("log/echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}

	// 同时输出到文件和终端
	e.Logger.SetOutput(io.MultiWriter(os.Stdout, echoLog))
}

func setMiddleware(e *echo.Echo) {
	// access log 输出到文件中
	accessLog, err := os.OpenFile("log/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	// 同时输出到终端和文件
	middleware.DefaultLoggerConfig.Output = accessLog
	e.Use(middleware.Logger())

	// 自定义 middleware
	e.Use(AutoLogin)

	e.Use(middleware.Recover())
}

// route.go
package main

import (
	"bytes"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		tpl, err := template.ParseFiles("template/login.html")
		if err != nil {
			ctx.Logger().Error("parse file error:", err)
			return err
		}

		ctx.Logger().Info("this is login page...")

		data := map[string]interface{}{
			"msg": ctx.QueryParam("msg"),
		}

		if user, ok := ctx.Get("user").(*User); ok {
			data["username"] = user.Username
			data["had_login"] = true
		} else {
			sess := getCookieSession(ctx)
			if flashes := sess.Flashes("username"); len(flashes) > 0 {
				data["username"] = flashes[0]
			}
			sess.Save(ctx.Request(), ctx.Response())
		}

		var buf bytes.Buffer
		err = tpl.Execute(&buf, data)
		if err != nil {
			return err
		}

		return ctx.HTML(http.StatusOK, buf.String())
	})

	// 登录
	e.POST("/login", func(ctx echo.Context) error {
		username := ctx.FormValue("username")
		passwd := ctx.FormValue("passwd")
		rememberMe := ctx.FormValue("remember_me")

		if username == "polaris" && passwd == "123567" {
			// 用标准库种 cookie
			cookie := &http.Cookie{
				Name:     "username",
				Value:    username,
				HttpOnly: true,
			}
			if rememberMe == "1" {
				cookie.MaxAge = 7*24*3600	// 7 天
			}
			ctx.SetCookie(cookie)

			return ctx.Redirect(http.StatusSeeOther, "/")
		}

		// 用户名或密码不对，用户名回填，通过 github.com/gorilla/sessions 包实现
		sess := getCookieSession(ctx)
		sess.AddFlash(username, "username")
		err := sess.Save(ctx.Request(), ctx.Response())
		if err != nil {
			return ctx.Redirect(http.StatusSeeOther, "/?msg="+err.Error())
		}

		return ctx.Redirect(http.StatusSeeOther, "/?msg=用户名或密码错误")
	})

	// 退出登录
	e.GET("/logout", func(ctx echo.Context) error {
		cookie := &http.Cookie{
			Name:    "username",
			Value:   "",
			Expires: time.Now().Add(-1e9),
			MaxAge:  -1,
		}
		ctx.SetCookie(cookie)

		return ctx.Redirect(http.StatusSeeOther, "/")
	})
}

func getCookieSession(ctx echo.Context) *sessions.Session {
	sess, _ := cookieStore.Get(ctx.Request(), "request-scope")
	return sess
}

// middleware.go
package main

import "github.com/labstack/echo/v4"

// AutoLogin 如果上次记住了，则自动登录
func AutoLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("username")
		if err == nil && cookie.Value != "" {
			// 实际项目这里可以通过 username 读库获取用户信息
			user := &User{Username: cookie.Value}

			// 放入 context 中
			ctx.Set("user", user)
		}

		return next(ctx)
	}
}

// user.go
package main

type User struct {
	UID      int
	Username string
	Passwd   string
}
