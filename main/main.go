package main

import (
	"artWebsite/database"
	"artWebsite/general/router"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	mainRouter := general.NewRouter()
	mainRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	database.DBConn = database.SetupDb()
	defer database.DBConn.CloseDb()
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
