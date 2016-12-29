package main

import (
	"github.com/astaxie/beego"
	_ "github.com/northtyphoon/ballpen/routers"
)

func main() {
	// beego.BConfig.Listen.EnableAdmin = true

	beego.Run()
}
