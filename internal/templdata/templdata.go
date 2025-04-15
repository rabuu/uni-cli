package templdata

import (
	"fmt"

	"github.com/rabuu/uni-cli/internal/cfgfile"
	"github.com/rabuu/uni-cli/internal/exit"
)

type TemplateData struct {
	CourseId string
	Number int
	NumberPadded string
	Config cfgfile.Config
	Course cfgfile.Course
}

func New(config *cfgfile.Config, courseName string, number int) TemplateData {
	if !config.ContainsCourse(courseName) {
		exit.ExitWithMsg("No course:", courseName)
	}

	course := config.Courses[courseName]

	data := TemplateData{
		CourseId: courseName,
		Number: number,
		NumberPadded: fmt.Sprintf("%02d", number),
		Config: *config,
		Course: course,
	}

	return data
}

