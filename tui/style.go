package tui

import (
	"github.com/Broderick-Westrope/barber/file"
	"github.com/charmbracelet/lipgloss"
)

type styles struct {
	Snippet snippetStyle
}

// Contains the style structs for the snippets pane when it is focused and unfocused.
type snippetStyle struct {
	Focused snippetBaseStyle
}

// Contains styles for each element of the snippets pane.
type snippetBaseStyle struct {
	SelectedTitle      lipgloss.Style
	UnselectedTitle    lipgloss.Style
	SelectedSubtitle   lipgloss.Style
	UnselectedSubtitle lipgloss.Style
}

// TODO: Add custom color settings using config file
func defaultStyles(config *file.Config) *styles {
	// white := lipgloss.Color(config.Collection.Style.WhiteColor)
	gray := lipgloss.Color(config.Collection.Style.GrayColor)
	// black := lipgloss.Color(config.Collection.Style.BackgroundColor)
	// brightBlack := lipgloss.Color(config.Collection.Style.BlackColor)
	// green := lipgloss.Color(config.Collection.Style.GreenColor)
	// brightGreen := lipgloss.Color(config.Collection.Style.BrightGreenColor)
	brightBlue := lipgloss.Color(config.Collection.Style.PrimaryColor)
	blue := lipgloss.Color(config.Collection.Style.PrimaryColorSubdued)
	// red := lipgloss.Color(config.Collection.Style.RedColor)
	// brightRed := lipgloss.Color(config.Collection.Style.BrightRedColor)

	return &styles{
		Snippet: snippetStyle{
			Focused: snippetBaseStyle{
				SelectedTitle:      lipgloss.NewStyle().Foreground(brightBlue),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(gray),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(blue),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
			},
		},
	}
}
