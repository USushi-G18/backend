package plate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/gorilla/mux"
)

func CreatePlate(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create plate: %v", err)
	}
	var req models.Plate
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	_, err = u_sushi.GetDB().NamedExec(`
		insert into plate 
		(name, price, category, menu, description, image_id, order_limit, pieces) 
		values (:name, :price, :category, :menu, :description, :image_id, :order_limit, :pieces)`,
		&req,
	)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ReadPlate(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read plate: %v", err)
	}
	plates := []models.Plate{}
	err := u_sushi.GetDB().Select(&plates, "select * from plate")
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	platesJson, err := json.Marshal(plates)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(platesJson))
}

func UpdatePlate(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("update plate: %v", err)
	}
	var req models.Plate
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	req.ID = id
	_, err = u_sushi.GetDB().NamedExec(`
		update plate set
			name = :name,
			price = :price,
			category = :category,
			menu = :menu,
			description = :description,
			image_id = :image_id,
			order_limit = :order_limit,
			pieces = :pieces
		where id = :id`,
		&req,
	)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}

func DeletePlate(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete plate: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().Exec("delete from plate where id = $1", vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}