package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

func AccountInsert(username, password string) (int, error) {
	accountCursor := cursors.AccountCursorTurn()

	/*
		@ We initiate transaction for the have control of query process
	*/

	tx, err := accountCursor.Begin()
	if err != nil {
		return 0, err
	}

	// Add account and turn its id
	result := tx.QueryRow(`INSERT INTO public.account(username, password) VALUES ($1, $2) RETURNING id;`, username, password)

	/*
		@ We catch the id of the returned query

	*/

	var id int
	if err := result.Scan(&id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("Error when account insert: %s", err)
	}

	tx.Commit()

	return id, nil
}


