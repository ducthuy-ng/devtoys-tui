package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Subprogram is an interface that all subprograms should implement.
// After user select a subprogram (e.g., Base64 Encode), the main program
// will switch its state to the selected subprogram.
// Also, the method `Init` is called to initialize the subprogram, if needed.
//
// Then, during a subprogram's lifecycle, the main program (main model) will
// redirect Bubble Tea messages to the subprogram's Update method.
type Subprogram interface {
	Init() tea.Cmd // Initialize the subprogram, if needed
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
}
