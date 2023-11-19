package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	queries_account "shared-library/database/queries/queries-account"
	queries_meme "shared-library/database/queries/queries-meme"
	connection_redis "shared-library/redis"
	utils "shared-library/utils"
	"strconv"

	types_game "shared-library/types/game"
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

	var cardData types_game.CardData

	err := json.Unmarshal([]byte(cards), &cardData)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	/*

		@ Connected to Redis, getting cards and converted to GO LANG objects
		@ Now, we keep other cards, except dropped card

	*/

	var stayedCards map[string]interface{} = map[string]interface{}{}

	for i := 1; i <= 4; i++ {
		var cardValue string

		// Kartı CardData yapısından alın.
		switch i {
		case 1:
			cardValue = cardData.Card1
		case 2:
			cardValue = cardData.Card2
		case 3:
			cardValue = cardData.Card3
		case 4:
			cardValue = cardData.Card4
		}

		if cardValue != "" {
			if cardValue == data["cardId"] {
				fmt.Println("Dropped card: ", cardValue)
			} else {
				stayedCards["card:"+strconv.Itoa(i)] = cardValue
			}

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
func GiveVoteService(data map[string]interface{}, roomid int) error {
	utils.EnvLoader()

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	accountKey := "account:" + data["accountId"].(string)

	refere_status := client.HGet(ctx, accountKey, "referee_status").Val()

	/*
		@ We checked account is referee
	*/

	if refere_status == "1" {

		stringAffected := data["affectedId"].(string)
		
		if stringAffected == data["accountId"].(string) {
			return fmt.Errorf("A referee can not give vote him")
		}

		intAffected, err := strconv.Atoi(stringAffected)
		if err != nil {
			return err
		}

		roomOfAccount := queries_account.GetRoomOfAccount(float64(intAffected))
		if roomOfAccount == nil {
			return fmt.Errorf("Account is not founded")

		}

		accountKey = "account:" + stringAffected

		votes := client.HGet(ctx, accountKey, "votes").Val()
		convertedVotes, _ := strconv.Atoi(votes)

		/*

			@ Note: AffecetId id own the given vote account by referee
			@ We taken votes of affected account and re-save to cache by increment one

		*/

		err = client.HSet(ctx, accountKey, map[string]interface{}{"votes": convertedVotes + 1}).Err()
		if err != nil {
			return fmt.Errorf("Redis HSet error: %s", err.Error())
		}

		roundKey := "round:" + strconv.Itoa(roomid)

		card_throwers := []string{}

		json_throwers, err := json.Marshal(card_throwers)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("JSON marshaling error: %s", err.Error())
		}

		err = client.HSet(ctx, roundKey, "card_throwers", string(json_throwers)).Err()
		if err != nil {
			return fmt.Errorf(err.Error())
		}

	} else {
		return fmt.Errorf("Account is not referee")
	}

	return nil
}
