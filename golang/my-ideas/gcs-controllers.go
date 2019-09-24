// main
r.Get("/",               controller.Home.Index)
r.Get("/dashboard",      controller.Dashboard.Index)
r.Get("/project/index",  controller.Project.Index)
r.Get("/project/detail", controller.Project.Detail)
r.Get("/user/login",     controller.User.Login)

//----------------------------------------------------------
// controllers.go
package controller

var Home      = &HomeController{}
var Dashboard = &DashboardController{}
var Project   = &ProjectController{}

/*
r.Get("/",               Home.Index)
r.Get("/dashboard",      Dashboard.Index)
r.Get("/project/index",  Project.Index)
r.Get("/project/detail", Project.Detail)
r.Get("/user/login",     User.Login)
*/

//----------------------------------------------------------
// home.go
package controller

type HomeController struct { }

func (z HomeController) Index() {
}

//----------------------------------------------------------
// dashboard.go
package controller

func (z DashboardController) Index() {
}

//----------------------------------------------------------
// project.go
package controller

func (z ProjectController) Index() {
}

//----------------------------------------------------------
// user.go
package controller

func (z UserController) Login() {
}
