package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ducthuy-ng/devtoys-tui/models"
)

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type DevToysModelFeatures struct {
	title       string
	description string
	Subprogram  models.Subprogram
}

func (f DevToysModelFeatures) FilterValue() string { return f.title }
func (f DevToysModelFeatures) Title() string       { return f.title }
func (f DevToysModelFeatures) Description() string { return f.description }

type DevToysModel struct {
	featureList  list.Model
	programStyle lipgloss.Style

	currentSubprogram models.Subprogram // The currently active subprogram
}

func initialModel() DevToysModel {
	model := DevToysModel{
		programStyle: lipgloss.NewStyle().Padding(2),
	}
	items := []list.Item{
		DevToysModelFeatures{title: "Base64 Encode", description: "Encode text to Base64", Subprogram: &models.Base64EncryptSubprogram{}},
		DevToysModelFeatures{title: "Base64 Decode", description: "Decode Base64 to text", Subprogram: &models.Base64DecryptSubprogram{}},
	}
	model.featureList = list.New(items, list.NewDefaultDelegate(), 50, 20)
	model.featureList.Title = "DevToys Features"
	model.featureList.DisableQuitKeybindings()

	return model
}

func (m DevToysModel) Init() tea.Cmd {
	return nil
}

func (m DevToysModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.currentSubprogram != nil {
			// Exit the program
			if msg.Type == tea.KeyEsc {
				m.currentSubprogram = nil
				return m, nil
			}

			_, cmd := m.currentSubprogram.Update(msg)
			return m, cmd
		}

		if msg.Type == tea.KeyEsc {
			return m, tea.Quit
		}

		if msg.Type == tea.KeyEnter {
			m.currentSubprogram = m.featureList.SelectedItem().(DevToysModelFeatures).Subprogram
			m.currentSubprogram.Init()
		}
	}

	var cmd tea.Cmd
	m.featureList, cmd = m.featureList.Update(msg)
	return m, cmd
}

func (m DevToysModel) View() string {
	if m.currentSubprogram != nil {
		// If a subprogram is active, delegate the view to it
		return m.programStyle.Render(m.currentSubprogram.View())
	}

	view := "Welcome to DevToys TUI!\n\n"
	view += m.featureList.View()
	view += "\n\n"
	view += "Press Esc to exit."
	return m.programStyle.Render(view)
}
