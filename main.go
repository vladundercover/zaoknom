package main

import (
	"zaoknom/helpers"
	"zaoknom/weather"
)

func main() {
	rawResp := helpers.GetRespBody(weather.GetData("MSK"))
	respData := helpers.TypefyResp(rawResp)
	hourlyForecast := respData["hourly"].(map[string]interface{})

	weather.DumpWeatherDigest(hourlyForecast, 0)
	weather.DumpWeatherDigest(hourlyForecast, 1)
}
