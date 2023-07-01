package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/bostigger/go-weather-tracker/model"
	"io/ioutil"
	"net/http"
)

const (
	OpenWeatherURL = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"
)

func GetCoordinates(city string, apiKey string) (model.WeatherCoordinates, error) {
	var coordinates model.WeatherCoordinates

	res, err := http.Get(fmt.Sprintf(OpenWeatherURL, city, apiKey))
	if err != nil {
		return model.WeatherCoordinates{}, fmt.Errorf("error reading from API: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.WeatherCoordinates{}, fmt.Errorf("API returned non-200 status code: %d", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.WeatherCoordinates{}, fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(data, &coordinates)
	if err != nil {
		return model.WeatherCoordinates{}, fmt.Errorf("error unmarshalling from body: %w", err)
	}

	return coordinates, nil
}

const SEPARATOR = "-------------------------------------------------------"
