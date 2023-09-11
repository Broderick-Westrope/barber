package cmd

import (
	"github.com/spf13/cobra"
)

var snippetCmd = &cobra.Command{
	Use: "snippet",
	Aliases: []string{"snpt"},
	Short: "Manage snippets within a collection",
	Run: displayHelp,
}