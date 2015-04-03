package main


import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	)

func main() {
	fmt.Println("Hello world")

	db, err := sql.Open("postgres", "user=pimples dbname=pimples host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.Query("SELECT * FROM playground")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	fmt.Printf("type is %T and \n",db)
}