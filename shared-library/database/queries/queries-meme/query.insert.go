package queries

import (
	"database/sql"
	"fmt"
	cursors "shared-library/utils"
)

var memeCursor *sql.DB

func PhotoIdInsert(photoId string) error {
	memeCursor = cursors.MemeCursorTurn()

	tx, err := memeCursor.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Query(`INSERT INTO public.memephoto(photoid) VALUES ($1);`, photoId)

	if err != nil {
		return fmt.Errorf("Error: when insert photo id %s", err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func TextInsert(text string) error {
	memeCursor = cursors.MemeCursorTurn()

	tx, err := memeCursor.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Query(`INSERT INTO public.memetext(text) VALUES ($1);`, text)
	if err != nil {
		return fmt.Errorf("Error: when insert text %s", err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
