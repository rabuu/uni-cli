package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/spf13/cobra"
)

var courseCmd = &cobra.Command{
	Use: "course",
	Short: "Manage courses",
}

var listCoursesCmd = &cobra.Command{
	Use: "list",
	Short: "List all courses",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintCourses()
	},
}

var newFullName string
var addCourseCmd = &cobra.Command{
	Use: "add",
	Short: "Add a course",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		var course cfg.Course
		if newFullName != "" {
			course.FullName = newFullName
		}

		if config.ContainsCourse(name) {
			fmt.Printf("The course %s already exists.\n", name)
			os.Exit(0)
		}

		config.Courses[name] = course
		config.WriteToFile(configFile)

		newCourseDirectory := filepath.Join(uniDirectory, name)
		err := os.MkdirAll(newCourseDirectory, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Success: Added course %s.\n", name)
	},
}

func init() {
	addCourseCmd.Flags().StringVarP(&newFullName, "full-name", "f", "", "provide a more detailed name")

	courseCmd.AddCommand(listCoursesCmd)
	courseCmd.AddCommand(addCourseCmd)
}
