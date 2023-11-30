package command

import (
	u_sushi "u-sushi"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

func checkCanOrderMore(tx *sqlx.Tx, command []CommandRequest, sessionID int) (bool, error) {
	var seating int
	err := tx.QueryRow("select seating from session where id = $1", sessionID).Scan(&seating)
	if err != nil {
		return false, err
	}

	var prevCommand []CommandRequest
	err = tx.Select(&prevCommand, "select product_id, quantity from command where session_id = $1", sessionID)
	if err != nil {
		return false, err
	}

	count := make(map[int]int)
	for _, c := range command {
		count[c.ProductID] += c.Quantity
	}
	for _, c := range prevCommand {
		count[c.ProductID] += c.Quantity
	}

	var products []struct {
		ID    int
		Limit null.Int
	}
	err = u_sushi.GetDB().Select(&products, "select distinct p.id, p.order_limit as limit from product p join command c on c.product_id = p.id where session_id = $1", sessionID)
	if err != nil {
		return false, err
	}

	for _, p := range products {
		if p.Limit.Valid && int(p.Limit.Int64)*seating < count[p.ID] {
			return false, nil
		}
	}

	return true, nil
}
