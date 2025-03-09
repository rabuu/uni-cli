package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use: "next",
	Short: "Generate next working directory",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		course := getCurrentCourse()

		prefix := config.Courses[course].Prefix
		number := testNextDir(prefix, course)
		nextDir := fmt.Sprintf("%s%02d", prefix, number)

		err := os.Mkdir(nextDir, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		templateDirPath := getTemplateDirPath(course)
		filepath.WalkDir(templateDirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}

			if d.IsDir() {
				stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
				if stripped == "" {
					return nil
				}

				target := filepath.Join(uniDirectory, course, nextDir, stripped)
				err := os.Mkdir(target, 0755)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error:", err)
					os.Exit(1)
				}
			}

			if !d.Type().IsRegular() {
				return nil
			}

			templ, err := template.ParseFiles(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ParseError:", err)
				os.Exit(1)
			}

			stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
			target := filepath.Join(uniDirectory, course, nextDir, stripped)

			data := struct {
				Course, CourseName string
				Number int
				NumberPadded string
				Group []cfg.GroupMember
			}{
				Course: course,
				CourseName: config.Courses[course].FullName,
				Number: number,
				NumberPadded: fmt.Sprintf("%02d", number),
				Group: config.Courses[course].Group,
			}

			file, err := os.Create(target)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}

			err = templ.Execute(file, data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}

			return nil
		})
	},
}

func getCurrentCourse() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	dir := filepath.Dir(cwd)
	if dir != uniDirectory {
		fmt.Fprintln(os.Stderr, "You must be in a course directory")
		os.Exit(1)
	}

	course := filepath.Base(cwd)
	if !config.ContainsCourse(course) {
		fmt.Fprintln(os.Stderr, "You must be in a course directory")
		os.Exit(1)
	}

	return course
}

func testNextDir(prefix string, course string) int {
	for i := 0; i <= 99; i++ {
		testDir := fmt.Sprintf("%s%02d", prefix, i)
		testDirPath := filepath.Join(uniDirectory, course, testDir)

		_, err := os.Stat(testDirPath)
		if os.IsNotExist(err) {
			return i
		}
	}

	fmt.Fprintln(os.Stderr, "Too many directories")
	os.Exit(1)

	return -1
}

func getTemplateDirPath(course string) string {
	templateDirPath := filepath.Join(uniDirectory, course, "template")
	templateDirInfo, err := os.Stat(templateDirPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if !templateDirInfo.IsDir() {
		fmt.Fprintln(os.Stderr, "There is no template directory")
		os.Exit(1)
	}

	return templateDirPath
}
