package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	connection_redis "shared-library/redis"
	utils "shared-library/utils"
	"strconv"
)

func CacheInitForGame(accountId float64) error {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// Convert to string account id to use as redis key
	strAccountId := strconv.Itoa(int(accountId))

	// Check votes variable is declared
	redisSlices := client.HMGet(ctx, "strAccountId", "votes")  
	resultInterface, err := redisSlices.Result()
	if err != nil {
		return fmt.Errorf("Redis slices error: %s", err.Error())  
	}

	fmt.Println("redis slices result: ", resultInterface[0] == nil, resultInterface[0])
	if resultInterface[0] == nil {
		return fmt.Errorf("This account is playing now")  
	}

	cards := []string{}

	// Convert the []string slice to a JSON string.
	jsonCards, err := json.Marshal(cards)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("JSON marshaling error: %s", err.Error())  
	}

	// Set the JSON string in Redis.
	err = client.HMSet(ctx, strAccountId, map[string]interface{}{
		"votes": 0,
		"cards": jsonCards,
	}).Err()

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Redis HMSet error: %s", err.Error())  
	}
	return nil
}

/*
// Update vote/score, if you take vote from other players this win value as +1
func UpdateVote(room, user string) bool {
	utils.EnvLoader()
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	return true
}

// Update the cards, if you play the any card then card counter lose value as -1
func UpdateCards(name string) bool {
	utils.EnvLoader()
	// client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	// ctx := context.Background()

	return true
}

// Update room status, if you leave which 'll  be 0
// if you join  to the room then 1
func UpdateRoom(name string) bool {
	utils.EnvLoader()
	// client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	// ctx := context.Background()

	return true
}
*/
