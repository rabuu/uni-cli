package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Manage the configuration file",
}

var printCmd = &cobra.Command{
	Use: "print",
	Short: "Print the configuration file",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintToStdout()
	},
}

var rewriteCmd = &cobra.Command{
	Use: "rewrite",
	Short: "Rewrite the configuration file",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteToFile(configFile)
	},
}

func init() {
	configCmd.AddCommand(printCmd)
	configCmd.AddCommand(rewriteCmd)
}
