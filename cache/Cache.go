package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"

)

var Cache *redis.Pool

func init(){
	redisHost := "0.0.0.0"
	redisPort := "6378"

	redisAddr := fmt.Sprintf("%s:%s", redisHost , redisPort)
	const maxConnections = 100

	Cache = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)	
}


func Create(key string , value string){
	
}