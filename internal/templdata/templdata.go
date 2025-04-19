package templdata

import (
	"github.com/rabuu/uni-cli/internal/cfgfile"
	"github.com/rabuu/uni-cli/internal/exit"
	"github.com/rabuu/uni-cli/internal/workingdir"
)

type TemplateData struct {
	CourseId string
	Number int
	NumberPadded string
	Config cfgfile.Config
	Course cfgfile.Course
}

func New(config *cfgfile.Config, courseId string, number int) TemplateData {
	if !config.ContainsCourse(courseId) {
		exit.ExitWithMsg("No course:", courseId)
	}

	course := config.Courses[courseId]

	data := TemplateData{
		CourseId: courseId,
		Number: number,
		NumberPadded: workingdir.FromNumber(number, ""),
		Config: *config,
		Course: course,
	}

	return data
}

