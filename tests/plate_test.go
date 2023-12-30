package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"u-sushi/models"

	"github.com/stretchr/testify/assert"
)

var plate = models.Plate{
	Price:  "2.0",
	Menu:   models.MenuLunch,
	Pieces: 2,
}

func createPlate(name string) int {
	p := plate
	p.Name = name
	categoryID := createCategory(fmt.Sprintf("category-plate-%s", p.Name))
	p.Category = categoryID
	w := executeRequest("POST", "/admin/plate", p)

	var id models.ReturningID
	err := json.Unmarshal(w.Body.Bytes(), &id)
	if err != nil {
		log.Fatalln(err)
	}
	return id.ID
}

func TestCreatePlate(t *testing.T) {
	p := plate
	p.Name = "test-create-plate"
	categoryID := createCategory(fmt.Sprintf("category-plate-%s", p.Name))
	p.Category = categoryID
	w := executeRequest("POST", "/admin/plate", p)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReadPlate(t *testing.T) {
	w := executeRequest("GET", "/admin/plate", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	w = executeRequest("GET", "/client/plate", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePlate(t *testing.T) {
	// create plate
	id := createPlate("test-update-plate")
	url := fmt.Sprintf("/admin/plate/%d", id)

	// update created plate
	p := plate
	p.Name = "test-update-plate-updated"
	categoryID := createCategory(fmt.Sprintf("category-plate-%s", p.Name))
	p.Category = categoryID
	w := executeRequest("PUT", url, p)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePlate(t *testing.T) {
	// create plate
	id := createPlate("test-delete-plate")
	url := fmt.Sprintf("/admin/plate/%d", id)

	// delete created plate
	w := executeRequest("DELETE", url, nil)
	assert.Equal(t, http.StatusOK, w.Code)
}
