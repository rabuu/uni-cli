package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/templating"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use: "next",
	Aliases: []string{"n"},
	Short: "Generate next working directory",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		course, cwd := dir.CwdCourseDir(uniDirectory, &config)

		prefix := config.Courses[course].Prefix
		number := testNextDir(cwd, prefix)
		nextDir := dir.FormatWorkdirName(number, prefix)

		err := os.Mkdir(nextDir, 0755)
		exit.ExitWithErr(err)

		templateDirPath := getTemplateDirPath(course)
		if templateDirPath != "" {
			filepath.WalkDir(templateDirPath, func(path string, d fs.DirEntry, err error) error {
				return applyTemplate(path, d, err, cwd, nextDir, course, number)
			})
		}

		fmt.Printf("Success: Added working directory %s.\n", nextDir)
	},
}

func testNextDir(cwd string, prefix string) int {
	for i := 1; i <= 99; i++ {
		testDir := dir.FormatWorkdirName(i, prefix)
		testDirPath := filepath.Join(cwd, testDir)

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
	if os.IsNotExist(err) {
		return ""
	}
	exit.ExitWithErr(err)

	if !templateDirInfo.IsDir() {
		exit.ExitWithMsg("Bad template directory")
	}

	return templateDirPath
}

func applyTemplate(path string, d fs.DirEntry, err error, cwd string, nextDir string, course string, number int) error {
	exit.ExitWithErr(err)

	if d.IsDir() {
		stripped := strings.TrimPrefix(path, filepath.Join(cwd, "template"))
		if stripped == "" {
			return nil
		}

		target := filepath.Join(cwd, nextDir, stripped)
		err := os.Mkdir(target, 0755)
		exit.ExitWithErr(err)
	}

	if !d.Type().IsRegular() {
		return nil
	}

	templ, err := template.ParseFiles(path)
	exit.ExitWithErr(err)

	stripped := strings.TrimPrefix(path, filepath.Join(cwd, "template"))
	target := filepath.Join(cwd, nextDir, stripped)

	file, err := os.Create(target)
	exit.ExitWithErr(err)

	data := templating.Data(&config, course, number)
	err = templ.Execute(file, data)
	exit.ExitWithErr(err)

	return nil
}
