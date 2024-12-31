package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// rawResp := GetRespBody(GetData("MSK"))
	// respData := TypefyResp(rawResp)
	// hourlyForecast := respData["hourly"].(map[string]interface{})

	// DumpWeatherDigest(hourlyForecast, 0)
	// DumpWeatherDigest(hourlyForecast, 1)
}
