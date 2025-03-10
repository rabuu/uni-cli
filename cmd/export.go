package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/rabuu/uni-cli/internal/cwd"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/templdata"
	"github.com/spf13/cobra"
)

var cleanFlag bool
var exportCmd = &cobra.Command{
	Use: "export",
	Aliases: []string{"x"},
	Short: "Export output from working directory",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if cleanFlag {
			err := os.RemoveAll(exportDirectory)
			exit.ExitWithErr(err)
			err = os.Mkdir(exportDirectory, 0755)
			exit.ExitWithErr(err)
			return
		}

		courseName, number := cwd.WorkingDir(uniDirectory, &config)
		course := config.Courses[courseName]

		if len(course.Export) == 0 {
			fmt.Println("No export rules specified.")
		}
		
		for _, export := range course.Export {
			fileInfo, err := os.Stat(export.Filename)
			
			if os.IsNotExist(err) {
				fmt.Println("Not found:", export.Filename)
				continue
			}
			exit.ExitWithErr(err)

			if !fileInfo.Mode().Type().IsRegular() {
				exit.ExitWithMsg("Error: Not a regular file:", export.Filename)
			}

			outTempl := template.Must(template.New("output file").Parse(export.Output))

			var outFileBuff bytes.Buffer
			data := templdata.New(&config, courseName, number)
			err = outTempl.Execute(&outFileBuff, data)
			exit.ExitWithErr(err)
			outFile := outFileBuff.String()
			outFilePath := filepath.Join(exportDirectory, outFile)

			cpCmd := exec.Command("cp", export.Filename, outFilePath)
			err = cpCmd.Run()
			exit.ExitWithErr(err)

			fmt.Printf("Exported %s to %s.\n", export.Filename, outFile)
		}
	},
}

func init() {
	exportCmd.Flags().BoolVarP(&cleanFlag, "clean", "c", false, "clean the export directory")
}
