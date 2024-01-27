package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	hauth "u-sushi/handlers/auth"
	"u-sushi/handlers/command"
	"u-sushi/handlers/image"
	"u-sushi/handlers/plate"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	rr := r.PathPrefix("/client").Subrouter()
	rr.Use(auth.ClientAuthMiddleware)
	rr.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	HandleAuth(rr)
	HandleCommand(rr)
	HandlePlate(rr)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	rr := r.PathPrefix("/auth").Subrouter()
	rr.HandleFunc("/login", hauth.ClientLogin).Methods("POST")
}

func HandleCommand(r *mux.Router) {
	r.HandleFunc("/command", command.CreateCommand).Methods("POST")
	r.HandleFunc("/history", command.ReadClientCommandHistory).Methods("GET")
}

func HandlePlate(r *mux.Router) {
	r.HandleFunc("/plate", plate.ReadPlate).Methods("GET")

	HandleImage(r)
	HandleCategory(r)
	HandleAllergen(r)
	HandleIngredient(r)
	HandlePlateIngredient(r)
}

func HandleImage(r *mux.Router) {
	r.HandleFunc("/image", image.ReadImage).Methods("GET")
}

func HandleCategory(r *mux.Router) {
	r.HandleFunc("/category", plate.ReadCategory).Methods("GET")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", plate.ReadAllergen).Methods("GET")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", plate.ReadIngredient).Methods("GET")
}

func HandlePlateIngredient(r *mux.Router) {
	r.HandleFunc("/plate/{plateID}/ingredient", plate.ReadPlateIngredient).Methods("GET")
}
