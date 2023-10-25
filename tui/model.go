package tui

import (
	"github.com/Broderick-Westrope/barber/file"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys KeyMap
	help help.Model

	currentDir  *Directory
	snippetList list.Model

	listStyle lipgloss.Style
}

func NewModel(rootDir *Directory) tea.Model {
	snippetItems := GetItems(rootDir.Snippets)

	styles := defaultStyles(file.DefaultConfig())

	m := &Model{
		keys:        DefaultKeyMap,
		help:        help.New(),
		currentDir:  rootDir,
		snippetList: *newSnippetList(snippetItems, 2, &styles.Snippet.Focused),
		listStyle:   lipgloss.NewStyle().Margin(1, 2),
	}

	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := m.listStyle.GetFrameSize()
		m.snippetList.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.snippetList, cmd = m.snippetList.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	panes := lipgloss.JoinHorizontal(lipgloss.Left,
		m.snippetList.View())

	return lipgloss.JoinVertical(lipgloss.Top,
		panes,
		m.help.View(m.keys))
}
