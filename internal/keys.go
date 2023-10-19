package internal

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit       key.Binding
	ToggleHelp key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit:       key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	ToggleHelp: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
		k.ToggleHelp,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.ToggleHelp},
	}
}
