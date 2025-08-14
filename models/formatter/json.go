package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ducthuy-ng/devtoys-tui/shared"
)

type JsonFormatter struct {
	inputTextArea     textarea.Model
	formattedTextArea textarea.Model
	encodedText       string

	keys shared.KeyMap
	help help.Model
}

func (b *JsonFormatter) Init() tea.Cmd {
	b.inputTextArea = textarea.New()
	b.inputTextArea.Focus()

	b.formattedTextArea = textarea.New()

	b.keys = shared.GetDefaultKeyMap()
	b.help = help.New()

	b.encodedText = ""

	return nil
}

func (b *JsonFormatter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.inputTextArea.SetWidth(msg.Width - 5)
		b.inputTextArea.SetHeight(msg.Height/2 - 5)

		b.formattedTextArea.SetWidth(msg.Width - 5)
		b.formattedTextArea.SetHeight(msg.Height/2 - 5)
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, b.keys.Copy):
			err := clipboard.WriteAll(b.formattedTextArea.Value())
			if err != nil {
				panic(err)
			}
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd

	b.inputTextArea, cmd = b.inputTextArea.Update(msg)

	var formattedValue bytes.Buffer
	err := json.Indent(&formattedValue, []byte(b.inputTextArea.Value()), "", "  ")
	if err != nil {
		b.formattedTextArea.SetValue("Invalid JSON input")
	}
	b.formattedTextArea.SetValue(formattedValue.String())
	return b, cmd
}

func (b *JsonFormatter) View() string {
	view := "Base64 Encrypt Subprogram\n\n"

	view += b.inputTextArea.View()
	view += "\n\n"

	view += b.formattedTextArea.View()
	view += "\n\n"

	view += b.help.View(&b.keys)

	return view
}
