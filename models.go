package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

type cityListModel struct {
	choices []string // items on the list
	chosen  bool     // what is chosen
	cursor  int      // which item our cursor is pointing at
}

func initialModel() cityListModel {
	return cityListModel{
		// Our city code list
		choices: getCityCodes(CityList),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Zaoknom")
}

// Updates
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
