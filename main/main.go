package main

import (
	"artWebsite/general/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "tharukadevendra"
	password = ""
	dbname   = "golang"
)

func closeDb(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully closed database")
}

func setupDb() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func main() {

	mainRouter := router.NewRouter()
	mainRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	db := setupDb()
	defer closeDb(db)
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
