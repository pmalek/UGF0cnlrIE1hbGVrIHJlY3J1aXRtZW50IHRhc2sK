package weather

import (
	"encoding/json"
	"strings"
    "testing"

	"github.com/stretchr/testify/require"
)

func TestCityWeatherUnmarshal(t *testing.T) {
	var sut CityWeather
	require.NoError(t,
		json.NewDecoder(strings.NewReader(cityWeatherJSONResponse)).Decode(&sut))
}

const cityWeatherJSONResponse = `{
    "coord": {
        "lon": -0.13,
        "lat": 51.51
    },
    "weather": [
        {
            "id": 803,
            "main": "Clouds",
            "description": "broken clouds",
            "icon": "04d"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 283.95,
        "feels_like": 276.26,
        "temp_min": 283.15,
        "temp_max": 285.37,
        "pressure": 1017,
        "humidity": 66
    },
    "visibility": 10000,
    "wind": {
        "speed": 9.3,
        "deg": 250,
        "gust": 14.9
    },
    "clouds": {
        "all": 75
    },
    "dt": 1578829240,
    "sys": {
        "type": 1,
        "id": 1414,
        "country": "GB",
        "sunrise": 1578816126,
        "sunset": 1578845677
    },
    "timezone": 0,
    "id": 2643743,
    "name": "London",
    "cod": 200
}`
