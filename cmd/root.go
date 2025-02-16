package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var uniDir string
var rootCmd = &cobra.Command{
	Use:   "uni",
	Short: "University workflow tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uniDir)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&uniDir, "uni-directory", "d", "", "main uni directory (default is $HOME/uni)")

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


