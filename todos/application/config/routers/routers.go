package routers

import (
	controller "project/todos/application/controller"
	adminModule "project/todos/application/modules/admin/controller"
	userModule "project/todos/application/modules/user/controller"
	errorComponent "project/todos/core/components/error"
	"net/http"
	"fmt"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprint(res, err)
			return
		}
	}()

	if err := fn(res, req); err != nil {
		errorComponent.ErrorHandler(res, req, http.StatusInternalServerError,err.Error())
		return	
	}
	
}

func Listen() {

	var indexController controller.IndexController
	var postController controller.PostController
	// one more variant of defining controller
	adminDefaultController := new(adminModule.DefaultController)
	var userDefaultController userModule.DefaultController

	http.Handle("/", appHandler(indexController.IndexAction))
	http.Handle("/posts", appHandler(postController.IndexAction))
	http.Handle("/about", appHandler(postController.AboutAction))
	http.Handle("/admin", appHandler(adminDefaultController.IndexAction))
	http.Handle("/user", appHandler(userDefaultController.IndexAction))
	http.Handle("/user/profile", appHandler(userDefaultController.UserProfileAction))	

}
