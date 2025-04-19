package workingdir

import "fmt"

func FromNumber(i int, prefix string) string {
	return fmt.Sprintf("%s%02d", prefix, i)
}
