package main

import (
	"log"
	"net/http"
	"artWebsite/general/router"
)

func main() {

	mainRouter := router.NewRouter()
	mainRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	db := setupDb()
	defer closeDb(db)
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
