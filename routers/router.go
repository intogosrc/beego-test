package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/intogosrc/beego-test/controllers"
	"github.com/intogosrc/beego-test/middlewares"
)

func init() {

	apiRoute := &controllers.ApiController{}

	beego.Any("/api/currentUser", apiRoute.CurrentUser)
	beego.Any("/api/users", apiRoute.Users)
	beego.Any("/api/login/account", apiRoute.Login)
	beego.Any("/api/goods/list", apiRoute.GoodsList)
	beego.Any("/api/user/list", apiRoute.UserList)
	beego.Any("/api/user/list2", apiRoute.UserList2)
	beego.Any("/api/user/add", apiRoute.UserAdd)

	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.UserController{})

	// allow CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// middleware
	beego.InsertFilter("/user/*", beego.BeforeRouter, middlewares.Test)

	beego.InsertFilter("/api/*", beego.BeforeRouter, middlewares.ParseJsonBody)
}
