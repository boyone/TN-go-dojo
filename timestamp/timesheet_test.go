package timestamp

import (
	"fmt"
	"testing"
	"time"
)

type H interface {
	Hour() int
}

type realTime struct {
}

func (r realTime) Hour(s string) int {
	return time.Now().Hour()
}

type fakeTime struct {
	hour int
}

func (f fakeTime) Hour() int {
	return f.hour
}

type Timesheet struct {
	H
}

func (t Timesheet) Record(s string) {
	result := "Late"
	if t.Now() <= 9 {
		result = "OnTime"
	}

	fmt.Println(s, " | ", result)
}

func (t Timesheet) Now() int {
	return t.Hour()
}

func TestOnTimeBefore9AM(t *testing.T) {
	timesheet := Timesheet{H: (fakeTime{8})}
	timesheet.Record("KWAN")
}

func TestNoomComeToOfficeAt10AMThatLateAfter9AM(t *testing.T) {
	timesheet := Timesheet{H: (fakeTime{10})}
	timesheet.Record("NOOM")
}

func TestPackComeToOfficeAt10AMThatLateAfter9AM(t *testing.T) {
	timesheet := Timesheet{H: (fakeTime{10})}
	timesheet.Record("PACK")
}

func TestOnTimeAt9AM(t *testing.T) {
	timesheet := Timesheet{H: fakeTime{hour: 9}}
	timesheet.Record("PHAK")
}
func TestOnTimeAt14AM(t *testing.T) {
	timesheet := Timesheet{H: fakeTime{hour: 14}}
	timesheet.Record("PACK")
}
