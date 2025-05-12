package cmd

import (
	"archive/zip"
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

		for _, zipCfg := range course.ExportZip {
			archiveFileName := templating.GenerateString(data, zipCfg.ArchiveFile)
			archiveFilePath := filepath.Join(exportDirectory, archiveFileName + ".zip")

			archive, err := os.Create(archiveFilePath)
			exit.ExitWithErr(err)
			defer archive.Close()
			fmt.Printf("Export ZIP archive %s...\n", archiveFileName)

			zipper := zip.NewWriter(archive)

			for _, fileMap := range zipCfg.Include {
				inFileName := templating.GenerateString(data, fileMap.From)
				inFilePath := filepath.Join(cwd, inFileName)

				inFile, err := os.Open(inFilePath)
				if os.IsNotExist(err) {
					fmt.Println("\tNot found:", inFileName)
					continue
				}
				exit.ExitWithErr(err)
				defer inFile.Close()

				outFileName := templating.GenerateString(data, fileMap.To)
				outFilePath := filepath.Join(archiveFileName, outFileName)

				outFile, err := zipper.Create(outFilePath)
				exit.ExitWithErr(err)

				_, err = io.Copy(outFile, inFile)
				exit.ExitWithErr(err)

				fmt.Printf("\tZipped %s to %s.\n", inFileName, outFilePath)
			}

			zipper.Close()
			fmt.Println("Closed ZIP archive.")
		}

		if len(course.ExportFile) + len(course.ExportZip) == 0 {
			fmt.Println("No export rules specified.")
		}
	},
}

func init() {
	exportCmd.Flags().BoolVarP(&cleanFlag, "clean", "c", false, "clean the export directory")
}
