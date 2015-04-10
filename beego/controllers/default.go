package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["xsrfdata"]=template.HTML(c.XsrfFormHtml())
	c.TplNames = "index.tpl"
}

func (c *MainController) Login() {
	c.Data["Website"] = "Login site"
	c.TplNames = "index.tpl"
}

func (this *MainController) Delete() {
	this.Data["Website"] = "Delete method"
	this.TplNames = "index.tpl"
}
