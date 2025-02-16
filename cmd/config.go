package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Read and write configuration data",
}

var getUniDirectoryCmd = &cobra.Command{
	Use: "get-uni-directory",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(mainConfig.GetString("uni-directory"))
	},
}

var setUniDirectoryCmd = &cobra.Command{
	Use: "set-uni-directory",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mainConfig.Set("uni-directory", args[0])
		err := mainConfig.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	configCmd.AddCommand(getUniDirectoryCmd)
	configCmd.AddCommand(setUniDirectoryCmd)
}
