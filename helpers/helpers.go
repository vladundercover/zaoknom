package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Returns responce body from an API call. JSON expected.
func RespBody(apiCall string) []byte {
	resp, err := http.Get(apiCall)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return body
}

// Converts JSON from raw responce into GO type
func TypefyResp(data []byte) map[string]interface{} {

	// Chew raw data: init variable with arbitrary data structure
	var digested map[string]interface{}
	if err := json.Unmarshal(data, &digested); err != nil {
		panic(err)
	}

	return digested
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
