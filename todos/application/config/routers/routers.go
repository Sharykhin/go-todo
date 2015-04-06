package routers

import (
	controller "project/todos/application/controller"
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
	
	http.Handle("/", appHandler(indexController.IndexAction))
	http.Handle("/create",appHandler(indexController.CreateTodoAction))	

}
