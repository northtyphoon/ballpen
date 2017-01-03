package routers

import (
	"github.com/astaxie/beego"
	"github.com/northtyphoon/ballpen/controllers"
)

func init() {
	// Remove default static path and add webpack output path
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/dist", "dist")

	// Add view router
	beego.Router("/", &controllers.MainController{}, "get:Home")

	// Add restful api router
	beego.Router("/api/hash", &controllers.HashController{})
}
