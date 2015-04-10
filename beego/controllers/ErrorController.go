package controllers

import (
    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error501() {
	beego.Trace("default2")
    c.Data["content"] = "server error"
    c.TplNames = "501.tpl"
}


