package main

import (
	"log"
	"net/http"
	"artWebsite/general/router"
)

func main() {

	mainRouter := router.NewRouter()
	db := setupDb()
	defer closeDb(db)
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
