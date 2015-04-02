package user

import "net/http"
import controller "project/todos/application/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	if err := ctrl.RenderView(res,req, "modules/user:user", nil, nil); err != nil {
		return err
	}
	return nil
}

func (ctrl *DefaultController) UserProfileAction(res http.ResponseWriter, req *http.Request) error {
	if err := ctrl.Render(res,req, "modules/user:profile", nil, struct {
		UserName string
	}{
		UserName: "John",
	}); err != nil {
		return err
	}
	return nil
}
