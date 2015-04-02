package controller

import "net/http"
import errorComponent "project/todos/core/components/error"


type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	// 404 error handler
	if req.URL.Path != "/" {
		errorComponent.ErrorHandler(res, req, http.StatusNotFound,"")
		return nil
	}		
	
	flashErrorMessages := ctrl.GetFlashMessages(res,req,"error")
	
	if err := ctrl.Render(res,req, "index", nil, struct {
		TestData string		
		FlashMessage []interface{}
	}{
		TestData: "Test string",		
		FlashMessage: flashErrorMessages,
	}); err != nil {
		return err
	}
	return nil
}
