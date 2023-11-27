package u_sushi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HttpError(w http.ResponseWriter, statusCode int, httpErr error) {
	jsonErr, err := json.Marshal(&struct {
		Error string `json:"error"`
	}{
		Error: httpErr.Error(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(httpErr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, string(jsonErr))
}
