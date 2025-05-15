package util

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rabuu/uni-cli/internal/exit"
)

func EscapeHomeDir(path string) string {
	after, found := strings.CutPrefix(path, "~")
	if found {
		homedir, err := os.UserHomeDir()
		if err != nil {
			exit.ExitWithErr(err)
		}
		return filepath.Join(homedir, after)
	}

	return path
}
