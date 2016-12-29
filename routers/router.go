package routers

import (
	"github.com/astaxie/beego"
	"github.com/northtyphoon/ballpen/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Home")
}
