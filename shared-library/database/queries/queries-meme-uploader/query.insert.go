package queries

import (
	"database/sql"
	"fmt"
	cursors "shared-library/utils"
)

var memeCursor *sql.DB

func InsertMemePhotoId(photoId string) error {
	memeCursor = cursors.TurnMemeCursor()
	_, err := memeCursor.Query(`INSERT INTO public.memephoto(photoid) VALUES ($1);`, photoId)

	if err != nil {
		return fmt.Errorf("Error insert photoId %s", err)
	}

	return nil
}

func InsertMemeText(text string) error {
	memeCursor = cursors.TurnMemeCursor()

	_, err := memeCursor.Query(`INSERT INTO public.memetext(text) VALUES ($1);`, text)
	if err != nil {
		return fmt.Errorf("Error insert text %s", err)
	}

	return nil
}
