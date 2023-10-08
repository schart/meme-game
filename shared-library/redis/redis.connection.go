package redis

import (
	"github.com/go-redis/redis/v8"
)

func Connect(host, port string) *redis.Client {
	// Do not forget, exactly set password
	/*client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{host + ":" + port}, DisableCache: true})
	if err != nil {
		panic(err)
	}*/

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // password set
		DB:       0,  // use default DB
	})

	return client
}
