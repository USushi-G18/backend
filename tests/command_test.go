package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"u-sushi/handlers/command"
	"u-sushi/models"

	"github.com/stretchr/testify/assert"
)

func createCommandRequest(i int) []command.OrderRequest {
	var c []command.OrderRequest
	for j := 0; j < 3; j++ {
		id := createPlate(fmt.Sprintf("create-plate-for-order-test-%d-%d", i, j))
		c = append(c, command.OrderRequest{
			PlateID:  id,
			Quantity: 1,
		})
	}
	return c
}

func TestCreateCommand(t *testing.T) {
	commandReq := createCommandRequest(0)
	w := executeRequest("POST", "/client/command", commandReq)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReadCommand(t *testing.T) {
	w := executeRequest("GET", "/employee/command", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateOrderStatus(t *testing.T) {
	// create some commands
	for i := 0; i < 2; i++ {
		commandReq := createCommandRequest(i + 1)
		_ = executeRequest("POST", "/client/command", commandReq)
	}

	// read the commands
	w := executeRequest("GET", "/employee/command", nil)
	var commands []models.Command
	err := json.Unmarshal(w.Body.Bytes(), &commands)
	if err != nil {
		log.Fatalln(err)
	}

	// update the commands
	for _, c := range commands {
		w := executeRequest("POST", "/employee/order-status", command.UpdateOrderStatusRequest{
			SessionID: c.SessionID,
			PlateID:   c.PlateID,
			At:        c.At,
			Status:    models.CommandPrepared,
		})
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
