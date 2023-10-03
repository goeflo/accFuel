package fuel

import "github.com/charmbracelet/bubbles/key"

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view.
// It's part of the key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Help, k.Quit},
	}
}

type keyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
	Help key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "increase extra lap"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "decrease extra lap"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}
