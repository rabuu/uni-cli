package cmd

import (
	"github.com/spf13/cobra"
)

var rewriteCmd = &cobra.Command{
	Use: "rewrite",
	Short: "Rewrite the configuration file",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteToFile(configFile)
	},
}
