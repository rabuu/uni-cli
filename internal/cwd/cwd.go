package cwd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rabuu/uni-cli/internal/cfgfile"
	"github.com/rabuu/uni-cli/internal/exit"
)

func CourseDir(uniDirectory string, config *cfgfile.Config) string {
	cwd, err := os.Getwd()
	exit.ExitWithErr(err)

	if filepath.Dir(cwd) != uniDirectory {
		exit.ExitWithMsg("You're not in a course directory.")
	}

	course := filepath.Base(cwd)
	if !config.ContainsCourse(course) {
		exit.ExitWithMsg("You're not in a course directory.")
	}

	return course
}

func WorkingDir(uniDirectory string, config *cfgfile.Config) (course string, number int) {
	cwd, err := os.Getwd()
	exit.ExitWithErr(err)

	courseDir := filepath.Dir(cwd)

	if filepath.Dir(courseDir) != uniDirectory {
		exit.ExitWithMsg("You're not in a working directory.")
	}

	course = filepath.Base(courseDir)
	if !config.ContainsCourse(course) {
		exit.ExitWithMsg("You're not in a working directory.")
	}

	prefix := config.Courses[course].Prefix

	workingDir := filepath.Base(cwd)
	numberStr, found := strings.CutPrefix(workingDir, prefix)
	if !found {
		exit.ExitWithMsg("You're not in a working directory.")
	}

	number, err = strconv.Atoi(numberStr)
	if err != nil {
		exit.ExitWithMsg("You're not in a working directory.")
	}

	return
}
