package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/models"

	"github.com/gorilla/mux"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/guregu/null.v4"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Password string
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	loginUser(w, r, models.UserAdmin)
}

func EmployeeLogin(w http.ResponseWriter, r *http.Request) {
	loginUser(w, r, models.UserEmployee)
}

func loginUser(w http.ResponseWriter, r *http.Request, userType models.UserType) {
	wrapErr := func(err error) error {
		return fmt.Errorf("login: %v", err)
	}
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var encodedHash string
	err = u_sushi.GetDB().QueryRow("select password from sushi_user where user_type = $1", userType).Scan(&encodedHash)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	if verified, err := auth.VerifyPassword(req.Password, encodedHash); err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	} else if !verified {
		u_sushi.HttpError(w, http.StatusUnauthorized, wrapErr(auth.ErrWrongPassword))
		return
	}
	token, err := auth.CreateJWT(userType, null.Int{})
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

type ClientLoginRequest struct {
	Password    string
	TableNumber int
	Menu        models.Menu
	Seating     int
}

func ClientLogin(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("login: %v", err)
	}
	var req ClientLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var encodedHash string
	err = u_sushi.GetDB().QueryRow("select password from sushi_user where user_type = $1", models.UserClient).Scan(&encodedHash)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	if verified, err := auth.VerifyPassword(req.Password, encodedHash); err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	} else if !verified {
		u_sushi.HttpError(w, http.StatusUnauthorized, wrapErr(auth.ErrWrongPassword))
		return
	}
	var sessionID int
	err = u_sushi.GetDB().QueryRow(
		"insert into session (start_at, table_number, menu, seating) values ($1, $2, $3, $4) returning id",
		time.Now(),
		req.TableNumber,
		req.Menu,
		req.Seating,
	).Scan(&sessionID)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	token, err := auth.CreateJWT(models.UserClient, null.IntFrom(int64(sessionID)))
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

type ChangePasswordRequest struct {
	OldPassword string
	NewPassword string
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("change password: %v", err)
	}
	var req ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var encodedHash string
	vars := mux.Vars(r)
	userType := cases.Title(language.English).String(strings.ToLower(vars["userType"]))
	err = u_sushi.GetDB().QueryRow("select password from sushi_user where user_type = $1", userType).Scan(&encodedHash)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	verified, err := auth.VerifyPassword(req.OldPassword, encodedHash)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	} else if !verified {
		u_sushi.HttpError(w, http.StatusUnauthorized, wrapErr(auth.ErrWrongPassword))
		return
	}
	newHash, err := auth.HashPassword(req.NewPassword, auth.StdArgon2Params)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	_, err = u_sushi.GetDB().Exec("update sushi_user set password = $1 where user_type = $2", newHash, userType)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
}
