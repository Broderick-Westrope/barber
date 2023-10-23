package internal

import (
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
	dirList     list.Model
	snippetList list.Model

	listStyle lipgloss.Style
}

type snippetDelegate struct{}

func NewModel(rootDir *Directory) tea.Model {
	snippetItems := GetItems(rootDir.Snippets)
	dirItems := GetItems(rootDir.SubDirs)

	listDelegate := list.NewDefaultDelegate()
	listDelegate.ShowDescription = false

	m := &Model{
		keys:        DefaultKeyMap,
		help:        help.New(),
		currentDir:  rootDir,
		snippetList: list.New(snippetItems, listDelegate, 0, 0),
		dirList:     list.New(dirItems, listDelegate, 0, 0),
		listStyle:   lipgloss.NewStyle().Margin(1, 2),
	}

	m.snippetList.Title = "Snippets"

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
		m.dirList.View(),
		m.snippetList.View())

	return lipgloss.JoinVertical(lipgloss.Top,
		panes,
		m.help.View(m.keys))
}
