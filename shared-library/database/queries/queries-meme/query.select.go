package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

func PhotoGetByCount(count int) ([]string, error) {
	memeCursor = cursors.MemeCursorTurn()
	rows, err := memeCursor.Query(`SELECT * FROM public.memephoto LIMIT $1;`, count)
	if err != nil {
		fmt.Errorf("Error Select all photo %s", err)
	}
	defer rows.Close()

	var results []string

	for rows.Next() {
		var id string
		var photoid string
		if err := rows.Scan(&id, &photoid); err != nil {
			return nil, fmt.Errorf("Error scanning row: %s", err)
		}
		results = append(results, photoid)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %s", err)
	}

	return results, nil
}

func TextGetByCount(count int) ([]string, error) {
	memeCursor = cursors.MemeCursorTurn()
	rows, err := memeCursor.Query(`SELECT * FROM public.memetext LIMIT $1;`, count)

	if err != nil {
		fmt.Errorf("Error Select all text %s", err)
	}
	defer rows.Close()

	var results []string

	for rows.Next() {
		var id string
		var text string
		if err := rows.Scan(&id, &text); err != nil {
			return nil, fmt.Errorf("Error scanning row: %s", err)
		}
		results = append(results, text)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %s", err)
	}

	return results, nil

}
