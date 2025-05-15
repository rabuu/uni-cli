package templating

import (
	"strings"
	"time"

	"github.com/rabuu/uni-cli/internal/exit"
)

type DateUtil struct {
	format string
}

func (util DateUtil) Today() (date string) {
	now := time.Now()
	date = now.Format(util.format)
	return
}

func (util DateUtil) NextWeekday(day string, weekOffset int) (date string) {
	now := time.Now()
	weekday := weekdayFromString(day)

	daysUntilWeekday := (int(weekday) - int(now.Weekday()) + 7) % 7

	if daysUntilWeekday == 0 {
		daysUntilWeekday = 7
	}

	t := now.AddDate(0, 0, daysUntilWeekday + (7 * weekOffset))
	date = t.Format(util.format)
	return
}

func weekdayFromString(day string) (weekday time.Weekday) {
	switch strings.ToLower(day) {
	case "monday": weekday = time.Monday
	case "tuesday": weekday = time.Tuesday
	case "wednesday": weekday = time.Wednesday
	case "thursday": weekday = time.Thursday
	case "friday": weekday = time.Friday
	case "saturday": weekday = time.Saturday
	case "sunday": weekday = time.Sunday
	default: exit.ExitWithMsg("No valid weekday: " + day)
	}
	return
}
