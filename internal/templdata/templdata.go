package templdata

import (
	"fmt"

	"github.com/rabuu/uni-cli/internal/cfgfile"
	"github.com/rabuu/uni-cli/internal/exit"
)

type TemplateData struct {
	Course, CourseName string
	Number int
	NumberPadded string
	Group []cfgfile.GroupMember
}

func New(config *cfgfile.Config, courseName string, number int) TemplateData {
	if !config.ContainsCourse(courseName) {
		exit.ExitWithMsg("No course:", courseName)
	}

	data := TemplateData{
		Course: courseName,
		CourseName: config.Courses[courseName].FullName,
		Number: number,
		NumberPadded: fmt.Sprintf("%02d", number),
		Group: config.Courses[courseName].Group,
	}

	return data
}

