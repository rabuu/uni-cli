package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/templating"
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

		courseId, number, cwd := dir.CwdWorkingDir(uniDirectory, &config)
		course := config.Courses[courseId]

		data := templating.Data(&config, courseId, number)

		for _, fileMap := range course.ExportFile {
			inFileName := templating.GenerateString(data, fileMap.From)
			inFilePath := filepath.Join(cwd, inFileName)

			inFile, err := os.Open(inFilePath)
			if os.IsNotExist(err) {
				fmt.Println("Not found:", inFileName)
				continue
			}
			exit.ExitWithErr(err)
			defer inFile.Close()

			outFileName := templating.GenerateString(data, fileMap.To)
			outFilePath := filepath.Join(exportDirectory, outFileName)

			outFile, err := os.Create(outFilePath)
			exit.ExitWithErr(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			exit.ExitWithErr(err)

			fmt.Printf("Exported %s to %s.\n", inFileName, outFileName)
		}

		if len(course.ExportFile) == 0 {
			fmt.Println("No export rules specified.")
		}
	},
}

func init() {
	exportCmd.Flags().BoolVarP(&cleanFlag, "clean", "c", false, "clean the export directory")
}
