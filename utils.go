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
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, string(jsonErr))
}

func ContentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}
