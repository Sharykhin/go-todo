package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

type Todo struct {
	Title string
}

var todos []Todo

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res, todos)
}

func editHandler(res http.ResponseWriter, req *http.Request) {
	todoTitle := req.URL.Path[len("/edit"):]
	fmt.Println(todoTitle)
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		return
	}
	todo := Todo{Title:todoTitle}
	t.Execute(res, todo)
}

func saveHandler(res http.ResponseWriter, req *http.Request) {

	todoTitle := req.FormValue("todo")
	todos = append(todos, Todo{Title: todoTitle})
	http.Redirect(res, req, "/", http.StatusFound)

}

func main() {

	http.Handle("/vendor/", http.StripPrefix("/vendor/",http.FileServer(http.Dir("public/vendor"))))
	http.HandleFunc("/", defaultHandler)
	//http.HandleFunc("/public/vendor/", func(res http.ResponseWriter, req *http.Request) {
	//	http.ServeFile(res, req, req.URL.Path[1:])
	//})
	
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}
