package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/northtyphoon/ballpen/routers"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("console")

	// The original view files are built and outputed to dist folder
	beego.SetViewsPath("dist/views")

	beego.Run()
}
