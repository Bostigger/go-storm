package model

type WeatherData struct {
	CityName        string           `json:"name"`
	Country         CountryDetails   `json:"sys"`
	MainWeatherData *MainWeatherData `json:"main"`
}

type MainWeatherData struct {
	Temperature    float64 `json:"temp"`
	TemperatureMin float64 `json:"temp_min"`
	TemperatureMax float64 `json:"temp_max"`
	Pressure       float64 `json:"pressure"`
}

type WeatherCoordinates struct {
	Coord Coordinates `json:"coord"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type CountryDetails struct {
	CountryName string `json:"country"`
}
