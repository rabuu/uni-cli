package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mainConfig = viper.New()

var mainConfigFile string
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
	cobra.OnInitialize(initMainConfig)

	rootCmd.PersistentFlags().StringVarP(&mainConfigFile, "config", "c", "", "configuration file (default: ~/.config/uni-cli/uni-cli.toml)")

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(infoCmd)
}

func initMainConfig() {
	if mainConfigFile != "" {
		if _, err := os.Stat(mainConfigFile); err == nil {
			mainConfig.SetConfigFile(mainConfigFile)
		} else {
			log.Fatal(err)
		}
	} else {
		mainConfig.SetConfigName("uni-cli")
		mainConfig.SetConfigType("toml")
		mainConfig.AddConfigPath("$HOME/.config/uni-cli/")
		mainConfig.AddConfigPath("$XDG_CONFIG_HOME/uni-cli/")
		mainConfig.AddConfigPath(".")
	}

	err := mainConfig.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	uniDirInfo, err := os.Stat(mainConfig.GetString("uni-directory"))
	if err != nil {
		log.Fatal(err)
	}

	if !uniDirInfo.IsDir() {
		log.Fatal(fmt.Sprintf("no directory: %s", mainConfig.GetString("uni-directory")))
	}
}
