package routers

import (
	"todo/beego/controllers"
	userController "todo/beego/modules/user/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.MainController{})
    beego.Router("/all/:key",&controllers.CMSController{},"get:AllBlock")
    beego.Router("/user",&userController.UserController{})
}
