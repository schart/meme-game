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

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	accountKey := "account:" + strconv.Itoa(int(accountId))

	_ = client.HGet(ctx, accountKey, "cards").Val()

	/*
		if cardSlices != "" {
			return fmt.Errorf("This account is playing now")
		}
	*/

	/*

		@ Connected redis and check account cache declared or un declared by getting card.
		@ Now, we finally create cache of account.

	*/

	cards := queries_meme.GetPhoto(5)
	jsonCards, err := json.Marshal(cards)
	if err != nil {
		return fmt.Errorf("JSON marshaling error: %s", err.Error())
	}

	err = client.HSet(ctx, accountKey, map[string]interface{}{
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

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	accountKey := "account:" + data["accountId"].(string)

	cards := client.HGet(ctx, accountKey, "cards").Val()

	if cards == "" {
		return fmt.Errorf("Player have not a card")
	}

	var newCards map[string]string

	err := json.Unmarshal([]byte(cards), &newCards)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	/*

		@ Connected to Redis, getting cards and converted to GO LANG objects
		@ Now, we keep other cards, except dropped card

	*/

	var stayedCards []string

	for i := 0; i < len(newCards); i++ {
		cardkey := "card:" + strconv.Itoa(i)

		if newCards[cardkey] == data["cardId"] {
			fmt.Println("Dropped card: ", newCards[cardkey])
		} else {
			fmt.Println(newCards[cardkey])
			stayedCards = append(stayedCards, newCards[cardkey])
		}
	}

	json_stayed, err := json.Marshal(stayedCards)
	if err != nil {
		return fmt.Errorf("Json marshall error: ", err)
	}

	err = client.HSet(ctx, accountKey, map[string]interface{}{
		"cards": json_stayed,
	}).Err()

	if err != nil {
		return fmt.Errorf("Redis HSet error: %s", err.Error())
	}

	/*
		@ We kept the cards except for the dropped one and updated the cache
		@ Let's add to the lineup of players who played cards in the round

	*/

	roundKey := "round:" + strconv.Itoa(roomid)

	card_throwers := client.HGet(ctx, roundKey, "card_throwers").Val()

	new_card_throwers := []string{}
	err = json.Unmarshal([]byte(card_throwers), &new_card_throwers)
	if err != nil {
		return fmt.Errorf("Json Umarshal error: ", err)
	}

	new_card_throwers = append(new_card_throwers, data["accountId"].(string))

	json_throwers, err := json.Marshal(new_card_throwers)
	if err != nil {
		return fmt.Errorf("Json marshall error: ", err)
	}

	err = client.HSet(ctx, roundKey, map[string]interface{}{
		"card_throwers": string(json_throwers),
	}).Err()

	return nil
}

// Update vote/score, if you take vote from other players this win value as +1
func GiveVoteService(data map[string]interface{}) error {
	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	accountKey := "account:" + strconv.Itoa(data["accountId"].(int))

	refere_status := client.HGet(ctx, accountKey, "refere_status").Val()

	/*
		@ We checked account is referee
	*/

	if refere_status == "true" {
		accountKey = "account:" + strconv.Itoa(data["affectedId"].(int))

		votes := client.HGet(ctx, accountKey, "votes").Val()
		convertedVotes, _ := strconv.Atoi(votes)

		/*

			@ Note: AffecetId id own the given vote account by referee
			@ We taken votes of affected account and re-save to cache by increment one

		*/

		err := client.HSet(ctx, accountKey, map[string]interface{}{"votes": convertedVotes + 1}).Err()
		if err != nil {
			return fmt.Errorf("Redis HSet error: %s", err.Error())
		}

	} else {
		return fmt.Errorf("Account is not referee")
	}

	return nil
}
