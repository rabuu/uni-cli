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
	Members []cfgfile.GroupMember
	Semester string
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
		Members: config.Courses[courseName].Members,
		Semester: config.Semester,
	}

	return data
}

