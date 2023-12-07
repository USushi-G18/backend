package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/command"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.EmployeeAuthMiddleware)
	sr := r.PathPrefix("/employee").Subrouter()
	HandleAuth(sr)
	HandleOrderStatus(sr)
	HandleCommand(sr)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	sr := r.PathPrefix("/auth").Subrouter()
	sr.HandleFunc("/login", auth.EmployeeLogin).Methods("POST")
}

func HandleOrderStatus(r *mux.Router) {
	r.HandleFunc("/order-status", command.UpdateOrderStatus).Methods("GET")
}

func HandleCommand(r *mux.Router) {
	r.HandleFunc("/command", command.ReadCommandHistory).Methods("GET")
}
