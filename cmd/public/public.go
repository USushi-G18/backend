package main

import (
	"log"
	"net/http"
	u_sushi "u-sushi"
	routes "u-sushi/routes/public"

	"github.com/gorilla/mux"
)

func main() {
	err := u_sushi.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	routes.HandleAll(r)
	err = http.ListenAndServe(":8082", r)
	if err != nil {
		log.Fatalln(err)
	}
}
