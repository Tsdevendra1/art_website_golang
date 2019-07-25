package models

import "database/sql"

func createUser(db *sql.DB, firstName, email, password string) (*User, error) {
	user := &User{}


	err := db.QueryRow(`
insert into users (first_name, email, password) 
values ($1, $2, $3)
returning id, first_name, email
`).Scan(&user.ID, &user.FirstName, &user.Email)

	if err != nil {
		panic(err)
	}
	return user, err
}
