package main

import (
	"fmt"
	"net/url"
)

const weatherAPIURL string = "https://api.open-meteo.com/v1/forecast?"

// City
type City struct {
	name     string
	timezone string
	coords   Coord
}

type Coord struct {
	latitude  float32
	longitude float32
}

var cityList = map[string]City{
	"MSK": {"Moscow", "Europe/Moscow", Coord{55.7522, 37.6156}},
	"SPB": {"Saint-Petersburg", "Europe/Moscow", Coord{59.9342, 30.3350}},
	"LPK": {"Lipetsk", "Europe/Moscow", Coord{52.3701, 39.3601}},
	"PTG": {"Pyatigorsk", "Europe/Moscow", Coord{44.0301, 43.0301}},
	"STW": {"Stavropol", "Europe/Moscow", Coord{45.0201, 41.5801}},
	"KSV": {"Kislovodsk", "Europe/Moscow", Coord{43.5401, 42.4301}},
	"IRK": {"Irkutsk", "Asia/Singapore", Coord{52.2978, 104.2964}},
}

func getCityCodes(catalog map[string]City) []string {
	var codes []string
	for code := range catalog {
		codes = append(codes, code)
	}
	return codes
}

// Weather API Call
func weatherAPICall(cityCode string) string {
	apiUrl := weatherAPIURL

	city := cityList[cityCode]

	forecast_days := 1
	forecast_hours := 12

	apiCall := fmt.Sprintf("%slatitude=%f&longitude=%f&hourly=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,precipitation,weather_code,surface_pressure,visibility,wind_speed_10m,temperature_80m,uv_index,is_day,sunshine_duration&wind_speed_unit=ms&timezone=%s&forecast_days=%d&forecast_hours=%d",
		apiUrl, city.coords.latitude, city.coords.longitude, url.PathEscape(city.timezone), forecast_days, forecast_hours)

	return apiCall
}

// Prints data from forecast. Set which next hour forecast in horsFwd
func dumpWeatherDigest(weatherData map[string]interface{}, horsFwd int) string {
	var header string
	switch horsFwd {
	case 0:
		header = "Current weather:\n"
	case 1:
		header = fmt.Sprintf("In %d hour expect:\n\n", horsFwd)
	default:
		header = fmt.Sprintf("In %d hours expect:\n\n", horsFwd)
	}

	s := fmt.Sprintf(`
%s
Temperature: %v°C
Feels like: %v°C
Chance of rain/snow %v%%
Wind speed: %vm/s
Precipitation %vmm
UV index %v`,
		header,
		weatherData["temperature_2m"].([]interface{})[horsFwd],
		weatherData["apparent_temperature"].([]interface{})[horsFwd],
		weatherData["precipitation_probability"].([]interface{})[horsFwd],
		weatherData["wind_speed_10m"].([]interface{})[horsFwd],
		weatherData["precipitation"].([]interface{})[horsFwd],
		weatherData["uv_index"].([]interface{})[horsFwd])

	return s
}
