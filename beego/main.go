package main

import (
	_ "project/beego/routers"
	_ "project/beego/conf/config"
	"project/beego/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

