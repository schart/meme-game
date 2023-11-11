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

	roomKey := "room:" + strconv.Itoa(int(roomId))

	_ = client.HMGet(ctx, roomKey, "meme_text").Val()

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

	var card_throwers []string = []string{}

	jsonThrowers, err := json.Marshal(card_throwers)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("JSON marshaling error: %s", err.Error())
	}

	roundKey := "round:" + strconv.Itoa(int(roomId))

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
	fmt.Println(roundKey)
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
	fmt.Println("new: ", newThrowers)
	accountid := data["accountId"]
	for i := 0; i < len(newThrowers); i++ {
		if accountid == newThrowers[i] {
			fmt.Println("test: ", accountid, newThrowers[i])
			return true
		}
	}

	return false
}

func ThrownCardService(roomid int) []string {
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
		return nil
	}

	return newThrowers
}

func DeleteThrownCardService(roomid int) error {

	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	roundKey := "round:" + strconv.Itoa(int(roomid))
	err := client.HDel(ctx, roundKey, "card_throwers").Err()
	if err != nil {
		fmt.Println("[DeleteThrownCardService] ERROR: ", err)
		return err
	}

	return nil
}

func IncrementRoundService(roomid int) error {

	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	roundKey := "round:" + strconv.Itoa(roomid)

	round := client.HGet(ctx, roundKey, "round").Val()
	convertedRound, _ := strconv.Atoi(round)

	err := client.HSet(ctx, roundKey, map[string]interface{}{"round": convertedRound + 1}).Err()
	if err != nil {
		return fmt.Errorf("Redis HSet error: %s", err.Error())
	}

	return nil
}
