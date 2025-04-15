package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var materialFlag bool
var exportFlag bool
var pathCmd = &cobra.Command{
	Use: "path",
	Short: "Get paths of uni directory and its courses",
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if exportFlag {
			if len(args) != 0 {
				exit.ExitWithMsg("The --export flag cannot be used with an argument")
			}
			if materialFlag {
				exit.ExitWithMsg("The --export flag cannot be used with the --material flag")
			}

			fmt.Println(exportDirectory)
			return
		}

		if len(args) == 0 {
			fmt.Println(uniDirectory)
			return
		}

		courseId := args[0]
		course, ok := config.Courses[courseId]
		if ok {
			if course.Link != "" {
				fmt.Println(course.Link)
				return
			}

			path := filepath.Join(uniDirectory, courseId)

			if materialFlag {
				materialDir := filepath.Join(path, "material")
				materialDirInfo, err := os.Stat(materialDir)
				if err != nil || !materialDirInfo.IsDir() {
					fmt.Fprintf(os.Stderr, "Error: There is no material directory in course %s.\n", courseId)
					os.Exit(1)
				}
				fmt.Println(materialDir)
				return
			}

			fmt.Println(path)
			return
		}

		fmt.Fprintf(os.Stderr, "Error: There is no course %s.\n", courseId)
		os.Exit(1)
	},
}

func init() {
	pathCmd.Flags().BoolVarP(&materialFlag, "material", "m", false, "material directory")
	pathCmd.Flags().BoolVarP(&exportFlag, "export", "x", false, "global export directory")
}
