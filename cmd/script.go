package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use: "script",
	Short: "Run global or course-local scripts",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		course := dir.CwdMaybeCourse(uniDirectory, &config)
		path := filepath.Join(uniDirectory, course, "scripts", arg)
		valid := validScript(path)

		if !valid && course != "" {
			path = filepath.Join(uniDirectory, "scripts", arg)
			valid = validScript(path)
		}

		if !valid {
			exit.ExitWithMsg("No valid script with name:", arg)
		}

		c := exec.Command(path)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()
		exit.ExitWithErr(err)
	},
}

func validScript(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}
