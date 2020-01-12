package flags

import "github.com/urfave/cli/v2"

import "time"

var Port string
var FlagPort = &cli.StringFlag{
	Name:        "port",
	Value:       "8080",
	Usage:       "Run the app on the specified `PORT`",
	EnvVars:     []string{"PORT"},
	Destination: &Port,
}

var OpenWeatherAPIKey string
var FlagOpenweatherAPIKey = &cli.StringFlag{
	Name:        "openweather-api-key",
	Usage:       "API key for accessing OpenWeather's API",
	EnvVars:     []string{"OPENWEATHER_API_KEY"},
	Required:    true,
	Destination: &OpenWeatherAPIKey,
}

var RedisAddress string
var FlagRedisAddress = &cli.StringFlag{
	Name:        "redis-address",
	Usage:       "Address where redis can be reached at",
	Required:    true,
	EnvVars:     []string{"REDIS_ADDRESS"},
	Destination: &RedisAddress,
}

var RedisTTL time.Duration
var FlagRedisTTL = &cli.DurationFlag{
	Name:        "redis-ttl",
	Usage:       "TTL for values saved in redis",
	Value:       10 * time.Minute,
	EnvVars:     []string{"REDIS_TTL"},
	Destination: &RedisTTL,
}
