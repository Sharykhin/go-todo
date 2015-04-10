package controllers

import (
	"github.com/astaxie/beego"	
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	beego.ViewsPath="modules/user/views"
	this.Data["Website"] = "user module"	
	this.TplNames="user.tpl"
} 
