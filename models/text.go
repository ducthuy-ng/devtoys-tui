package models

import (
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextEscape struct {
	textArea       textarea.Model
	resultTextArea textarea.Model
}

func (prog *TextEscape) Init() tea.Cmd {
	prog.textArea = textarea.New()
	prog.textArea.Focus()

	prog.resultTextArea = textarea.New()

	return nil
}
func (prog *TextEscape) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.Type {
		case tea.KeyCtrlP:
			err := clipboard.WriteAll(prog.resultTextArea.Value())
			if err != nil {
				panic(err)
			}

		case tea.KeyTab:
			if prog.textArea.Focused() {
				prog.textArea.Blur()
				prog.resultTextArea.Focus()
			} else {
				prog.resultTextArea.Blur()
				prog.textArea.Focus()
			}
		}
	}

	var cmd tea.Cmd
	prog.textArea, cmd = prog.textArea.Update(msg)
	return prog, cmd
}

func (b *TextEscape) View() string {
	view := "Base64 Encrypt Subprogram\n\n"
	view += b.textArea.View()

	view += "\n\n"
	view += "Escaped text:\n" + b.resultTextArea.View() + "\n"
	return view
}
