package models

import (
	"encoding/base64"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ducthuy-ng/devtoys-tui/shared"
)

type Base64EncryptSubprogram struct {
	inputTextArea textarea.Model
	encodedText   string

	keys shared.KeyMap
	help help.Model
}

func (b *Base64EncryptSubprogram) Init() tea.Cmd {
	b.inputTextArea = textarea.New()
	b.inputTextArea.Focus()

	b.keys = shared.GetDefaultKeyMap()
	b.help = help.New()

	b.encodedText = ""

	return nil
}

func (b *Base64EncryptSubprogram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, b.keys.Copy):
			err := clipboard.WriteAll(b.encodedText)
			if err != nil {
				panic(err)
			}
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd
	b.inputTextArea, cmd = b.inputTextArea.Update(msg)
	b.encodedText = base64.StdEncoding.EncodeToString([]byte(b.inputTextArea.Value()))
	return b, cmd
}

func (b *Base64EncryptSubprogram) View() string {
	view := "Base64 Encrypt Subprogram\n\n"

	view += b.inputTextArea.View()
	view += "\n\n"

	if b.encodedText == "" {
		view += "Encoded text:\n> "
	} else {
		view += "Encoded text:\n> " + b.encodedText
	}
	view += "\n\n"

	view += b.help.View(&b.keys)

	return view
}

/* ============================================ */
type Base64DecryptSubprogram struct {
	inputTextArea  textarea.Model
	resultTextArea textarea.Model

	keys shared.KeyMap
	help help.Model
}

func (b *Base64DecryptSubprogram) Init() tea.Cmd {
	b.inputTextArea = textarea.New()
	b.inputTextArea.Focus()

	b.resultTextArea = textarea.New()

	b.keys = shared.GetDefaultKeyMap()
	b.help = help.New()

	return nil
}

func (b *Base64DecryptSubprogram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, b.keys.Copy):
			err := clipboard.WriteAll(b.resultTextArea.Value())
			if err != nil {
				panic(err)
			}
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd
	b.inputTextArea, cmd = b.inputTextArea.Update(msg)
	decoded, err := base64.StdEncoding.DecodeString(b.inputTextArea.Value())
	if err == nil {
		b.resultTextArea.SetValue(string(decoded))
	} else {
		b.resultTextArea.SetValue("")
	}
	return b, cmd
}

func (b *Base64DecryptSubprogram) View() string {
	view := "Base64 Decrypt Subprogram\n\n"

	view += b.inputTextArea.View()
	view += "\n\n"

	view += "Decoded text:\n"
	view += b.resultTextArea.View()
	view += "\n\n"

	view += b.help.View(&b.keys)

	return view
}
