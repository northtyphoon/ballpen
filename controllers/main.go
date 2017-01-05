package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is the main controller
type MainController struct {
	beego.Controller
}

// Home is the / action
func (c *MainController) Home() {
	c.TplName = "main/home.tpl"
}
