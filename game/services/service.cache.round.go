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
	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	strRoomtId := "room:" + strconv.Itoa(int(roomId))

	_ = client.HMGet(ctx, strRoomtId, "meme_text").Val()

	/*
		if meme_text[0] != nil {
			return fmt.Errorf("Room is declared")
		}
	*/

	/*

		@ We, Connected to Redis, taken meme_text and check presence of room thanks to meme_text param in room cache.
		@ Now, We Creating Round Cache.

	*/

	new_meme := queries_meme.GetText(1)

	var card_throwers []string = []string{"0"}

	jsonThrowers, err := json.Marshal(card_throwers)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("JSON marshaling error: %s", err.Error())
	}

	roundKey := "round:" + strRoomtId

	err = client.HMSet(ctx, roundKey, map[string]interface{}{
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

func PlayerThrownACardService(roomid int, data map[string]interface{}) bool {
	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	roundKey := "round:" + strconv.Itoa(int(roomid))
	throwers := client.HGet(ctx, roundKey, "card_throwers").Val()

	var newThrowers []string

	/*

		@ We, Connected to Redis, taken card_throwers
		@ And checking account is played a card in this round.

	*/

	err := json.Unmarshal([]byte(throwers), &newThrowers)
	if err != nil {
		fmt.Println("JSON decode error:  ", err)
		return false
	}

	accountid := data["accountId"]
	for i := 0; i < len(newThrowers); i++ {
		if accountid == newThrowers[i] {
			return true
		}
	}

	return false
}
