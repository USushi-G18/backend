package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	u_sushi "u-sushi"
	"u-sushi/auth"
	"u-sushi/models"
)

var (
	ErrLimitReached = errors.New("limit for this product has been reached")
)

type CommandRequest struct {
	ProductID int `db:"product_id"`
	Quantity  int
}

func CreateCommand(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("create command: %v", err)
	}
	claims, err := auth.ExtractClaims(r)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	var req []CommandRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	tx, err := u_sushi.GetDB().Beginx()
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	if ok, err := checkCanOrderMore(tx, req, int(claims.SessionID.Int64)); err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	} else if !ok {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(ErrLimitReached))
		return
	}
	at := time.Now()
	var commands []models.Command
	for _, command := range req {
		commands = append(commands, models.Command{
			SessionID: int(claims.SessionID.Int64),
			ProductID: command.ProductID,
			At:        at,
			Quantity:  command.Quantity,
			Status:    models.CommandOrdered,
		})
	}
	_, err = tx.NamedExec(
		`insert into command
				(session_id, product_id, at, quantity, status) values
				(:session_id, :product_id, :at, :quantity, :status)`,
		commands)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	err = tx.Commit()
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
}

func CommandHistory(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("command history: %v", err)
	}
	claims, err := auth.ExtractClaims(r)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
	command := []models.Command{}
	err = u_sushi.GetDB().Select(&command, "select * from command where session_id = $1", claims.SessionID.Int64)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	commandJson, err := json.Marshal(command)
	if err != nil {
		u_sushi.HttpError(w, http.StatusInternalServerError, wrapErr(err))
		return
	}
	fmt.Fprint(w, string(commandJson))
}
