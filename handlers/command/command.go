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
	ErrLimitReached = errors.New("limit for this plate has been reached")
)

type OrderRequest struct {
	PlateID  int `db:"plate_id"`
	Quantity int
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
	var req []OrderRequest
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
	if ok, err := checkValidCommand(tx, req, int(claims.SessionID.Int64)); err != nil {
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
			PlateID:   command.PlateID,
			At:        at,
			Quantity:  command.Quantity,
			Status:    models.CommandOrdered,
		})
	}
	_, err = tx.NamedExec(
		`insert into command
				(session_id, plate_id, at, quantity, status) values
				(:session_id, :plate_id, :at, :quantity, :status)`,
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
	w.WriteHeader(http.StatusCreated)
}

func ReadClientCommandHistory(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read client command history: %v", err)
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

func ReadCommandHistory(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("read command history: %v", err)
	}
	command := []models.Command{}
	err := u_sushi.GetDB().Select(&command, "select * from command")
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

type UpdateOrderStatusRequest struct {
	SessionID int `db:"session_id"`
	PlateID   int `db:"plate_id"`
	At        time.Time
	Status    models.CommandStatus
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	wrapErr := func(err error) error {
		return fmt.Errorf("update order status: %v", err)
	}
	var req UpdateOrderStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}

	_, err = u_sushi.GetDB().NamedExec(
		"update command set status = :status where session_id = :session_id and plate_id = :plate_id and at = :at",
		req,
	)
	if err != nil {
		u_sushi.HttpError(w, http.StatusBadRequest, wrapErr(err))
		return
	}
}
