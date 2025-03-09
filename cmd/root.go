package cmd

import (
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal"
	"github.com/spf13/cobra"
)

var configFile string
var config internal.Config

var uniDirectory string
var rootCmd = &cobra.Command{
	Use:   "uni",
	Short: "University workflow tool",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(validation)

	rootCmd.PersistentFlags().StringVarP(&uniDirectory, "directory", "d", "", "uni directory (default: ~/uni)")

	rootCmd.AddCommand(courseCmd)
	rootCmd.AddCommand(pathCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(nextCmd)
	rootCmd.AddCommand(configCmd)
}

func validation() {
	if uniDirectory == "" {
		home, err := os.UserHomeDir()
		internal.ExitWithErr(err)

		uniDirectory = filepath.Join(home, "uni")
	}

	uniDirectoryInfo, err := os.Stat(uniDirectory)
	internal.ExitWithErr(err)

	if !uniDirectoryInfo.IsDir() {
		internal.ExitWithMsg("Error: no directory:", uniDirectory)
	}

	configFile = filepath.Join(uniDirectory, "uni-cli.toml")

	configFileInfo, err := os.Stat(configFile)
	internal.ExitWithErr(err)

	if configFileInfo.IsDir() {
		internal.ExitWithMsg("Error: is directory", configFile)
	}

	config = internal.ParseConfig(configFile)
}
