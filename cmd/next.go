package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rabuu/uni-cli/internal"
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
		internal.ExitWithErr(err)

		templateDirPath := getTemplateDirPath(course)
		filepath.WalkDir(templateDirPath, func(path string, d fs.DirEntry, err error) error {
			internal.ExitWithErr(err)

			if d.IsDir() {
				stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
				if stripped == "" {
					return nil
				}

				target := filepath.Join(uniDirectory, course, nextDir, stripped)
				err := os.Mkdir(target, 0755)
				internal.ExitWithErr(err)
			}

			if !d.Type().IsRegular() {
				return nil
			}

			templ, err := template.ParseFiles(path)
			internal.ExitWithErr(err)

			stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
			target := filepath.Join(uniDirectory, course, nextDir, stripped)

			data := struct {
				Course, CourseName string
				Number int
				NumberPadded string
				Group []internal.GroupMember
			}{
				Course: course,
				CourseName: config.Courses[course].FullName,
				Number: number,
				NumberPadded: fmt.Sprintf("%02d", number),
				Group: config.Courses[course].Group,
			}

			file, err := os.Create(target)
			internal.ExitWithErr(err)

			err = templ.Execute(file, data)
			internal.ExitWithErr(err)

			return nil
		})
	},
}

func getCurrentCourse() string {
	cwd, err := os.Getwd()
	internal.ExitWithErr(err)

	dir := filepath.Dir(cwd)
	if dir != uniDirectory {
		internal.ExitWithMsg("You must be in a course directory")
	}

	course := filepath.Base(cwd)
	if !config.ContainsCourse(course) {
		internal.ExitWithMsg("You must be in a course directory")
	}

	return course
}

func testNextDir(prefix string, course string) int {
	for i := 1; i <= 99; i++ {
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
	internal.ExitWithErr(err)

	if !templateDirInfo.IsDir() {
		internal.ExitWithMsg("There is no template directory")
	}

	return templateDirPath
}
