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

func CreateAccountCacheService(accountId float64) error {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// Convert to string account id to use as redis key
	strAccountId := strconv.Itoa(int(accountId))

	// Check player is declared
	_ = client.HGet(ctx, "account:"+strAccountId, "cards").Val()

	/*if cardSlices != "" {
		return fmt.Errorf("This account is playing now")
	}*/

	cards := queries_meme.PhotoGetByCount(5)

	// Convert the []string slice to a JSON string.
	jsonCards, err := json.Marshal(cards)
	if err != nil {
		return fmt.Errorf("JSON marshaling error: %s", err.Error())
	}

	// Set the JSON string in Redis.
	err = client.HSet(ctx, "account:"+strAccountId, map[string]interface{}{
		"votes":          0,
		"cards":          jsonCards,
		"referee_status": true,
	}).Err()

	if err != nil {
		return fmt.Errorf("Redis HMSet error: %s", err.Error())
	}

	return nil
}

func DropCardService(data map[string]interface{}, roomid int) error {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// Set account key
	accountKey := "account:" + data["accountId"].(string)

	// Get the cards of player
	cards := client.HGet(ctx, accountKey, "cards").Val()

	// for keep changed data
	var newCards []string
	if cards == "" {
		return fmt.Errorf("Player have not a card")
	}

	//  converting to array via json
	err := json.Unmarshal([]byte(cards), &newCards)
	if err != nil {
		return fmt.Errorf("JSON decode error:  ", err)
	}

	// Keep stayed cards
	var stayedCards []string

	// Delete card in card
	for i := 0; i < len(newCards); i++ {
		if newCards[i] == data["cardId"] {
			fmt.Println("Dropped card: ", newCards[i])
		} else {
			fmt.Println(newCards[i])
			stayedCards = append(stayedCards, newCards[i])
		}
	}

	// Convert to json stayed cards array
	json_stayed, err := json.Marshal(stayedCards)
	if err != nil {
		return fmt.Errorf("Json marshall error: ", err)
	}

	// Set new cards
	err = client.HSet(ctx, accountKey, map[string]interface{}{
		"cards": json_stayed,
	}).Err()
	if err != nil {
		return fmt.Errorf("Redis HSet error: %s", err.Error())
	}

	/*

		@ Let's add to the lineup of players who played cards in the round

	*/

	// Set redis access key
	roundKey := "round:" + strconv.Itoa(roomid)

	// Get the accounts the played a card
	card_throwers := client.HGet(ctx, roundKey, "card_throwers").Val()

	// Decode the stored json data
	new_card_throwers := []string{}
	err = json.Unmarshal([]byte(card_throwers), &new_card_throwers)
	if err != nil {
		return fmt.Errorf("Json Umarshal error: ", err)
	}

	// Add to throwers
	new_card_throwers = append(new_card_throwers, data["accountId"].(string))

	// Convert to json for storage in redis
	json_throwers, err := json.Marshal(new_card_throwers)
	if err != nil {
		return fmt.Errorf("Json marshall error: ", err)
	}

	// Add new card throwers
	err = client.HSet(ctx, roundKey, map[string]interface{}{
		"card_throwers": string(json_throwers),
	}).Err()

	return nil
}

// Update vote/score, if you take vote from other players this win value as +1
func GiveVote(data map[string]interface{}) error {
	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	accountId := data["accountId"]
	refere_status := client.HGet(ctx, fmt.Sprintf("account:%s", accountId), "refere_status").Val()

	if refere_status == "true" {
		votes := client.HGet(ctx, fmt.Sprintf("account:%s", data["affectedId"]), "votes").Val()
		convertedVotes, _ := strconv.Atoi(votes)

		err := client.HSet(ctx, fmt.Sprintf("account:%s", data["affectedId"]), map[string]interface{}{
			"votes": convertedVotes + 1}).Err()
		if err != nil {
			return fmt.Errorf("Redis HSet error: %s", err.Error())
		}

	} else {
		return fmt.Errorf("Account is not referee")
	}

	return nil
}
