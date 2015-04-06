package controller

import "net/http"
import errorComponent "project/todos/core/components/error"
import Todo "project/todos/application/model/todo"
import sessionComponent "project/todos/core/components/session"


type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	// 404 error handler
	if req.URL.Path != "/" {
		errorComponent.ErrorHandler(res, req, http.StatusNotFound,"")
		return nil
	}		

	todoModel := Todo.New()
	
	flashMessage := ctrl.GetFlashMessages(res,req,"success")

	todos,err := todoModel.FindAll();
	
	if err != nil {
	   return err	
	}
	
	if err := ctrl.Render(res,req, "index", nil, struct {	
		Todos []map[string]interface{}	
		FlashMessage []interface{}
	}{			
		FlashMessage: flashMessage,
		Todos: todos,
	}); err != nil {
		return err
	}
	return nil
}

func (ctrl *IndexController) CreateTodoAction(res http.ResponseWriter, req *http.Request) error {

	if req.Method != "POST" {
		errorComponent.ErrorHandler(res, req, http.StatusMethodNotAllowed,"Method Not Allowed")
		return nil
		
	}

	todoModel := Todo.New()

	title := req.FormValue("title")

	todoModel.SetTitle(title)

	if err := todoModel.Save(); err != nil {
		return err
	}

	session, _ := sessionComponent.Store.Get(req, "session")
	session.AddFlash("Todo has been created","success")
	session.Save(req,res)

	http.Redirect(res,req, "/" , http.StatusFound)
	return nil

}
