package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "barber",
	Short: "Barber is a tool for managing snippets",
	Long: `Barber is a CLI & TUI for managing snippets.
			It is targeted towards code snippets, but can be used for any text format.
			Documentation is available at https://github.com/Broderick-Westrope/barber`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Launch interactive mode
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
