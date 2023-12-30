package tests

import (
	"fmt"
	"net/http"
	"testing"
	"u-sushi/auth"
	"u-sushi/models"

	"github.com/stretchr/testify/assert"
)

func TestChangeAdminPassword(t *testing.T) {
	testChangePassword(t, models.UserAdmin)
}

func TestChangeClientPassword(t *testing.T) {
	testChangePassword(t, models.UserClient)
}

func TestChangeEmployeePassword(t *testing.T) {
	testChangePassword(t, models.UserEmployee)
}

func testChangePassword(t *testing.T, userType models.UserType) {
	url := fmt.Sprintf("/admin/auth/password/%s", userType)
	w := executeRequest("POST", url, auth.ChangePasswordRequest{
		OldPassword: "u-sushi",
		NewPassword: "u-sushi",
	})
	assert.Equal(t, http.StatusOK, w.Code)
}
