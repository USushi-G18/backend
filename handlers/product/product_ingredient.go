package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/gorilla/mux"
)

type CreateProductIngredientRequest struct {
	IngredientID int
}

func CreateProductIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create product_ingredient: %v", err)
	}
	var req CreateProductIngredientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	vars := mux.Vars(r)
	_, err = u_sushi.GetDB().Exec("insert into product_ingredient (product_id, ingredient_id) values ($1, $2)", vars["productID"], req.IngredientID)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}

func ReadProductIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read product_ingredient: %v", err)
	}
	vars := mux.Vars(r)
	ingredients := []models.Ingredient{}
	err := u_sushi.GetDB().Select(&ingredients, `
		select i.* from product_ingredient pi 
			join ingredient i on pi.ingredient_id = i.id
		where pi.product_id = $1`,
		vars["productID"],
	)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	ingredientsJson, err := json.Marshal(&ingredients)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(ingredientsJson))
}

func DeleteProductIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete product_ingredient: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().NamedExec("delete from product_ingredient where product_id = :product_id and ingredient_id = :ingredient_id", &vars)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
