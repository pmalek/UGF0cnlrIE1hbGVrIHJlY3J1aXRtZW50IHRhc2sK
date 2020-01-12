package app

import (
	"time"

	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/weather"
)

type Config struct {
	RedisAddress string
	Port         string
	WeatherAPI   weather.API
	CacheTTL     time.Duration
}
