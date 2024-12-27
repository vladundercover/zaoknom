package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"zaoknom/weather"
)

func main() {
	weatherAPICall := weather.CheckFor("MSK")
	resp, _ := http.Get(weatherAPICall)
	body, _ := io.ReadAll(resp.Body)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	hourly := dat["hourly"].(map[string]interface{})

	fmt.Printf(`Current temp: %v°C
	Feels like: %v°C
	Chance of rain/snow %v%%
	Wind speed: %vm/s
	Precipitation %vmm
	UV index %v`,
		hourly["temperature_2m"].([]interface{})[1],
		hourly["apparent_temperature"].([]interface{})[1],
		hourly["precipitation_probability"].([]interface{})[1],
		hourly["wind_speed_10m"].([]interface{})[1],
		hourly["precipitation"].([]interface{})[1],
		hourly["uv_index"].([]interface{})[1])

	fmt.Printf(`

		In the next hour expect:

		Temp: %v°C
		Will feel like: %v°C
		Chance of rain/snow %v%%
		Wind speed: %vm/s
		Precipitation %vmm
		UV index %v`,
		hourly["temperature_2m"].([]interface{})[2],
		hourly["apparent_temperature"].([]interface{})[2],
		hourly["precipitation_probability"].([]interface{})[2],
		hourly["wind_speed_10m"].([]interface{})[2],
		hourly["precipitation"].([]interface{})[2],
		hourly["uv_index"].([]interface{})[2])
}
