package main

import (
	"html/template"
	"net/http"
	"log"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

type Todo struct {
	Title string
	Id int
}

var db *sql.DB

func init() {

	


}

func DefaultHandler(res http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("postgres", "user=test dbname=testdb password=test host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("public/pg.html")
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.Query("SELECT id,title FROM homedevice")
	if err != nil {
	    log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
	    var (
	    	title string
	    	id int
	    	todo Todo
	    	)

	    if err := rows.Scan(&id,&title); err != nil {
	        log.Fatal(err)
	    }

	    todo = Todo{Id:id,Title:title}
	    todos = append(todos,todo)
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}


	t.Execute(res,todos)
}

func EditHandler(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "user=test dbname=testdb password=test host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}

	var id string
	var title string
	id = req.URL.Path[len("/edit/"):]
	err = db.QueryRow(`SELECT title FROM homedevice WHERE id=` + id).Scan(&title)
	t, err := template.ParseFiles("public/pg.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)
	t.Execute(res,struct{
		TodoItem Todo 
		}{TodoItem: Todo{Id:2, Title:title}})
	
}

func main() {
	
		
	http.Handle("/vendor/",http.StripPrefix("/vendor/",http.FileServer(http.Dir("public/vendor"))))

	http.HandleFunc("/",DefaultHandler)
	http.HandleFunc("/edit/",EditHandler)
	http.ListenAndServe(":8080",nil)

}