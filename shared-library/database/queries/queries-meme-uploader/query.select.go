package queries

import (
	"database/sql"
	"fmt"
	cursors "shared-library/utils"
)

func AllPhotoGet() *sql.Rows {
	memeCursor = cursors.MemeCursorTurn()
	res, err := memeCursor.Query(`SELECT * FROM public.memephoto;`)

	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}

	return res
}

func AllTextGet() *sql.Rows {
	memeCursor = cursors.MemeCursorTurn()
	res, err := memeCursor.Query(`SELECT * FROM public.memetext;`)

	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}

	return res

}
