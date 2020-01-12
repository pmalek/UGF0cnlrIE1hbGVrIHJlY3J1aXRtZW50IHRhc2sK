package app

import (
	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/api"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

func Run(config Config) error {
	redisClient, err := createRedisClient(config.RedisAddress)
	if err != nil {
		return errors.Wrapf(err,
			"failed to create redis client for address:%s", config.RedisAddress)
	}

	r := gin.Default()
	r.GET("/weather",
		api.GetWeatherHandler(config.WeatherAPI, redisClient, config.CacheTTL))

	address := "0.0.0.0:" + config.Port
	if err := r.Run(address); err != nil {
		return errors.Wrapf(err, "failed to run the server at %s", address)
	}
	return nil
}

func createRedisClient(redisAddress string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddress,
		// TODO configure password?
		Password: "",
		// NOTE: do we need to configure the db or it's fine to use the default one?
		DB: 0,
	})

	if err := redisClient.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "failed to ping redis")
	}

	return redisClient, nil
}
