package general

import "encoding/json"

type User struct {
	ID        int
	Name      string
	Email     string
	FirstName string
	LastName  string
	Password string
}

func FieldToString(structToConvert interface{}) string {
	out, err := json.Marshal(structToConvert)
	if err != nil {
		panic(err)
	}
	return string(out)

}

func (user *User) FieldsToUpdate() []interface{} {
	return []interface{}{&user.ID, &user.FirstName, &user.Email}
}
