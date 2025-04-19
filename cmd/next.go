package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rabuu/uni-cli/internal/cwd"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/templdata"
	"github.com/rabuu/uni-cli/internal/workingdir"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use: "next",
	Aliases: []string{"n"},
	Short: "Generate next working directory",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		course := cwd.CourseDir(uniDirectory, &config)

		prefix := config.Courses[course].Prefix
		number := testNextDir(prefix, course)
		nextDir := workingdir.FromNumber(number, prefix)

		err := os.Mkdir(nextDir, 0755)
		exit.ExitWithErr(err)

		templateDirPath := getTemplateDirPath(course)
		filepath.WalkDir(templateDirPath, func(path string, d fs.DirEntry, err error) error {
			exit.ExitWithErr(err)

			if d.IsDir() {
				stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
				if stripped == "" {
					return nil
				}

				target := filepath.Join(uniDirectory, course, nextDir, stripped)
				err := os.Mkdir(target, 0755)
				exit.ExitWithErr(err)
			}

			if !d.Type().IsRegular() {
				return nil
			}

			templ, err := template.ParseFiles(path)
			exit.ExitWithErr(err)

			stripped := strings.TrimPrefix(path, filepath.Join(uniDirectory, course, "template"))
			target := filepath.Join(uniDirectory, course, nextDir, stripped)

			file, err := os.Create(target)
			exit.ExitWithErr(err)

			data := templdata.New(&config, course, number)
			err = templ.Execute(file, data)
			exit.ExitWithErr(err)

			return nil
		})

		fmt.Printf("Success: Added working directory %s.\n", nextDir)
	},
}

func testNextDir(prefix string, course string) int {
	for i := 1; i <= 99; i++ {
		testDir := workingdir.FromNumber(i, prefix)
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
	exit.ExitWithErr(err)

	if !templateDirInfo.IsDir() {
		exit.ExitWithMsg("There is no template directory")
	}

	return templateDirPath
}
