package cmd

import (
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/cfgfile"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var configFile string
var config cfgfile.Config

var exportDirectory string

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
	rootCmd.AddCommand(exportCmd)
}

func validation() {
	if uniDirectory == "" {
		home, err := os.UserHomeDir()
		exit.ExitWithErr(err)

		uniDirectory = filepath.Join(home, "uni")
	}

	uniDirectoryInfo, err := os.Stat(uniDirectory)
	exit.ExitWithErr(err)

	if !uniDirectoryInfo.IsDir() {
		exit.ExitWithMsg("Error: no directory:", uniDirectory)
	}

	configFile = filepath.Join(uniDirectory, "uni-cli.toml")

	configFileInfo, err := os.Stat(configFile)
	exit.ExitWithErr(err)

	if configFileInfo.IsDir() {
		exit.ExitWithMsg("Error: is directory", configFile)
	}

	config = cfgfile.ParseConfig(configFile, uniDirectory)
	exportDirectory = filepath.Join(uniDirectory, config.ExportDir)
}
