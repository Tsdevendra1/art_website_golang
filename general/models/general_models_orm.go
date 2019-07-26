package models

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) CreateUser(firstName, email, password string) (*User, error) {
	user := &User{}

	err := db.DB.QueryRow(fmt.Sprintf(`
insert into users (first_name, email, password)
values ($1, $2, crypt($3, gen_salt('bf')))
returning %s
`, FieldToString(user)), firstName, email, password).Scan(user.FieldsToUpdate())

	if err != nil {
		panic(err)
	}
	return user, err
}

func (db *Database) SelectUser(queryString string, queryStringValues ...interface{}) (*User, error) {
	user := &User{}
	err := db.DB.QueryRow(fmt.Sprintf(`
	select %s from users where %s
`, FieldToString(user), queryString), queryStringValues).Scan(user.FieldsToUpdate())

	if err != nil {
		panic(err)
	}

	return user, nil

}

func (db *Database) FindUserByFirstName(firstName string) (*User, error) {
	return db.SelectUser("lower(first_name) = lower($1)", firstName)
}
