package models

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

func createUser(db *sql.DB, firstName, email, password string) (*User, error) {
	user := &User{}

	err := db.QueryRow(fmt.Sprintf(`
insert into users (first_name, email, password)
values ($1, $2, crypt($3, gen_salt('bf')))
returning %s
`, FieldToString(user)), firstName, email, password).Scan(user.FieldsToUpdate())

	if err != nil {
		panic(err)
	}
	return user, err
}
