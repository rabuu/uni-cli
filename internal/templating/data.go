package templating

import (
	"github.com/rabuu/uni-cli/internal/cfg"
	"github.com/rabuu/uni-cli/internal/dir"
	"github.com/rabuu/uni-cli/internal/exit"
)

type TemplateData struct {
	CourseId string
	Number int
	NumberPadded string
	Config cfg.Config
	Course cfg.Course
}

func Data(config *cfg.Config, courseId string, number int) TemplateData {
	if !config.ContainsCourse(courseId) {
		exit.ExitWithMsg("No course:", courseId)
	}

	course := config.Courses[courseId]

	data := TemplateData{
		CourseId: courseId,
		Number: number,
		NumberPadded: dir.FormatWorkdirName(number, ""),
		Config: *config,
		Course: course,
	}

	return data
}

