package internal

import (
	"fmt"
	"os"
)

func ExitWithErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func ExitWithMsg(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
