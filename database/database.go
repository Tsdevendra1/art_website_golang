package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tharukadevendra"
	password = ""
	dbname   = "golang"
)

type Database struct {
	DB *sql.DB
}

func (database Database) CloseDb() {
	if err := database.DB.Close(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully closed database")
}

func SetupDb() *Database {

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
	return &Database{db}
}

var (
	// DBCon is the connection handle
	// for the database
	DBConn *Database
)

