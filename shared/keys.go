package shared

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	// Keys map[string]key.Binding
	Copy key.Binding
	Quit key.Binding
}

func GetDefaultKeyMap() KeyMap {
	return KeyMap{

		Copy: key.NewBinding(
			key.WithKeys("ctrl+p"),
			key.WithHelp("ctrl+p", "copy to clipboard"),
		),
		Quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "quit"),
		),
	}
}

func (k *KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Copy, k.Quit}
}

func (k *KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Copy},
		{k.Quit},
	}
}
