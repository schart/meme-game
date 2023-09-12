package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

func AllPhotoGet() {
	memeCursor = cursors.TurnMemeCursor()
	res, err := memeCursor.Query(`SELECT * FROM public.memephoto;`)

	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}

	fmt.Println(res)

}

func AllTextGet() {
	memeCursor = cursors.TurnMemeCursor()
	res, err := memeCursor.Query(`SELECT * FROM public.memetext;`)

	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}

	fmt.Println(res)

}
