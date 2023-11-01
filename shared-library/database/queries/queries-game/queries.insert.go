package queriesgame

import (
	cursors "shared-library/utils"

	"github.com/gofrs/uuid"
)

func RoomInsert(accountId float64, link uuid.UUID) error {
	gameCursor := cursors.GameCursorTurn()

	tx, err := gameCursor.Begin()
	if err != nil {
		return err
	}

	rows := tx.QueryRow(`INSERT INTO public.rooms(accountid, link) VALUES($1, $2) returning id`, accountId, link)

	var roomid int
	err = rows.Scan(&roomid)
	if err != nil {
		return err
	}

	err = RoomJoin(accountId, roomid, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func RoomJoin(accountId float64, roomId int, is_owner bool) error {
	accountCursor := cursors.AccountCursorTurn()

	tx, err := accountCursor.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query("INSERT INTO public.account_rooms(accountid, roomid, is_owner) VALUES($1, $2, $3)", accountId, roomId, is_owner)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
