package controllers

import (
	"encoding/json"
	"fmt"
	service_redis "game/services"
	"net/http"
	queries_account "shared-library/database/queries/queries-account"
	queries_meme "shared-library/database/queries/queries-meme"
	utils "shared-library/utils"
	websocket_connect "shared-library/websocket-connection"
	"strconv"
)

func StartRoundsController(w http.ResponseWriter, r *http.Request) {
	conn := websocket_connect.Connect(w, r)

	// check connection is websocket conn?
	status := utils.CheckWebsocketConnection(r)
	if status == false {
		utils.HandleErrorWS(conn, "This connection is not web socket")
		return
	}

	// Check room is avaiable to ready for the play
	// here

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		message := string(p)
		fmt.Printf("Received message: %s\n", message)

		var data map[string]interface{}
		if err := json.Unmarshal([]byte(message), &data); err != nil {
			fmt.Println("JSON parsing error")
			continue
		}

		// Process json data
		dataType, ok := data["type"].(string)
		if !ok {
			utils.HandleErrorWS(conn, "Undefiend process type: "+dataType)
			return
		}

		intAccountId, _ := strconv.Atoi(data["accountId"].(string))
		roomOfAccount := queries_account.GetRoomOfAccount(float64(intAccountId))

		fmt.Println("room: ", roomOfAccount, roomOfAccount["roomid"])
		roomid := roomOfAccount["roomid"].(int)

		switch dataType {

		/*
			@ Capture the event based on the information received in the websocket connection
			@ Process the activity by prioritizing
		*/

		case "drop-card":
			// Check player thrown a card
			status := service_redis.IsPlayerThrownACardService(roomid, data)
			if status == true {
				utils.HandleErrorWS(conn, "Played a card in this round!")
				return
			}

			/*
				@ If player not played a card then we proceed
			*/

			// Check presence of card/photo
			presence := queries_meme.IsTherePhoto(data)
			if presence == false {
				utils.HandleErrorWS(conn, "This card is not founded")
				return
			}

			/*
				@ If card there proceed
			*/

			// Finally drop a card
			err = service_redis.DropCardService(data, roomid)
			if err != nil {
				utils.HandleErrorWS(conn, err.Error())
				return
			}

			// Turn the dropped card
			utils.HandleSuccessWS(conn, map[string]interface{}{"dropped_card": data["cardId"]})

		case "give-vote":
			/*
				@ Check the priorty of process
			*/
			fmt.Println(dataType)

		default:
			fmt.Println("Undefiend process type:", dataType)
		}
	}

}
