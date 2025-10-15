package dir

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/rabuu/uni-cli/internal/exit"
)

func getCwd(followSymlinks bool) (cwd string) {
	cwd, err := os.Getwd()
	exit.ExitWithErr(err)

	if followSymlinks {
		cwd, err = filepath.EvalSymlinks(cwd)
		exit.ExitWithErr(err)
	}

	return
}

func CwdCourseDir(uniDirectory string, config *cfg.Config) (course string, cwd string) {
	cwd = getCwd(config.FollowSymlinks)

	if filepath.Dir(cwd) != uniDirectory {
		exit.ExitWithMsg("You're not in a course directory.")
	}

	course = filepath.Base(cwd)
	if !config.ContainsCourse(course) {
		exit.ExitWithMsg("You're not in a course directory.")
	}

	return
}

func CwdWorkingDir(uniDirectory string, config *cfg.Config) (course string, number int, cwd string) {
	cwd = getCwd(config.FollowSymlinks)

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

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		exit.ExitWithMsg("You're not in a working directory.")
	}

	return
}

func CwdMaybeCourse(uniDirectory string, config *cfg.Config) (course string) {
	cwd := getCwd(config.FollowSymlinks)

	rest := cwd
	for rest != "/" {
		dir := filepath.Dir(rest)

		if dir == uniDirectory {
			maybeCourse := filepath.Base(rest)
			if config.ContainsCourse(maybeCourse) {
				course = maybeCourse
			}

			return
		}

		rest = dir
	}

	return
}
