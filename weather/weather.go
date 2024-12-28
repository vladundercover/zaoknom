package weather

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

var citiList = map[string]City{
	"MSK": {"Moscow", "Europe/Moscow", Coord{55.7522, 37.6156}},
	"SPB": {"Saint-Petersburg", "Europe/Moscow", Coord{59.9342, 30.3350}},
}

// Weather API Call
func GetData(cityCode string) string {
	apiUrl := weatherAPIURL

	city := citiList[cityCode]

	forecast_days := 1
	forecast_hours := 12

	apiCall := fmt.Sprintf("%slatitude=%f&longitude=%f&hourly=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,precipitation,weather_code,surface_pressure,visibility,wind_speed_10m,temperature_80m,uv_index,is_day,sunshine_duration&wind_speed_unit=ms&timezone=%s&forecast_days=%d&forecast_hours=%d",
		apiUrl, city.coords.latitude, city.coords.longitude, url.PathEscape(city.timezone), forecast_days, forecast_hours)

	return apiCall
}

// Prints data from forecast. Set which next hour forecast in horsFwd
func DumpWeatherDigest(weatherData map[string]interface{}, horsFwd int) {
	var header string
	switch horsFwd {
	case 0:
		header = "Current weather:"
	case 1:
		header = fmt.Sprintf("In %d hour expect:", horsFwd)
	default:
		header = fmt.Sprintf("In %d hours expect:", horsFwd)
	}

	fmt.Printf(`%s
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
}
