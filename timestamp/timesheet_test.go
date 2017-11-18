package timestamp

import (
	"testing"
	"fmt"
	"time"
)

type H interface {
	Hour() int
}

type realTime struct {
}

func (r realTime) Hour() int {
	return time.Now().Hour()
}

type fakeTime struct {
}

func (f fakeTime) Hour() int {
	return 8
}

type Timesheet struct {
	H
}

func (t Timesheet) Record(s string) {
	result := "Late"
	if t.Now() < 9 {
		result = "OnTime"
	}

	fmt.Println(s, " | ", result)
}

func (t Timesheet) Now() int {
	return t.Hour()
}

func TestOnTimeBefore9AM(t *testing.T) {
	timesheet := Timesheet{H: new(fakeTime)}
	//timesheet := Timesheet{H: new(realTime)}

	timesheet.Record("KWAN")
}
