package command

import (
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

func checkValidCommand(tx *sqlx.Tx, command []OrderRequest, sessionID int) (bool, error) {
	var session models.Session
	err := tx.Get(&session, "select * from session where id = $1", sessionID)
	if err != nil {
		return false, err
	}

	if session.Menu == models.MenuCarte {
		return true, nil
	}

	var prevCommand []OrderRequest
	err = tx.Select(&prevCommand, "select plate_id, quantity from command where session_id = $1", sessionID)
	if err != nil {
		return false, err
	}

	count := make(map[int]int)
	for _, c := range command {
		count[c.PlateID] += c.Quantity
	}
	for _, c := range prevCommand {
		count[c.PlateID] += c.Quantity
	}

	var plates []struct {
		ID    int
		Limit null.Int
	}
	err = u_sushi.GetDB().Select(
		&plates,
		`select distinct p.id, p.order_limit as limit 
			from plate p 
			join command c on c.plate_id = p.id 
		where session_id = $1`,
		sessionID,
	)
	if err != nil {
		return false, err
	}

	for _, p := range plates {
		if p.Limit.Valid && int(p.Limit.Int64)*session.Seating < count[p.ID] {
			return false, nil
		}
	}

	return true, nil
}
