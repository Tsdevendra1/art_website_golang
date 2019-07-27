package general

import (
	general "artWebsite/general/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	t := general.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}
	fmt.Println(t)

	_, err := fmt.Fprintln(w, "hi")

	if err != nil {
		panic(err)
	}
}
