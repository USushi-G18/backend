package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/category"
	"u-sushi/handlers/command"
	"u-sushi/handlers/image"
	"u-sushi/handlers/product"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.ClientAuthMiddleware)
	sr := r.PathPrefix("/admin").Subrouter()
	HandleAuth(sr)
	HandleCommand(sr)
	HandleProduct(sr)
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

func HandleProduct(r *mux.Router) {
	r.HandleFunc("/product", product.ReadProduct).Methods("GET")

	HandleImage(r)
	HandleCategory(r)
	HandleAllergen(r)
	HandleIngredient(r)
	HandleProductIngredient(r)
}

func HandleImage(r *mux.Router) {
	r.HandleFunc("/image", image.ReadImage).Methods("GET")
}

func HandleCategory(r *mux.Router) {
	r.HandleFunc("/category", category.ReadCategory).Methods("GET")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", product.ReadAllergen).Methods("GET")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", product.ReadIngredient).Methods("GET")
}

func HandleProductIngredient(r *mux.Router) {
	r.HandleFunc("/product-ingredient/{productID}", product.ReadProductIngredient).Methods("GET")
}
