package image

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/gorilla/mux"
)

func CreateImage(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create image: %v", err)
	}
	var req models.Image
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var id int
	err = u_sushi.NamedGet(&id, "insert into image (image) values (:image) returning id", &req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	idJson, err := json.Marshal(models.ReturningID{ID: id})
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(idJson))
}

func ReadImage(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read image: %v", err)
	}
	var err error

	limitS := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitS)
	offsetS := r.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(offsetS)

	images := []models.Image{}
	if limit == 0 {
		err = u_sushi.GetDB().Select(&images, "select * from image order by id offset $1", offset)
	} else {
		err = u_sushi.GetDB().Select(&images, "select * from image order by id offset $1 limit $2", offset, limit)
	}
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	imagesJson, err := json.Marshal(images)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(imagesJson))
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("update image: %v", err)
	}
	var req models.Image
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	vars := mux.Vars(r)
	_, err = u_sushi.GetDB().Exec("update image set image = $1 where id = $2", req.Image, vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("delete image: %v", err)
	}
	vars := mux.Vars(r)
	_, err := u_sushi.GetDB().Exec("delete from image where id = $1", vars["id"])
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
