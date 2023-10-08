package services

/*
import (
	"context"
	"fmt"
	"os"
	connection_redis "shared-library/redis"
	utils "shared-library/utils"
)

func CreateRoom(data map[string]interface{}) bool {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	// generate id for name of room
	// id := utils.Generator()

	ctx := context.Background()


		err := client.Do(ctx, client.B().Set().Key("x" data["room"]+":"+id.String()).
			Value(fmt.Sprintf(`{"%v:%v": { "users" : {"%v": {"votes": "0", "cards": "0"} } }}`, data["room"], id.String(), data["user"])).Build()).Error()

		if err != nil {
			fmt.Println("Error: ", err)
			return false
		}


	err := client.HSetNX(ctx, "room", "users", map[string]interface{}{
		"vote": "2",
	}).Err()
	if err != nil {
		// Handle error
	}

	fmt.Println(client.HGet(ctx, "room", "users"))

	return true
}


// Update vote/score, if you take vote from other players this win value as +1
func UpdateVote(room, user string) bool {
	utils.EnvLoader()
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	votes, _ := client.Do(ctx, client.B().Get().Key("x").Build()).ToString()
	parsed_votes := utils.JsonToMap(votes)

	defer fmt.Print("result votes: ", parsed_votes)
	for x, y := range parsed_votes {
		fmt.Println("x, y: ", x, y)
	}

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
