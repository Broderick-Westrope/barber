package tui

import "github.com/charmbracelet/bubbles/key"

// TODO: Look into adding a way to specify a keymap for a specific view/pane for more specific help text.
type KeyMap struct {
	Quit       key.Binding
	ToggleHelp key.Binding
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit:       key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	ToggleHelp: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	Left:       key.NewBinding(key.WithKeys("h", "left"), key.WithHelp("h", "left")),
	Right:      key.NewBinding(key.WithKeys("l", "right"), key.WithHelp("l", "right")),
	Up:         key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k", "up")),
	Down:       key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j", "down")),
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
