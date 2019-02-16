package action 

import (
	"fmt"
	"net/http"
	lib "http_server/library"
	"github.com/go-redis/redis"
)

//https://godoc.org/github.com/go-redis/redis#pkg-examples

func RedisGet(writer http.ResponseWriter, request *http.Request) {
	key := "aaa"
	value, err := lib.RedisClient.Get(key).Result()
	if err == redis.Nil {
		fmt.Fprintf(writer, "this is redis key " + key + " value is not exists")
	} else if err != nil {
	    fmt.Fprintf(writer, "this is redis key " + key + " get error")
	}else{
		fmt.Fprintf(writer, "this is redis key " + key + " value is " + value)
	}
}

func RedisSet(writer http.ResponseWriter, request *http.Request) {
	key := "aaa"
	value := "bbb"
	value, err := lib.RedisClient.Set(key, value, 0).Result()
	if err != nil {
	    fmt.Fprintf(writer, "set redis key failed")
		return
	}
	
	fmt.Fprintf(writer, "set redis key sucess")
}

func RedisListLen(writer http.ResponseWriter, request *http.Request) {
	key := "list"
	value, err := lib.RedisClient.HLen(key).Result()
	if err != nil {
	    fmt.Fprintf(writer, "redis hlen failed ")
		return
	}
	
	fmt.Fprintf(writer, "key list len " + string(value))
}