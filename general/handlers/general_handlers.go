package general

import (
	general "artWebsite/general/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	userJson := general.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userJson); err != nil {
		panic(err)
	}
	fmt.Println(userJson.FirstName, userJson.Email)

	user, err := general.CreateUser(userJson.FirstName, userJson.Email, userJson.Password)

	fmt.Println(user)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintln(w, "hi")

	if err != nil {
		panic(err)
	}
}

func UserSelect(w http.ResponseWriter, r *http.Request) {

	user, err := general.SelectUser("ID = $1", 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
