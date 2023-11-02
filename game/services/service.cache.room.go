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

	client := connection_redis.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ctx := context.Background()

	strRoomtId := strconv.Itoa(int(roomId))

	_ = client.HMGet(ctx, "room:"+strRoomtId, "room_link").Val()

	/*
		if link[0] != "" {
			return fmt.Errorf("Room is declared")
		}
	*/

	/*
		@ We connected to Redis and taken room_link in the room cache with key
		@ Also checked room link is there for check presence of room in room cache
		@ Finally, we create a room in room cache

	*/

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
