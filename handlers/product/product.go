package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create product: %v", err)
	}
	var req models.Product
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	_, err = u_sushi.GetDB().NamedExec(`
		insert into product 
		(name, price, category, menu, description, image_id, order_limit, pieces) 
		values (:name, :price, :category, :menu, :description, :image_id, :order_limit, :pieces)`,
		&req,
	)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}

func ReadProduct(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read product: %v", err)
	}
	products := []models.Product{}
	err := u_sushi.GetDB().Select(&products, "select * from product")
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	productsJson, err := json.Marshal(products)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(productsJson))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("update product: %v", err)
	}
	var req models.Product
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
		update product set
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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete product: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().Exec("delete from product where id = $1", vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
