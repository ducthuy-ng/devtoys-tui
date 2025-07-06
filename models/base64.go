package models

import (
	"encoding/base64"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Base64EncryptSubprogram struct {
	textInput   textinput.Model
	encodedText string
}

func (b *Base64EncryptSubprogram) Init() tea.Cmd {
	b.textInput = textinput.New()
	b.textInput.Focus()

	return nil
}

func (b *Base64EncryptSubprogram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEsc:
			return b, tea.Quit
		}

	}

	var cmd tea.Cmd
	b.textInput, cmd = b.textInput.Update(msg)
	b.encodedText = base64.StdEncoding.EncodeToString([]byte(b.textInput.Value()))
	return b, cmd
}

func (b *Base64EncryptSubprogram) View() string {
	view := "Base64 Encrypt Subprogram\n\n"
	view += "Enter text to encode:\n> " + b.textInput.Value() + "\n"

	if b.encodedText == "" {
		view += "Encoded text:\n> "
	} else {
		view += "Encoded text:\n> " + b.encodedText
	}

	view += "\n\nPress Esc to exit."
	return view
}

/* ============================================ */
type Base64DecryptSubprogram struct {
	textInput   textinput.Model
	decodedText string
}

func (b *Base64DecryptSubprogram) Init() tea.Cmd {
	b.textInput = textinput.New()
	b.textInput.Focus()

	return nil
}

func (b *Base64DecryptSubprogram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return b, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd
	b.textInput, cmd = b.textInput.Update(msg)
	decoded, err := base64.StdEncoding.DecodeString(b.textInput.Value())
	if err == nil {
		b.decodedText = string(decoded)
	} else {
		b.decodedText = ""
	}
	return b, cmd
}

func (b *Base64DecryptSubprogram) View() string {
	view := "Base64 Decrypt Subprogram\n\n"
	view += "Enter text to decode:\n> " + b.textInput.Value() + "\n"

	if b.decodedText == "" {
		view += "Decoded text:\n> "
	} else {
		view += "Decoded text:\n> " + b.decodedText
	}

	view += "\n\nPress Esc to exit."
	return view
}
