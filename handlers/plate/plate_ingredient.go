package plate

import (
	"encoding/json"
	"fmt"
	"net/http"
	u_sushi "u-sushi"

	"github.com/gorilla/mux"
)

type plateIngredientID struct {
	IngredientID int `json:"ingredientID" db:"ingredient_id"`
}

func CreatePlateIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create plate_ingredient: %v", err)
	}
	var req plateIngredientID
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	vars := mux.Vars(r)
	_, err = u_sushi.GetDB().Exec("insert into plate_ingredient (plate_id, ingredient_id) values ($1, $2)", vars["plateID"], req.IngredientID)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ReadPlateIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read plate_ingredient: %v", err)
	}
	vars := mux.Vars(r)
	ingredients := []plateIngredientID{}
	err := u_sushi.GetDB().Select(&ingredients, `
		select pi.ingredient_id from plate_ingredient pi 
			join ingredient i on pi.ingredient_id = i.id
		where pi.plate_id = $1`,
		vars["plateID"],
	)
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

func DeletePlateIngredient(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete plate_ingredient: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().Exec("delete from plate_ingredient where plate_id = $1 and ingredient_id = $2", vars["plateID"], vars["ingredientID"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
