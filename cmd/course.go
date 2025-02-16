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

var newLongName string
var addCourseCmd = &cobra.Command{
	Use: "add",
	Short: "Add a course",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newName := args[0]

		var newCourse cfg.Course
		newCourse.Name = newName

		if newLongName != "" {
			newCourse.LongName = newLongName
		}

		if config.ContainsCourse(newCourse.Name) {
			fmt.Printf("The course %s already exists.\n", newCourse.Name)
			os.Exit(0)
		}

		config.Courses = append(config.Courses, newCourse)
		config.WriteToFile(configFile)

		newCourseDirectory := filepath.Join(uniDirectory, newCourse.Name)
		err := os.Mkdir(newCourseDirectory, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Success: Added course '%s'.\n", newCourse.Name)
	},
}

func init() {
	addCourseCmd.Flags().StringVarP(&newLongName, "long-name", "l", "", "provide a more detailed name")

	courseCmd.AddCommand(addCourseCmd)
}
