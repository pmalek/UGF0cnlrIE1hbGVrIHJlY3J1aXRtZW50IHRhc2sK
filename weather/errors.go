package weather

type CityNotFoundError string

func (e CityNotFoundError) Error() string {
	return "City '" + string(e) + "' not found"
}
