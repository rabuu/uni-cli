package dir

import "fmt"

func FormatWorkdirName(i int, prefix string) string {
	return fmt.Sprintf("%s%02d", prefix, i)
}
