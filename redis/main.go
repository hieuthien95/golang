package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

// Object is ok
type Object struct {
	Str string
	Num int
}

type valueEx struct {
	Name  string
	Email string
}

func main() {

	// ----------- Sample Ping - Pong START------------
	// New connect
	client := newConn()

	// Test ping connect
	ping(client)
	// ----------- Sample Ping - Pong END------------

	// ----------- Sample Database: Key - Value START------------
	// Set value
	set(client)

	// Get value
	get(client)
	// ----------- Sample Database: Key - Value END ------------

	// ----------- Sample Cache START------------
	// Cache
	sampleCache()
	// ----------- Sample Cache END------------

	// ----------- Sample Redis Cluster START------------
	// Redis Cluster
	redisCluster()
	// ----------- Sample Redis Cluster END------------

}

func redisCluster() {

	addr := []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"}
	redisCluterClient, err := initCl(addr)
	if err != nil {
		fmt.Println("Error connect: ", err)
		return
	}
	key1 := "myKeyCluster"
	value1 := &valueEx{Name: "congpv", Email: "congpv@lozi.vn"}
	err = redisCluterClient.setKey(key1, value1, time.Minute*1)
	if err != nil {
		fmt.Println("Error: %v", err.Error())
	}
	value2 := &valueEx{}
	err = redisCluterClient.getKey(key1, value2)
	if err != nil {
		fmt.Println("Error: %v", err.Error())
	}

	fmt.Println(value2.Email)
	fmt.Println(value2.Name)
}

func newConn() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	return rdb
}

func ping(client *redis.Client) error {

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func set(client *redis.Client) error {

	// Set param 0: key has no expiration time.
	if err := client.Set("name", "CongPV", 0).Err(); err != nil {
		return err
	}

	if err := client.Set("fullname", "Phan Van Cong 2", 0).Err(); err != nil {
		return err
	}

	return nil
}

func get(client *redis.Client) error {

	value, err := client.Get("name").Result()
	if err != nil {
		return err
	}

	fmt.Println("name is value: ", value)

	return nil
}

func sampleCache() {

	// Set cache
	setCache()

	// Get cache
	getCache()

}

// set cache value
func setCache() {

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Minute,
	})
}

// get cache value
func getCache() {

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	var wanted Object
	if err := codec.Get("mykey", &wanted); err == nil {
		fmt.Println(wanted)
	} else {
		fmt.Println(err)
	}
}

type redisCluterClient struct {
	c *redis.ClusterClient
}

var client = &redisCluterClient{}

func initCl(addr []string) (*redisCluterClient, error) {

	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addr,
	})

	if err := c.Ping().Err(); err != nil {
		return nil, err
	}

	client.c = c
	return client, nil
}

func (client *redisCluterClient) setKey(key string, value interface{}, expiration time.Duration) error {

	json, err := json.Marshal(value)

	if err != nil {
		return err
	}

	if err = client.c.Set(key, json, expiration).Err(); err != nil {
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
