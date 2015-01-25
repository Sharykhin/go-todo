package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
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


	t.Execute(res, struct {
		Todos []Todo
		Title string
	}{Todos: todos})

	/*tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("public/views/index.html", "public/views/base.html"))
	tmpl["index.html"].ExecuteTemplate(res, "base", struct {
		Todos []Todo
		Title string
	}{Todos: todos})*/
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

func deletHandler(res http.ResponseWriter, req *http.Request) {
	todoTitle := req.URL.Path[len("/delete/"):]
	var tmpTodos []Todo
	for _, value := range todos {
		
		if string(value.Title) != string(todoTitle) {

			tmpTodos = append(tmpTodos, Todo{Title: string(value.Title)})

		}
	}
	todos = tmpTodos
	http.Redirect(res, req, "/", http.StatusFound)

}

func saveHandler(res http.ResponseWriter, req *http.Request) {

	//Get todo
	todoTitle := req.FormValue("todo")
	prevTodo := req.FormValue("prev_todo")

	if len(prevTodo) > 0 {
		var tmpTodos []Todo
		var toAppend bool = true
		for _, value := range todos {

			if string(value.Title) == string(prevTodo) {
				tmpTodos = append(tmpTodos, Todo{Title: todoTitle})
				toAppend = false
			} else {
				tmpTodos = append(tmpTodos, Todo{Title: string(value.Title)})
			}
		}
		todos = tmpTodos
		if toAppend == true {

			todos = append(todos, Todo{Title: todoTitle})
		}
	} else {
		todos = append(todos, Todo{Title: todoTitle})
	}

	http.Redirect(res, req, "/", http.StatusFound)

}

func main() {


	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("public/vendor"))))


	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	http.HandleFunc("/delete/", deletHandler)


	http.ListenAndServe(":8080", nil)

}
