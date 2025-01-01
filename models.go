package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor       int
	choices      []string
	screen       string
	forecastHour int
	city         string
}

func initialModel() model {
	return model{
		choices: getCityCodes(cityList),
		screen:  "city",
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.screen {
			case "city":
				m.city = m.choices[m.cursor]
				m.screen = "hours"
				m.cursor = 0
				m.choices = []string{"Now", "1 hour"}
			case "hours":
				m.forecastHour = m.cursor
				m.screen = "weather"
				m.cursor = 0
			case "weather":
			}
		}
	}

	return m, nil
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	s := "Check weather\n\n"

	switch m.screen {
	case "city":
		s += citiesView(m)
	case "hours":
		s += hoursView(m)
	case "weather":
		s += weatherView(m)
	}

	s += "\n\nPress q to quit"
	return s
}

func citiesView(m model) string {
	s := "Choose city\n\n"
	return s + choicesView(m.choices, m)
}

func hoursView(m model) string {
	s := "Choose hours\n\n"
	return s + choicesView(m.choices, m)
}

func weatherView(m model) string {
	var s string
	rawResp := getRespBody(weatherAPICall(m.city))
	respData := typefyResp(rawResp)
	hourlyForecast := respData["hourly"].(map[string]interface{})
	s = cityList[m.city].name + "\n"
	s += dumpWeatherDigest(hourlyForecast, m.forecastHour)
	return s
}

func choicesView(choices []string, m model) string {
	var s string
	for choice := range choices {
		cursor := " "
		if m.cursor == choice {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choices[choice])
	}

	return s
}
