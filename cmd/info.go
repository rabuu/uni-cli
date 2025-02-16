package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use: "info",
	Short: "Show information about the uni directory",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Uni Directory: %s\n", mainConfig.GetString("uni-directory"))
	},
}
