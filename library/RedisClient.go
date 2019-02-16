package library

import (
	"os"
	"encoding/json"
	"github.com/go-redis/redis"
	"http_server/utils"
)

type RedisConfiguration struct {
    Address      string
    Password       string
}

var redisConfig RedisConfiguration
var RedisClient *redis.Client

func init() {
    loadRedisConfig()
    initRedisClient()
}

func loadRedisConfig() {
	file, _ := os.Open("config/redis.config.json")
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	redisConfig = RedisConfiguration{}
	err := jsonDecoder.Decode(&redisConfig)

	if err == nil {
    	utils.Slog("redis config content decode sucess")
    } else {
    	utils.Slog("redis config content decode failed")
    }
}

func initRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password, // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()

	if err == nil {
		utils.Slog("redis client init sucess")
	} else {
		utils.Slog("redis client init failed")
	}
}