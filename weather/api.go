package weather

type API interface {
	// ForCities returns weather information for the list of cities and an error.
	ForCities(cities []string) (ForCitiesResponse, error)
}

type ForCitiesResponse map[string]*CityWeather
