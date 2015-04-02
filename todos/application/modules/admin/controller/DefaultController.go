package admin

import "net/http"
import controller "project/todos/application/controller"


type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) error {


	if err := ctrl.RenderView(res,req, "modules/admin:index", nil, nil); err != nil {
		return err
	}
	return nil
}
