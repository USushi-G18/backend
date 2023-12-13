package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/category"
	"u-sushi/handlers/image"
	"u-sushi/handlers/plate"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.AdminAuthMiddleware)
	HandleAuth(r)
	HandlePlate(r)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	sr := r.PathPrefix("/auth").Subrouter()
	sr.HandleFunc("/login", auth.AdminLogin).Methods("POST")
	sr.HandleFunc("/password/{userType}", auth.ChangePassword).Methods("POST")
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
	r.HandleFunc("/category", category.CreateCategory).Methods("POST")
	r.HandleFunc("/category", category.ReadCategory).Methods("GET")
	r.HandleFunc("/category/{id}", category.UpdateCategory).Methods("PUT")
	r.HandleFunc("/category/{id}", category.DeleteCategory).Methods("DELETE")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", plate.CreateAllergen).Methods("POST")
	r.HandleFunc("/allergen", plate.ReadAllergen).Methods("GET")
	r.HandleFunc("/allergen/{id}", plate.UpdateAllergen).Methods("PUT")
	r.HandleFunc("/allergen/{id}", plate.DeleteAllergen).Methods("DELETE")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", plate.CreateIngredient).Methods("POST")
	r.HandleFunc("/ingredient", plate.ReadIngredient).Methods("GET")
	r.HandleFunc("/ingredient/{id}", plate.UpdateIngredient).Methods("PUT")
	r.HandleFunc("/ingredient/{id}", plate.DeleteIngredient).Methods("DELETE")
}

func HandlePlateIngredient(r *mux.Router) {
	r.HandleFunc("/plate-ingredient/{plateID}", plate.CreatePlateIngredient).Methods("POST")
	r.HandleFunc("/plate-ingredient/{plateID}", plate.ReadPlateIngredient).Methods("GET")
	r.HandleFunc("/plate-ingredient/{plateID}/{ingredientID}", plate.DeletePlateIngredient).Methods("DELETE")
}
