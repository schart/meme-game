package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	queries_meme "shared-library/database/queries/queries-meme"
	connection_redis "shared-library/redis"
	utils "shared-library/utils"
	"strconv"
)

func CreateRoundCacheService(roomId float64, room_link string) error {
	/*
		For keep dynamic data in round
	*/
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// Convert to string room id to use as redis key
	strRoomtId := strconv.Itoa(int(roomId))

	// Check room is declared
	_ = client.HMGet(ctx, "room:"+strRoomtId, "meme_text").Val()

	/*if meme_text[0] != nil {
		return fmt.Errorf("Room is declared")
	}*/

	new_meme := queries_meme.TextGetByCount(1)

	var card_throwers []string = []string{"0"}

	jsonThrowers, err := json.Marshal(card_throwers)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("JSON marshaling error: %s", err.Error())
	}

	// Set the JSON string in Redis.
	err = client.HMSet(ctx, "round:"+strRoomtId, map[string]interface{}{
		"round":         0,
		"card_throwers": jsonThrowers,
		"referee_voted": false,
		"meme_text":     new_meme["text:1"],
	}).Err()

	if err != nil {
		return fmt.Errorf("Redis HMSet error: %s", err.Error())
	}

	return nil
}

// Is player thrown a card in this round?
func IsPlayerThrownACardService(roomid int, data map[string]interface{}) bool {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// set round redis key
	roundKey := "round:" + strconv.Itoa(int(roomid))

	throwers := client.HGet(ctx, roundKey, "card_throwers").Val()

	// for keep changed data
	var newThrowers []string

	// convert to json for storage in redis
	err := json.Unmarshal([]byte(throwers), &newThrowers)
	if err != nil {
		fmt.Println("JSON decode error:  ", err)
		return false
	}

	accountid := data["accountId"]
	for i := 0; i < len(newThrowers); i++ {
		// If player thrown a card
		if accountid == newThrowers[i] {
			return true
		}
	}

	return false
}
