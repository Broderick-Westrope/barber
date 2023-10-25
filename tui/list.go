package tui

import (
	"fmt"
	"io"

	"github.com/Broderick-Westrope/barber/file"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func newSnippetList(items []list.Item, height int, styles *snippetBaseStyle) *list.Model {
	snippetList := list.New(items, snippetDelegate{styles: styles}, 25, height)
	snippetList.SetShowHelp(false)
	snippetList.SetShowFilter(false)
	snippetList.SetShowTitle(false)
	snippetList.Styles.StatusBar = lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("240")).MaxWidth(35 - 2)
	snippetList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color("8")).MaxWidth(35 - 2)
	snippetList.FilterInput.Prompt = "Find: "
	snippetList.SetStatusBarItemName("snippet", "snippets")
	snippetList.DisableQuitKeybindings()

	return &snippetList
}

type snippetDelegate struct {
	styles *snippetBaseStyle
}

func (d snippetDelegate) Height() int {
	return 2
}

func (d snippetDelegate) Spacing() int {
	return 1
}

func (d snippetDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d snippetDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if item == nil {
		return
	}
	s, ok := item.(file.Snippet)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprintln(w, d.styles.SelectedTitle.Render(s.Path))
		return
	}
	fmt.Fprintln(w, d.styles.UnselectedTitle.Render(s.Path))
}
