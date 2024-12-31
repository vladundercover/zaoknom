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

	// rawResp := getRespBody(weatherAPICall("MSK"))
	// respData := typefyResp(rawResp)
	// hourlyForecast := respData["hourly"].(map[string]interface{})

	// fmt.Println(dumpWeatherDigest(hourlyForecast, 0))
}
