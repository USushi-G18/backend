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

func createCategory(name string) int {
	w := executeRequest("POST", "/admin/category", models.Category{
		Name: name,
	})

	var id models.ReturningID
	err := json.Unmarshal(w.Body.Bytes(), &id)
	if err != nil {
		log.Fatalln(err)
	}
	return id.ID
}

func TestCreateCategory(t *testing.T) {
	w := executeRequest("POST", "/admin/category", models.Category{
		Name: "test-create-category",
	})
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReadCategory(t *testing.T) {
	w := executeRequest("GET", "/admin/category", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	w = executeRequest("GET", "/client/category", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCategory(t *testing.T) {
	// create category
	id := createCategory("test-update-category")
	url := fmt.Sprintf("/admin/category/%d", id)

	// update created category
	w := executeRequest("PUT", url, models.Category{
		Name: "test-update-category-updated",
	})
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCategory(t *testing.T) {
	// create category
	id := createCategory("test-delete-category")
	url := fmt.Sprintf("/admin/category/%d", id)

	// delete created category
	w := executeRequest("DELETE", url, nil)
	assert.Equal(t, http.StatusOK, w.Code)
}
