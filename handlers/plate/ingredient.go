package plate

import (
	"encoding/json"
	"fmt"
	"net/http"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/gorilla/mux"
)

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create ingredient: %v", err)
	}
	var req models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var id int
	err = u_sushi.NamedGet(&id, "insert into ingredient (name, allergen_id) values (:name, :allergen_id) returning id", &req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	idJson, err := json.Marshal(models.ReturningID{ID: id})
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(idJson))
}

func ReadIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read ingredient: %v", err)
	}
	ingredients := []models.Ingredient{}
	err := u_sushi.GetDB().Select(&ingredients, "select * from ingredient")
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	ingredientsJson, err := json.Marshal(ingredients)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(ingredientsJson))
}

func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("update ingredient: %v", err)
	}
	var req models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	vars := mux.Vars(r)
	_, err = u_sushi.GetDB().Exec("update ingredient set name = $1, allergen_id = $2 where id = $3", req.Name, req.AllergenID, vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}

func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete ingredient: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().Exec("delete from ingredient where id = $1", vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
