package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	hauth "u-sushi/handlers/auth"
	"u-sushi/handlers/image"
	"u-sushi/handlers/plate"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	rr := r.PathPrefix("/admin").Subrouter()
	rr.Use(auth.AdminAuthMiddleware)
	rr.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	HandleAuth(rr)
	HandlePlate(rr)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	rr := r.PathPrefix("/auth").Subrouter()
	rr.HandleFunc("/login", hauth.AdminLogin).Methods("POST")
	rr.HandleFunc("/password/{userType}", hauth.ChangePassword).Methods("POST")
}

func HandlePlate(r *mux.Router) {
	r.HandleFunc("/plate", plate.CreatePlate).Methods("POST")
	r.HandleFunc("/plate", plate.ReadPlate).Methods("GET")
	r.HandleFunc("/plate/{id}", plate.UpdatePlate).Methods("PUT")
	r.HandleFunc("/plate/{id}", plate.DeletePlate).Methods("DELETE")

	HandleImage(r)
	HandleCategory(r)
	HandleAllergen(r)
	HandleIngredient(r)
	HandlePlateIngredient(r)
}

func HandleImage(r *mux.Router) {
	r.HandleFunc("/image", image.CreateImage).Methods("POST")
	r.HandleFunc("/image", image.ReadImage).Methods("GET")
	r.HandleFunc("/image/{id}", image.UpdateImage).Methods("PUT")
	r.HandleFunc("/image/{id}", image.DeleteImage).Methods("DELETE")
}

func HandleCategory(r *mux.Router) {
	r.HandleFunc("/category", plate.CreateCategory).Methods("POST")
	r.HandleFunc("/category", plate.ReadCategory).Methods("GET")
	r.HandleFunc("/category/{id}", plate.UpdateCategory).Methods("PUT")
	r.HandleFunc("/category/{id}", plate.DeleteCategory).Methods("DELETE")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", plate.ReadAllergen).Methods("GET")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", plate.ReadIngredient).Methods("GET")
}

func HandlePlateIngredient(r *mux.Router) {
	r.HandleFunc("/plate/{plateID}/ingredient", plate.CreatePlateIngredient).Methods("POST")
	r.HandleFunc("/plate/{plateID}/ingredient", plate.ReadPlateIngredient).Methods("GET")
	r.HandleFunc("/plate/{plateID}/ingredient/{ingredientID}", plate.DeletePlateIngredient).Methods("DELETE")
}
