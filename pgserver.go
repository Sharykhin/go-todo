package main

import (
	"html/template"
	"net/http"
	"log"
	_ "github.com/lib/pq"
	"database/sql"	
	"fmt"
	"strings"
)

type Todo struct {
	Title string
	Id int
}

var db *sql.DB

func init() {

	


}

func DefaultHandler(res http.ResponseWriter, req *http.Request) {


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


	t.Execute(res,struct {
		TodoItem Todo
		Todos []Todo
		}{Todos:todos} )
}

func EditHandler(res http.ResponseWriter, req *http.Request) {
	
	var id string
	var title string
	var todoId int
	id = req.URL.Path[len("/edit/"):]
	err := db.QueryRow(`SELECT id,title FROM homedevice WHERE id=` + id).Scan(&todoId,&title)
	if err != nil {
		log.Fatal(err)
	}
	t, err := template.ParseFiles("public/pg.html")
	if err != nil {
		log.Fatal(err)
	}
	
	t.Execute(res,struct{
		TodoItem Todo 
		Todos []Todo
		}{TodoItem: Todo{Id:todoId, Title:title}})
	
}

func DeleteHandler(res http.ResponseWriter, req *http.Request) {
	var id string
	id = req.URL.Path[len("/delete/"):]

	if _,err := db.Exec(`DELETE FROM homedevice WHERE id=` + id); err != nil {
		log.Fatal(err)
	}

	http.Redirect(res,req,"/",http.StatusFound)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {

	
	todoTitle := req.FormValue("todo")
	todoId := req.FormValue("todoid")
	
	
	if string(todoId) == "0" {
		
		if _,err := db.Exec(`INSERT INTO homedevice(title) VALUES('` + strings.TrimSpace(todoTitle) + `')`); err != nil {
			log.Fatal(err)
		} 				
		
	} else {
		if _,err := db.Exec(`UPDATE homedevice SET title='` + strings.TrimSpace(todoTitle) + `' WHERE id=`+todoId); err != nil {
			log.Fatal(err)
		} 
	}
	
	http.Redirect(res, req, "/",http.StatusFound)
	fmt.Fprint(res, "Todo title: " + todoTitle + ". and todoId: " + todoId)

}


func init() {
	var err error
	db,err = sql.Open("postgres", "user=test dbname=testdb password=test host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	
		
	http.Handle("/vendor/",http.StripPrefix("/vendor/",http.FileServer(http.Dir("public/vendor"))))

	http.HandleFunc("/",DefaultHandler)
	http.HandleFunc("/save/",SaveHandler)
	http.HandleFunc("/edit/",EditHandler)
	http.HandleFunc("/delete/",DeleteHandler)
	
	http.ListenAndServe(":8080",nil)

}