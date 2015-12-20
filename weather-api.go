package wb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL_OPEN_WEATHER_MAP_API = "http://api.openweathermap.org/data/2.5/weather"
const APP_ID = "460dd10fb0504e5a7b6fb4a0b8daf916"

type Forecast struct {
	Id   float64   `json:"id"`
	Name string    `json:"name"`
	Cod  float64   `json:"cod"`
	Info MainBlock `json:"main"`
}

type MainBlock struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

func Search(city string) (*Forecast, error) {
	resp, err := GetCurrentWeather(city)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var f Forecast
	err = json.Unmarshal(body, &f)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func GetCurrentWeather(city string) (*http.Response, error) {

	url := fmt.Sprintf(URL_OPEN_WEATHER_MAP_API+"?q=%s&appid=%s", city, APP_ID)

	res, err := http.Get(url)

	if err != nil {
		return res, err
	}
	return res, nil
}
