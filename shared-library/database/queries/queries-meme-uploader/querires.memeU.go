package queries

import (
	"database/sql"
	"fmt"
	cursors "shared-library/utils"
)

var memeCursor *sql.DB

func InsertMemePhotoId(photoId string) {
	memeCursor = cursors.TurnMemeCursor()
	res, err := memeCursor.Query(`INSERT INTO public.memephoto(photoid) VALUES ($1);`, photoId)

	if err != nil {
		fmt.Errorf("Error insert photoId %s", err)
	}

	fmt.Println(res)
}
