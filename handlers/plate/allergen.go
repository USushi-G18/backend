package plate

import (
	"encoding/json"
	"fmt"
	"net/http"
	u_sushi "u-sushi"
	"u-sushi/models"
)

func ReadAllergen(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read allergen: %v", err)
	}
	allergens := []models.Allergen{}
	err := u_sushi.GetDB().Select(&allergens, "select * from allergen")
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	allergensJson, err := json.Marshal(allergens)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(allergensJson))
}
