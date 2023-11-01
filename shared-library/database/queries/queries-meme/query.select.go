package queries

import (
	"fmt"
	cursors "shared-library/utils"
	"strconv"
)

func PhotoGetByCount(count int) map[string]interface{} {
	memeCursor = cursors.MemeCursorTurn()

	rows, err := memeCursor.Query(`SELECT * FROM public.memephoto LIMIT $1;`, count)
	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}
	defer rows.Close()

	var counter int = 1
	var cards map[string]interface{} = map[string]interface{}{}

	for rows.Next() {
		var id string
		var photoid string

		if err := rows.Scan(&id, &photoid); err != nil {
			fmt.Errorf("Error scanning row: %s", err)
			return nil
		}

		cards["card:"+strconv.Itoa(counter)] = photoid
		counter += 1

	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error iterating over rows: %s", err)
		return nil
	}

	return cards
}

func TextGetByCount(count int) map[string]interface{} {
	memeCursor = cursors.MemeCursorTurn()

	rows, err := memeCursor.Query(`SELECT * FROM public.memetext LIMIT $1;`, count)
	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}
	defer rows.Close()

	var counter int = 1
	var cards map[string]interface{} = map[string]interface{}{}

	for rows.Next() {
		var id string
		var text string

		if err := rows.Scan(&id, &text); err != nil {
			fmt.Errorf("Error scanning row: %s", err)
			return nil
		}

		cards["text:"+strconv.Itoa(counter)] = text
		counter += 1

	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error iterating over rows: %s", err)
		return nil
	}

	return cards
}

func IsTherePhoto(informations map[string]interface{}) bool {
	memeCursor = cursors.MemeCursorTurn()

	row := memeCursor.QueryRow(`SELECT id FROM public.memephoto WHERE photoid = $1;`, informations["cardId"])

	var id int

	err := row.Scan(&id)
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	if id == 0 {
		return false
	}

	return true
}
