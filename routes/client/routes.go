package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/category"
	"u-sushi/handlers/command"
	"u-sushi/handlers/image"
	"u-sushi/handlers/plate"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.ClientAuthMiddleware)
	sr := r.PathPrefix("/admin").Subrouter()
	HandleAuth(sr)
	HandleCommand(sr)
	HandlePlate(sr)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	sr := r.PathPrefix("/auth").Subrouter()
	sr.HandleFunc("/login", auth.ClientLogin).Methods("POST")
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
	r.HandleFunc("/category", category.ReadCategory).Methods("GET")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", plate.ReadAllergen).Methods("GET")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", plate.ReadIngredient).Methods("GET")
}

func HandlePlateIngredient(r *mux.Router) {
	r.HandleFunc("/plate-ingredient/{plateID}", plate.ReadPlateIngredient).Methods("GET")
}
