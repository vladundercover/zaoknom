package main

import (
	h "zaoknom/helpers"
	"zaoknom/weather"
)

func main() {
	rawResp := h.RespBody(weather.APICallFor("MSK"))
	respData := h.TypefyResp(rawResp)
	hourlyForecast := respData["hourly"].(map[string]interface{})

	h.DumpWeatherDigest(hourlyForecast, 0)
	h.DumpWeatherDigest(hourlyForecast, 1)
}
