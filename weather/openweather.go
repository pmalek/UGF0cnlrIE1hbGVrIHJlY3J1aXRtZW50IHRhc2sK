package weather

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/friendsofgo/errors"
	log "github.com/sirupsen/logrus"
)

const (
	OpenWeatherAPIBaseURL = "https://api.openweathermap.org/data/2.5/"
)

type OpenWeatherAPI struct {
	APIKey     string
	httpClient *http.Client
}

func NewOpenWeatherAPI(APIkey string, httpClient *http.Client) OpenWeatherAPI {
	return OpenWeatherAPI{
		APIKey:     APIkey,
		httpClient: httpClient,
	}
}

func (o OpenWeatherAPI) ForCities(cities []string) (ForCitiesResponse, error) {
	var (
		baseURL = OpenWeatherAPIBaseURL + `weather?appid=` + o.APIKey
		ret     = make(ForCitiesResponse, len(cities))
	)

	// NOTE: This could most likely be improved by storing city to city ID mapping
	// in a key, value store and then using https://openweathermap.org/current#severalid
	// to get those in bulk (note that this is only available by ID not by city name)
	// but that would make the application significantly more complex for this
	// task (should it though?).
	//
	// For now just make concurrent requests for each city.

	type Entry struct {
		City    string
		Weather *CityWeather
		Err     error
	}
	var (
		err error
		ch  = make(chan Entry)
		wg  = sync.WaitGroup{}
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for r := range ch {
			if r.Err != nil {
				err = r.Err
				return
			}
			ret[r.City] = r.Weather
		}
	}()

	wgRequests := sync.WaitGroup{}
	wgRequests.Add(len(cities))
	for _, c := range cities {
		go func(city string) {
			defer wgRequests.Done()

			url := baseURL + `&q=` + city
			log.Debugf("Calling OpenWeather API at %v", url)

			resp, err := o.httpClient.Get(url)
			if err != nil {
				ch <- Entry{
					Err: errors.Wrapf(err,
						"problem reaching OpenWeather's API for city: %s", city,
					),
				}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				ch <- Entry{
					Err: CityNotFoundError(city),
				}
				return
			} else if resp.StatusCode != http.StatusOK {
				ch <- Entry{
					Err: errors.Errorf(
						"OpenWeather API returned status code: %s for city: %s",
						resp.Status, city,
					),
				}
				return
			}

			var cityWeather CityWeather
			if err := json.NewDecoder(resp.Body).Decode(&cityWeather); err != nil {
				ch <- Entry{
					Err: errors.Wrapf(err,
						"failed to decode OpenWeather API response status code: %d for city: %s",
						resp.StatusCode, city,
					),
				}
				return
			}

			ch <- Entry{
				City:    city,
				Weather: &cityWeather,
			}
		}(c)
	}
	wgRequests.Wait()
	close(ch)

	wg.Wait()
	return ret, err
}
