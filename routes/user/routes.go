package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/command"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.ClientAuthMiddleware)
	HandleAuth(r)
	HandleCommand(r)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	sr := r.PathPrefix("/auth").Subrouter()
	sr.HandleFunc("/login", auth.ClientLogin).Methods("POST")
}

func HandleCommand(r *mux.Router) {
	r.HandleFunc("/command", command.CreateCommand).Methods("POST")
	r.HandleFunc("/history", command.CommandHistory).Methods("GET")
}
