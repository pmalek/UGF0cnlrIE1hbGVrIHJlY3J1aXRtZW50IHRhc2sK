package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/pmalek/UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK/weather"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

func GetWeatherHandler(weatherAPI weather.API, redisClient *redis.Client, cacheTTL time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		citiesRaw, ok := c.GetQuery("cities")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no city specified",
			})
			return
		}

		// NOTE: split by comma, no spec in task definition
		cities := strings.Split(citiesRaw, ",")
		if len(cities) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no city specified",
			})
			return
		}

		// NOTE: Maybe think about initial city name validation (names stored in cache?)

		ret := make(weather.ForCitiesResponse, len(cities))

		// TODO:
		// This (commented) version could allow to use MarshalBinary and
		// UnmarshalBinary defined on weather.CityWeather and to simplify the
		// code but on the expense of making multiple redis requests because
		// MGet()'s response (*SliceCmd) doesn't support struct scanning.
		// For now just use the more complex but slightly faster version using
		// MGet() below
		// We could probably use a pipeline but that would also be a bit
		// complex (arguably).

		// for _, city := range cities {
		// 	var cw weather.CityWeather
		// 	if err := redisClient.Get(city).Scan(&cw); err != nil {
		// 		if errors.Is(err, redis.Nil) {
		// 			// No data for this city in Redis.
		// 			continue
		// 		}

		// 		// If we encounter an error trying to get data from redis then
		// 		// there's probably a problem with network or redis server.
		// 		// Log it and continue with using OpenWeather's API.
		// 		log.WithError(err).Errorf("Problem getting data from redis for cities: %v", cities)
		// 		continue
		// 	}

		// 	ret[city] = &cw
		// }

		// Try to get data for requested cities from cache.
		cachedInts, err := redisClient.MGet(cities...).Result()
		if err != nil {
			// If we encounter an error connecting to redis and it wasn't a redis.Nil
			// error then there's probably a problem with network or redis server.
			// Log it and continue with using OpenWeather's API.
			log.WithError(err).Errorf("Problem getting data from redis for cities: %v", cities)
		} else if err == nil {
			for i, cachedInt := range cachedInts {
				if cachedInt == nil {
					// NOTE: Maybe add metrics for cache misses/hits?
					log.Infof("Cache miss for city %s", cities[i])
					continue
				}
				log.Infof("Cache hit for city %s", cities[i])
				cachedStr, ok := cachedInt.(string)
				if !ok {
					log.Errorf("Invalid data in cache for city %s, type: %T, data: %v",
						cities[i], cachedInt, cachedInt)
					continue
				}

				var cw weather.CityWeather
				err := json.NewDecoder(strings.NewReader(cachedStr)).Decode(&cw)
				if err != nil {
					log.WithError(err).
						Errorf("Couldn't decode weather data from cache for %s", cities[i])
					continue
				}
				ret[cities[i]] = &cw
			}

			if len(ret) == len(cities) {
				// Got all data from cache.
				c.JSON(http.StatusOK, ret)
				return
			}
		}

		// NOTE:
		// Don't bother reusing incomplete cache data. We'll fetch new data below
		// anyway so we might use it (and refresh cache with it).

		// Otherwise call OpenWeather API and try to cache its return value.
		resp, err := weatherAPI.ForCities(cities)
		if err != nil {
			log.WithError(err).
				Errorf("Problem with weather's API for cities: %v", cities)

			var cnf weather.CityNotFoundError
			if errors.As(err, &cnf) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "problem calling external weather API",
			})
			return
		}

		if err := saveToCache(redisClient, cacheTTL, resp); err != nil {
			log.WithError(err).WithField("cities", cities).
				Errorf("Hit weather API, failed saving data to cache")
		} else {
			log.WithField("cities", cities).
				Infof("Hit weather API and successfully saved data in cache for reuse")
		}

		c.JSON(http.StatusOK, resp)
	}
}

func saveToCache(redisClient *redis.Client, cacheTTL time.Duration, data weather.ForCitiesResponse) error {
	p := redisClient.Pipeline()
	defer p.Close()

	for city, weather := range data {
		if err := p.Set(city, weather, cacheTTL).Err(); err != nil {
			return errors.Wrapf(err, "problem saving data to cache for '%s'", city)
		}
	}
	_, err := p.Exec()
	return errors.Wrapf(err, "failed saving data to cache")
}
