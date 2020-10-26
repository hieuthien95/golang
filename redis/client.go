package cluster

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

type redisCluterClient struct {
	c *redis.ClusterClient
}

var client = &redisCluterClient{}

func initCl(hostname string) *redisCluterClient {

	addr := strings.Split(hostname, ",")

	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addr,
	})

	if err := c.Ping().Err(); err != nil {
		return nil
	}

	client.c = c
	return client
}

func (client *redisCluterClient) setKey(key string, value interface{}, expiration time.Duration) error {

	_, err := json.Marshal(value)

	if err != nil {
		return err
	}

	if err = client.c.Set(key, value, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (client *redisCluterClient) getKey(key string, src interface{}) error {

	val, err := client.c.Get(key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}
