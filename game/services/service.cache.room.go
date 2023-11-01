package services

import (
	"context"
	"fmt"
	"os"
	connection_redis "shared-library/redis"
	utils "shared-library/utils"
	"strconv"
)

func CreateRoomCacheService(roomId float64, room_link string, playerCount, ownerId int) error {
	utils.EnvLoader()

	// Create redis connections
	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	// Convert to string room id to use as redis key
	strRoomtId := strconv.Itoa(int(roomId))

	// Check room is declared
	_ = client.HMGet(ctx, "room:"+strRoomtId, "room_link").Val()

	/*if link[0] != "" {
		return fmt.Errorf("Room is declared")
	}*/

	// Set the JSON string in Redis.
	err := client.HMSet(ctx, "room:"+strRoomtId, map[string]interface{}{
		"room_link":  room_link,
		"ownerId":    ownerId,
		"count_user": playerCount,
	}).Err()

	if err != nil {
		return fmt.Errorf("Redis HMSet error: %s", err.Error())
	}

	return nil
}
