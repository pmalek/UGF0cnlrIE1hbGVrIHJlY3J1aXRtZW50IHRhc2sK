package weather

import "encoding/json"

type CityWeather struct {
	Coord    Coord     `json:"coord"`
	Weather  []Weather `json:"weather"`
	Base     string    `json:"base"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Clouds   Clouds    `json:"clouds"`
	Dt       int64     `json:"dt"`
	Sys      Sys       `json:"sys"`
	Timezone int64     `json:"timezone"`
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Cod      int64     `json:"cod"`
}

func (cw *CityWeather) MarshalBinary() ([]byte, error) {
	return json.Marshal(cw)
}

func (cw *CityWeather) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, cw)
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64   `json:"type"`
	ID      int64   `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}
