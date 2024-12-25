package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	api_call := "https://api.open-meteo.com/v1/forecast?latitude=55.7522&longitude=37.6156&current=temperature_2m,relative_humidity_2m,apparent_temperature,is_day,weather_code,surface_pressure,wind_speed_10m&hourly=precipitation_probability,uv_index,is_day&wind_speed_unit=ms&timezone=Europe%2FMoscow&forecast_days=1"
	resp, _ := http.Get(api_call)
	body, _ := io.ReadAll(resp.Body)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	current := dat["current"].(map[string]interface{})
	fmt.Printf(`Current temp: %v°C
Feels like: %v°C
Wind speed: %vm/s
`,
		current["temperature_2m"],
		current["apparent_temperature"],
		current["wind_speed_10m"])
}
