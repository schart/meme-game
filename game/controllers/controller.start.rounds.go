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

	status := utils.CheckWebsocketConnection(r)
	if status == false {
		utils.HandleErrorWS(conn, "This connection is not web socket")
		return
	}

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

		dataType, ok := data["type"].(string)
		if !ok {
			utils.HandleErrorWS(conn, "Undefiend process type: "+dataType)
			return
		}

		intAccountId, _ := strconv.Atoi(data["accountId"].(string))
		roomOfAccount := queries_account.GetRoomOfAccount(float64(intAccountId))
		if roomOfAccount == nil {
			utils.HandleErrorWS(conn, "undefiend account")
			return
		}

		roomid := roomOfAccount["roomid"].(int)

		/*

			@ We was read stream of websocket connection and taken data by converted to GO lang objects
			@ And was get room id thanks to account id of taken the websocket steram

		*/

		switch dataType {

		/*
			@ Capture the event based on the information received in the websocket connection
			@ Process the activity by prioritizing
		*/

		case "drop-card":
			status := service_redis.PlayerThrownACardService(roomid, data)
			if status == true {
				utils.HandleErrorWS(conn, "Played a card in this round!")
				return
			}

			/*
				@ If player not played a card then we proceed
			*/

			presence := queries_meme.PhotoAvailable(data)
			if presence == false {
				utils.HandleErrorWS(conn, "This card is not founded")
				return
			}

			/*
				@ If card there proceed
			*/

			err = service_redis.DropCardService(data, roomid)
			if err != nil {
				utils.HandleErrorWS(conn, err.Error())
				return
			}

			utils.HandleSuccessWS(conn, map[string]interface{}{"dropped_card": data["cardId"]})

		case "give-vote":
			/*

				@ Check the priorty of process

			*/

			thrownCardPlayer := service_redis.ThrownCardService(roomid)
			if len(thrownCardPlayer) < 1 {
				utils.HandleErrorWS(conn, "Each player must throw a card in round!")
				return
			}

			err := service_redis.GiveVoteService(data, roomid)
			if err != nil {
				utils.HandleErrorWS(conn, err.Error())
				return
			}

			/*

				@ We checked that each player played a card.
				@ And we checked player is referee or  referee try to give vote self..
				@ Finally, we delete thrown card and increment to round counter

				@ Notice: if players in last round, we turn winner after of voting

			*/

			err = service_redis.DeleteThrownCardService(roomid)
			if err != nil {
				utils.HandleErrorWS(conn, err.Error())
				return
			}

			roundCounting := service_redis.GetRoundService(roomid)
			if roundCounting == 5 {
				winner := service_redis.FindWinnerService(roomid)
				fmt.Println("winner: ", winner)

				utils.HandleSuccessWS(conn, map[string]interface{}{"winner": winner})
				return
			}

			err = service_redis.IncrementRoundService(roomid)
			if err != nil {
				utils.HandleErrorWS(conn, err.Error())
				return
			}

			utils.HandleSuccessWS(conn, map[string]interface{}{"winner_of_round": data["affectedId"]})

		default:
			fmt.Println("Undefiend process type:", dataType)
			utils.HandleErrorWS(conn, "Undefiend process type:"+dataType)
			return
		}
	}

}
