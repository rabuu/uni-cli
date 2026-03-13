package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/spf13/cobra"
)

var exclude = []string { ".direnv", ".envrc" }

var archiveCmd = &cobra.Command{
	Use: "archive",
	Short: "Archive the current state of the uni directory",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		outpath := args[0]

		err := archive(uniDirectory, outpath, exclude)
		exit.ExitWithErr(err)

		fmt.Println("Successfully archived the uni directory to", outpath)
	},
}

func archive(src, dest string, exclude []string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	gw := gzip.NewWriter(out)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	toplevel := filepath.Base(src)

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if rel == "." {
			return nil
		}

		// exclusion check
		if slices.Contains(exclude, info.Name()) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = filepath.Join(toplevel, rel)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(tw, f)
		return err
	})
}
