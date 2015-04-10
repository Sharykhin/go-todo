package routers

import (
	"project/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.MainController{},"post:CreateUser")
    beego.Router("/list",&controllers.MainController{},"get:Users")
    
}
