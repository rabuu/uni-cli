package cmd

import (
	"fmt"
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
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		uniDirectory = filepath.Join(home, "uni")
	}

	uniDirectoryInfo, err := os.Stat(uniDirectory)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if !uniDirectoryInfo.IsDir() {
		fmt.Fprintln(os.Stderr, "Error: no directory:", uniDirectory)
		os.Exit(1)
	}

	configFile = filepath.Join(uniDirectory, "uni-cli.toml")

	configFileInfo, err := os.Stat(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if configFileInfo.IsDir() {
		fmt.Fprintln(os.Stderr, "Error: is directory", configFile)
		os.Exit(1)
	}

	config = internal.ParseConfig(configFile)
}
