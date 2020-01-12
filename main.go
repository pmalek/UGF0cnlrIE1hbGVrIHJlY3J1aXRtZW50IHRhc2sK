package main

import (
	"net/http"
	"os"
	"time"

	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/app"
	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/flags"
	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/weather"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "app",
		Flags: []cli.Flag{
			flags.FlagOpenweatherAPIKey,
			flags.FlagPort,
			flags.FlagRedisAddress,
			flags.FlagRedisTTL,
		},
		Action: func(c *cli.Context) error {
			weatherAPI := weather.NewOpenWeatherAPI(flags.OpenWeatherAPIKey, httpClient())

			return app.Run(app.Config{
				Port:         flags.Port,
				RedisAddress: flags.RedisAddress,
				WeatherAPI:   weatherAPI,
				CacheTTL:     flags.RedisTTL,
			})
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Failed to run the app")
	}
}

func httpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
