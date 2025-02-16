package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use: "path",
	Short: "Get paths of uni directory and its courses",
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(uniDirectory)
			return
		}

		course := args[0]
		if config.ContainsCourse(course) {
			path := filepath.Join(uniDirectory, course)
			fmt.Println(path)
			return
		}

		fmt.Fprintf(os.Stderr, "Error: There is no course %s.", course)
		os.Exit(1)
	},
}
