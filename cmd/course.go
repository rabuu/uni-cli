package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var courseCmd = &cobra.Command{
	Use: "course",
	Aliases: []string{"c"},
	Short: "Manage the registered courses",
}

var fishFlag bool
var listCoursesCmd = &cobra.Command{
	Use: "list",
	Short: "List all courses",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if fishFlag {
			config.PrintCoursesFishCompletion()
		} else {
			config.PrintCoursesHumanReadable()
		}
	},
}

var newFullName string
var newPrefix string
var newLink string
var addCourseCmd = &cobra.Command{
	Use: "add",
	Short: "Add a course",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if config.ContainsCourse(name) {
			fmt.Printf("The course %s already exists.\n", name)
			os.Exit(0)
		}

		if newFullName == "" {
			newFullName = name
		}

		if newLink != "" {
			newPrefix = ""
		}

		var course cfg.Course
		course.Name = newFullName
		course.Prefix = newPrefix
		course.Link = newLink

		config.Courses[name] = course
		config.WriteToFile(configFile)

		if newLink == "" {
			newCourseDirectory := filepath.Join(uniDirectory, name)
			err := os.MkdirAll(newCourseDirectory, 0755)
			exit.ExitWithErr(err)
		}

		fmt.Printf("Success: Added course %s.\n", name)
	},
}

var deleteRemovedCourse bool
var removeCourseCmd = &cobra.Command{
	Use: "remove",
	Aliases: []string { "rm" },
	Short: "Remove a course",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if !config.ContainsCourse(name) {
			fmt.Printf("The course %s does not exist.\n", name)
			os.Exit(0)
		}

		delete(config.Courses, name)
		config.WriteToFile(configFile)

		if deleteRemovedCourse {
			courseDirectory := filepath.Join(uniDirectory, name)
			err := os.RemoveAll(courseDirectory)
			exit.ExitWithErr(err)
		}

		fmt.Printf("Success: Removed course %s.\n", name)
	},
}

func init() {
	listCoursesCmd.Flags().BoolVar(&fishFlag, "fish", false, "fish completion syntax")
	addCourseCmd.Flags().StringVarP(&newFullName, "name", "n", "", "the full course name")
	addCourseCmd.Flags().StringVarP(&newPrefix, "prefix", "p", "", "the prefix of the working directories")
	addCourseCmd.Flags().StringVarP(&newLink, "link", "l", "", "a link to another directory location")
	removeCourseCmd.Flags().BoolVarP(&deleteRemovedCourse, "delete", "D", false, "delete course directory")

	courseCmd.AddCommand(listCoursesCmd)
	courseCmd.AddCommand(addCourseCmd)
	courseCmd.AddCommand(removeCourseCmd)
}
