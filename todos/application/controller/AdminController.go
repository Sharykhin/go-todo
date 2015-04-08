package controller

import (
	"net/http"
	auth "github.com/abbot/go-http-auth"
	"fmt"
) 


type AdminController struct {
	BaseController
}

func (ctrl *AdminController) IndexAction(w http.ResponseWriter, r *auth.AuthenticatedRequest)  {

 	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
	
}


