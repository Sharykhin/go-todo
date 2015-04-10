package routers

import (

	userController "project/beego/modules/user/controllers"
	"project/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.MainController{})
    beego.Router("/all/:key",&controllers.CMSController{},"get:AllBlock")
    beego.Router("/user",&userController.UserController{})
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/create",&controllers.MainController{},"post:CreateUser")
    beego.Router("/list",&controllers.MainController{},"get:Users")
    
}
