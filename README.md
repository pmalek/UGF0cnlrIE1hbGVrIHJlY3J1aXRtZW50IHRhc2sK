# gogoapp / UGF0cnlrIE1hbGVrIHJlY3J1aXRtZW50IHRhc2sK

## How to run

To build and start the app `make compose-up-build` should be enough.
Environment variables can be tweaked in the `.env` file.

    PORT=8080
    REDIS_ADDRESS=redis:6379
    REDIS_TTL=10m
    OPENWEATHER_API_KEY=<PLACEHOLDER>

### Exemplar request

    curl "localhost:8080/weather?cities=Warsaw,Gdansk"

### Exemplar response:

    {
        "Gdansk": {
            "coord": {
                "lon": 18.65,
                "lat": 54.35
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
                "temp": 277.17,
                "feels_like": 271.68,
                "temp_min": 276.15,
                "temp_max": 278.15,
                "pressure": 1014,
                "humidity": 93
            },
            "wind": {
                "speed": 5.7,
                "deg": 220
            },
            "clouds": {
                "all": 75
            },
            "dt": 1578837255,
            "sys": {
                "type": 1,
                "id": 1696,
                "message": 0,
                "country": "PL",
                "sunrise": 1578812461,
                "sunset": 1578840326
            },
            "timezone": 3600,
            "id": 3099434,
            "name": "Gdansk",
            "cod": 200
        },
        "Warsaw": {
            "coord": {
                "lon": 21.01,
                "lat": 52.23
            },
            "weather": [
                {
                    "id": 804,
                    "main": "Clouds",
                    "description": "overcast clouds",
                    "icon": "04d"
                }
            ],
            "base": "stations",
            "main": {
                "temp": 278.3,
                "feels_like": 273.43,
                "temp_min": 277.04,
                "temp_max": 279.26,
                "pressure": 1021,
                "humidity": 69
            },
            "wind": {
                "speed": 4.1,
                "deg": 220
            },
            "clouds": {
                "all": 100
            },
            "dt": 1578837080,
            "sys": {
                "type": 1,
                "id": 1713,
                "message": 0,
                "country": "PL",
                "sunrise": 1578811255,
                "sunset": 1578840399
            },
            "timezone": 3600,
            "id": 756135,
            "name": "Warsaw",
            "cod": 200
        }
    }

## TODOs and notes

* Redis communication/requests maybe could be improved (pipelining for Get operation?)
* Storage in Redis can be improved, for now values are stored as JSON). Maybe
    more efficient approach can be chosen, e.g. storing as protobufs?
* For the API JSON REST was chosen as it wasn't specified wether to use Protobuf
    or similar, based API.
* Only a very minimal set of tests (**a test**) has been provided (no mention of that
    in task description and mocking/stubbing OpenWeather's API would increase the
    effort and time to finish the task overall)
* Geting weather data from OpenWeather API is done in a single request for single
    city manner. This could be improved by storing a city to city ID mapping in
    some sort of key, value store but that would make the application more complex
    (should it though?)
