package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var Cache *redis.Pool

func init() {
	redisHost := "0.0.0.0"
	redisPort := "6379"

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	const maxConnections = 100

	Cache = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)
}

func Create(set string, value string) error {
	conn := Cache.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", set, string(value))
	if err != nil {
		return err
	}

	return nil

}

func GetAll(set string) ([]interface{}, error) {

	conn := Cache.Get()
	defer conn.Close()
	allValues, err := redis.Values(conn.Do("SMEMBERS", set))

	if err != nil {
		return nil, err
	}
	return allValues, nil
}

func Delete(set string, value string) error {
	conn := Cache.Get()
	defer conn.Close()

	_, err := conn.Do("SREM", set, value)

	if err != nil {
		return err
	}

	return nil
}
