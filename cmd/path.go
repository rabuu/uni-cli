package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/workingdir"
	"github.com/spf13/cobra"
)

var materialFlag bool
var exportFlag bool
var pathCmd = &cobra.Command{
	Use: "path",
	Short: "Get paths of uni directory and its courses",
	Args: cobra.RangeArgs(0, 2),
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

		if materialFlag {
			if len(args) != 1 {
				exit.ExitWithMsg("The --material flag can only be used for course directories")
			}
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
					exit.ExitWithMsg("Error: There is no material directory in course", courseId)
				}
				fmt.Println(materialDir)
				return
			}

			if len(args) == 2 {
				i, err := strconv.Atoi(args[1])
				if err != nil {
					exit.ExitWithErr(err)
				}

				workingDirName := workingdir.FromNumber(i, course.Prefix)

				workingDir := filepath.Join(path, workingDirName)
				workingDirInfo, err := os.Stat(workingDir)
				if err != nil || !workingDirInfo.IsDir() {
					exit.ExitWithMsg("Error: There is no directory", workingDirName, "in course", courseId)
				}

				fmt.Println(workingDir)
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
