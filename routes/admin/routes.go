package routes

import (
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/handlers/category"
	"u-sushi/handlers/image"
	"u-sushi/handlers/product"

	"github.com/gorilla/mux"
)

func HandleAll(r *mux.Router) {
	r.Use(u_sushi.ContentTypeApplicationJsonMiddleware)
	r.Use(auth.AdminAuthMiddleware)
	HandleAuth(r)
	HandleProduct(r)
}

func HandleAuth(r *mux.Router) {
	auth.LoadKey()
	sr := r.PathPrefix("/auth").Subrouter()
	sr.HandleFunc("/login", auth.AdminLogin).Methods("POST")
	sr.HandleFunc("/password/{userType}", auth.ChangePassword).Methods("POST")
}

func HandleProduct(r *mux.Router) {
	r.HandleFunc("/product", product.CreateProduct).Methods("POST")
	r.HandleFunc("/product", product.ReadProduct).Methods("GET")
	r.HandleFunc("/product/{id}", product.UpdateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", product.DeleteProduct).Methods("DELETE")

	HandleImage(r)
	HandleCategory(r)
	HandleAllergen(r)
	HandleIngredient(r)
	HandleProductIngredient(r)
}

func HandleImage(r *mux.Router) {
	r.HandleFunc("/image", image.CreateImage).Methods("POST")
	r.HandleFunc("/image", image.ReadImage).Methods("GET")
	r.HandleFunc("/image/{id}", image.UpdateImage).Methods("POST")
	r.HandleFunc("/image/{id}", image.DeleteImage).Methods("DELETE")
}

func HandleCategory(r *mux.Router) {
	r.HandleFunc("/category", category.CreateCategory).Methods("POST")
	r.HandleFunc("/category", category.ReadCategory).Methods("GET")
	r.HandleFunc("/category/{id}", category.UpdateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", category.DeleteCategory).Methods("DELETE")
}

func HandleAllergen(r *mux.Router) {
	r.HandleFunc("/allergen", product.CreateAllergen).Methods("POST")
	r.HandleFunc("/allergen", product.ReadAllergen).Methods("GET")
	r.HandleFunc("/allergen/{id}", product.UpdateAllergen).Methods("POST")
	r.HandleFunc("/allergen/{id}", product.DeleteAllergen).Methods("DELETE")
}

func HandleIngredient(r *mux.Router) {
	r.HandleFunc("/ingredient", product.CreateIngredient).Methods("POST")
	r.HandleFunc("/ingredient", product.ReadIngredient).Methods("GET")
	r.HandleFunc("/ingredient/{id}", product.UpdateIngredient).Methods("POST")
	r.HandleFunc("/ingredient/{id}", product.DeleteIngredient).Methods("DELETE")
}

func HandleProductIngredient(r *mux.Router) {
	r.HandleFunc("/product-ingredient/{productID}", product.CreateProductIngredient).Methods("POST")
	r.HandleFunc("/product-ingredient/{productID}", product.ReadProductIngredient).Methods("GET")
	r.HandleFunc("/product-ingredient/{productID}/{ingredientID}", product.DeleteProductIngredient).Methods("DELETE")
}