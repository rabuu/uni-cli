package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/spf13/cobra"
)

var configFile string
var config cfg.Config

var uniDirectory string
var rootCmd = &cobra.Command{
	Use:   "uni",
	Short: "University workflow tool",
	// Run: func(cmd *cobra.Command, args []string) { },
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
}

func validation() {
	if uniDirectory == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		// TODO: change this to ~/uni when it's ready
		uniDirectory = filepath.Join(home, "uni-test")
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

	config = cfg.ParseConfig(configFile)
}
