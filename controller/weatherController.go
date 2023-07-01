package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bostigger/go-weather-tracker/helpers"
	"github.com/bostigger/go-weather-tracker/model"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	OpenWeatherURL = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s"
)

func GetWeatherData(city string, apikey string) (model.WeatherData, error) {

	var w model.WeatherData

	coordinates, err := helpers.GetCoordinates(city, apikey)

	if err != nil {
		return model.WeatherData{}, err
	}
	os.Getenv("apikey")
	res, err := http.Get(fmt.Sprintf(OpenWeatherURL, coordinates.Coord.Latitude, coordinates.Coord.Longitude, apikey))

	if err != nil {
		return model.WeatherData{}, fmt.Errorf("failed to read response from api %s", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.WeatherData{}, fmt.Errorf("failed to read response body  %s", err)
	}
	err = json.Unmarshal(data, &w)
	if err != nil {
		return model.WeatherData{}, fmt.Errorf("failed to unmarshall the response data %s", err)
	}

	defer res.Body.Close()

	mainWeatherData := model.MainWeatherData{
		Temperature:    w.MainWeatherData.Temperature,
		TemperatureMin: w.MainWeatherData.TemperatureMin,
		TemperatureMax: w.MainWeatherData.TemperatureMax,
		Pressure:       w.MainWeatherData.Pressure,
	}
	w.MainWeatherData = &mainWeatherData

	return w, nil

}

func ShowWeatherData(data *model.WeatherData) {
	separator := helpers.SEPARATOR
	fmt.Printf("\n" + separator)
	fmt.Printf("\nCountry: %s\n", data.Country.CountryName)
	fmt.Printf("City Name: %s\n", data.CityName)
	fmt.Printf("Minimum Temperature: %v%s\n", helpers.ConvertToCelsius(data.MainWeatherData.TemperatureMin), "°C")
	fmt.Printf("Maximum Temperature: %v%s\n", helpers.ConvertToCelsius(data.MainWeatherData.TemperatureMax), "°C")
	fmt.Printf("Current Temperature: %v%s\n", helpers.ConvertToCelsius(data.MainWeatherData.Temperature), "°C")
	fmt.Printf("Current Pressure: %v%s\n", data.MainWeatherData.Pressure, "Pa")
	fmt.Println(separator)
}
