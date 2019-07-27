package general

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	t := struct {
		Test int `json:"test"` // pointer so we can test for field absence
	}{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}

	fmt.Println(t.Test)

	_, err := fmt.Fprintln(w, "hi")

	if err != nil {
		panic(err)
	}
}
