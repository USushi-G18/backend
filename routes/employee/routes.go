package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/command"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	rr := r.PathPrefix("/employee").Subrouter()
	rr.Use(auth.EmployeeAuthMiddleware)
	rr.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	HandleAuth(rr)
	HandleOrderStatus(rr)
	HandleCommand(rr)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	rr := r.PathPrefix("/auth").Subrouter()
	rr.HandleFunc("/login", auth.EmployeeLogin).Methods("POST")
}

func HandleOrderStatus(r *mux.Router) {
	r.HandleFunc("/order-status", command.UpdateOrderStatus).Methods("POST")
}

func HandleCommand(r *mux.Router) {
	r.HandleFunc("/command", command.ReadCommandHistory).Methods("GET")
}
