package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/templating"
	"github.com/rabuu/uni-cli/internal/util"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use: "retrieve",
	Aliases: []string{"r"},
	Short: "Retrieve material from other locations",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		courseId, number, cwd := dir.CwdWorkingDir(uniDirectory, &config)
		course := config.Courses[courseId]

		data := templating.Data(&config, courseId, number)

		for _, fileMap := range course.RetrieveFile {
			inFileName := templating.GenerateString(data, fileMap.From)
			inFilePath := util.EscapeHomeDir(inFileName)

			fmt.Println(inFilePath)

			inFile, err := os.Open(inFilePath)
			if os.IsNotExist(err) {
				fmt.Println("Not found:", inFileName)
				continue
			}
			exit.ExitWithErr(err)
			defer inFile.Close()

			var outFileName string
			if fileMap.To == "" {
				outFileName = filepath.Base(inFilePath)
			} else {
				templating.GenerateString(data, fileMap.To)
			}

			outFilePath := filepath.Join(cwd, outFileName)

			outFile, err := os.Create(outFilePath)
			exit.ExitWithErr(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			exit.ExitWithErr(err)

			fmt.Printf("Retrieved %s from %s.\n", outFileName, inFileName)
		}

		if len(course.RetrieveFile) == 0 {
			fmt.Println("No retrieving rules specified.")
		}
	},
}
