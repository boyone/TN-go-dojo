package timetable

import (
	"time"
	"fmt"
)

func NewTimeReport(now func() time.Time, print func(str string)) func(staff string) {
	return func(staff string) {
		record := "Late"
		time := now()

		if time.Hour() < 8 {
			record = "Ontime"
		}

		if time.Hour() == 8 && time.Minute() == 0 {
			record = "Ontime"
		}

		s := fmt.Sprintf("%s : %s", staff, record)
		print(s)
	}
}
