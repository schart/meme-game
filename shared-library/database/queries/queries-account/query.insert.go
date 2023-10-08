package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

func AccountInsert(username, password string) (int, error) {
	accountCursor := cursors.AccountCursorTurn()

	// Add account and turn its id
	row := accountCursor.QueryRow(`INSERT INTO public.account(username, password) VALUES ($1, $2) RETURNING id;`, username, password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Error register: %s", err)
	}

	return id, nil
}
