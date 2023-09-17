package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

func AccountInsert(username, password string) error {
	accountCursor = cursors.AccountCursorTurn()

	_, err := accountCursor.Query(`INSERT INTO public.account(username, password) VALUES ($1, $2);`, username, password)

	if err != nil {
		return fmt.Errorf("Error: when insert account  %s", err)
	}

	return nil
}
