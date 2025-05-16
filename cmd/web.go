package cmd

import (
	"os/exec"

	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use: "web",
	Short: "Open pre-configured web links",
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var arg string
		if len(args) > 0 {
			arg = args[0]
		}
		if arg == "" {
			arg = "default"
		}

		course := dir.CwdMaybeCourse(uniDirectory, &config)
		if course != "" {
			for name, link := range config.Courses[course].Web {
				if arg == name {
					cmd := exec.Command("xdg-open", "http://" + link)
					err := cmd.Run()
					exit.ExitWithErr(err)
					return
				}
			}
		}

		for name, link := range config.Web {
			if arg == name {
				cmd := exec.Command("xdg-open", "http://" + link)
				err := cmd.Run()
				exit.ExitWithErr(err)
				return
			}
		}

		exit.ExitWithMsg("No pre-configured web link with name:", arg)
	},
}
