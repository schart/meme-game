package queries

import (
	"fmt"
	cursors "shared-library/utils"
)

// Use before insert for update count
func CountUpdate() (int, error) {
	memeCursor := cursors.TurnMemeCursor()
	res, err := memeCursor.Query(`SELECT count FROM public.memeText ORDER BY id DESC LIMIT 1;`)

	if err != nil {
		fmt.Errorf("Error count update:  %s", err)
		return -1, err
	}

	defer res.Close()

	var count int

	for res.Next() {
		if err := res.Scan(&count); err != nil {
			fmt.Errorf("Error: row scan %s", err)
			continue
		}

		// fmt.Printf("Count: %d\n", count)
	}
	return count + 1, nil
}

//
