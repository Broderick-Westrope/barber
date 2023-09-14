package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	yesFlag bool // skip confirmation prompts
	collectionFlag  string // path to collection root directory
	keepFlag bool // keep snippets that are not found in the filesystem
	dryRunFlag bool // display proposed changes without performing them
)

// The base command.
// When called without any subcommands, it will launch interactive mode.
var rootCmd = &cobra.Command{
	Use:   "barber",
	Short: "Barber is a tool for managing snippets",
	Long: `Barber is a CLI & TUI for managing snippets.
It is targeted towards code snippets, but can be used for any text format.
Documentation is available at https://github.com/Broderick-Westrope/barber`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Launch interactive mode
		fmt.Println("Interactive mode not yet implemented")
	},
}

// Execute runs the root command.
// This is the entry point for the CLI.
func Execute() {
	// Collection
	collectionCmd.PersistentFlags().StringVarP(&collectionFlag, "collection", "c", ".", "path to collection root directory")
	rootCmd.AddCommand(collectionCmd)

	// Collection Init
	collectionCmd.AddCommand(colInitCmd)

	// Collection Reset
	colResetCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "confirm removal without prompting")
	collectionCmd.AddCommand(colResetCmd)

	// Collection Remove
	colRemoveCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "confirm removal without prompting")
	collectionCmd.AddCommand(colRemoveCmd)

	// Snippet
	snippetCmd.PersistentFlags().StringVarP(&collectionFlag, "collection", "c", ".", "path to collection root directory")
	rootCmd.AddCommand(snippetCmd)

	// Snippet Add
	snippetCmd.AddCommand(snptAddCmd)

	// Snippet Remove
	snippetCmd.AddCommand(snptRemoveCmd)

	// Scan
	scanCmd.Flags().StringVarP(&collectionFlag, "collection", "c", ".", "path to collection root directory")
	scanCmd.Flags().BoolVar(&keepFlag, "keep", false, "keep snippets that are not found in the filesystem")
	scanCmd.Flags().BoolVar(&dryRunFlag, "dry-run", false, "display proposed changes without performing them")
	rootCmd.AddCommand(scanCmd)

	cobra.CheckErr(rootCmd.Execute())
}

// Displays the help message for a command.
func displayHelp(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("Failed to display help: %v", err)
	}
}