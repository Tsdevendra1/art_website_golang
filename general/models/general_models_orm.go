package general

import (
	"artWebsite/database"
	"fmt"
)

func CreateUser(firstName, email, password string) (*User, error) {
	user := &User{}

	fmt.Println("hi")
	fmt.Println(FieldToString(user))
	fmt.Println("hi")
	err := database.DBConn.DB.QueryRow(`
insert into users (first_name, email, password)
values ($1, $2, crypt($3, gen_salt('bf')))
returning id, first_name, email
`, firstName, email, password).Scan(user.FieldsToUpdate()...)

	if err != nil {
		panic(err)
	}
	return user, err
}

func SelectUser(queryString string, queryStringValues ...interface{}) (*User, error) {
	user := &User{}
	err := database.DBConn.DB.QueryRow(fmt.Sprintf(`
	select id, first_name, email from users where %s
`, queryString), queryStringValues...).Scan(user.FieldsToUpdate()...)

	if err != nil {
		panic(err)
	}

	return user, nil

}

func FindUserByFirstName(firstName string) (*User, error) {
	return SelectUser("lower(first_name) = lower($1)", firstName)
}
