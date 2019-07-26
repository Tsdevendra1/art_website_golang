package general

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
	_, err := fmt.Fprintln(w, "hi")

	if err != nil {
		panic(err)
	}
}
