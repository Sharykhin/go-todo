package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"log"
)

type Todo struct {
	Title string
}

type Data struct {
	Todos []Todo
	Title string
}

var todos []Todo



func defaultHandler(res http.ResponseWriter, req *http.Request) {
	
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	
	data := Data{Todos:todos}

	t.Execute(res, data)
}

func editHandler(res http.ResponseWriter, req *http.Request) {

	todoTitle := req.URL.Path[len("/edit/"):]	
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := Data{Title: todoTitle}
	t.Execute(res, data)
}

func saveHandler(res http.ResponseWriter, req *http.Request) {

	todoTitle := req.FormValue("todo")
	todos = append(todos, Todo{Title: todoTitle})
	http.Redirect(res, req, "/", http.StatusFound)

}

func main() {

	http.Handle("/vendor/", http.StripPrefix("/vendor/",http.FileServer(http.Dir("public/vendor"))))

	http.HandleFunc("/", defaultHandler)	
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	http.ListenAndServe(":8080", nil)
}
