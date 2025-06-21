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
	"github.com/rabuu/uni-cli/internal/util"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use: "retrieve",
	Aliases: []string{"r"},
	Short: "Import files into a course directory",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		courseId, number, cwd := dir.CwdWorkingDir(uniDirectory, &config)
		course := config.Courses[courseId]

		data := templating.Data(&config, courseId, number)

		for _, fileMap := range course.RetrieveFile {
			inFileName := templating.GenerateString(data, fileMap.From)
			inFilePath := util.EscapeHomeDir(inFileName)

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
				outFileName = templating.GenerateString(data, fileMap.To)
			}

			outFilePath := filepath.Join(cwd, outFileName)

			outFile, err := os.Create(outFilePath)
			exit.ExitWithErr(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			exit.ExitWithErr(err)

			fmt.Printf("Retrieved %s from %s.\n", outFileName, inFileName)

			if fileMap.Move {
				err := os.Remove(inFilePath)
				exit.ExitWithErr(err)
				fmt.Println("\t(and deleted source file)")
			}
		}

		for _, zipMap := range course.RetrieveZip {
			zipFileName := templating.GenerateString(data, zipMap.From)
			zipFilePath := util.EscapeHomeDir(zipFileName)

			zipFile, err := zip.OpenReader(zipFilePath)
			if os.IsNotExist(err) {
				fmt.Println("Not found:", zipFileName)
				continue
			}
			exit.ExitWithErr(err)
			defer zipFile.Close()

			for _, f := range zipFile.File {
				fpath := filepath.Join(zipMap.To, f.Name)

				if f.FileInfo().IsDir() {
					err = os.MkdirAll(fpath, os.ModePerm)
					exit.ExitWithErr(err)
					continue
				}

				err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
				exit.ExitWithErr(err)

				outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				exit.ExitWithErr(err)
				defer outFile.Close()

				rc, err := f.Open()
				exit.ExitWithErr(err)
				defer rc.Close()

				_, err = io.Copy(outFile, rc)
				exit.ExitWithErr(err)

				fmt.Printf("Retrieved %s from ZIP %s.\n", fpath, zipFileName)
			}

			if zipMap.Move {
				err := os.Remove(zipFilePath)
				exit.ExitWithErr(err)
				fmt.Println("\t(and deleted source file)")
			}
		}

		if len(course.RetrieveFile) + len(course.RetrieveZip) == 0 {
			fmt.Println("No retrieving rules specified.")
		}
	},
}
