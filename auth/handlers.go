package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	u_sushi "u-sushi"
	"u-sushi/models"

	"gopkg.in/guregu/null.v4"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Password string
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("login: %v", err)
	}
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var encodedHash struct {
		EncodedHash string `db:"password"`
	}
	err = u_sushi.GetDB().Get(&encodedHash, "select password from sushi_user where user_type = $1", models.UserAdmin)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	verified, err := VerifyPassword(req.Password, encodedHash.EncodedHash)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	} else if !verified {
		u_sushi.HttpError(w, http.StatusUnauthorized, fmt.Errorf("wrong password"))
		return
	}
	token, err := CreateJWT(models.UserAdmin, null.Int{})
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	tokenJson, err := json.Marshal(TokenResponse{
		Token: token,
	})
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(tokenJson))
}
