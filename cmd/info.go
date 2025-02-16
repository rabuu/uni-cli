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
		fmt.Printf("Uni Directory:      %s\n", uniDirectory)
		fmt.Printf("Configuration File: %s\n", configFile)
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Courses:")
		for i := 0; i < len(conf.Courses); i++ {
			course := conf.Courses[i]
			if course.LongName == "" {
				fmt.Printf("  %d. %s\n", i + 1, course.Name)
			} else {
				fmt.Printf("  %d. %s (%s)\n", i + 1, course.Name, course.LongName)
			}
		}
	},
}
