package timetable

import (
	"testing"
	"time"
)

func TestOntimeBefore9AM(t *testing.T) {
	record := ""
	timeReport := NewTimeReport(
		func() time.Time { return time.Date(2017, 11, 16, 7, 30, 0, 0, time.UTC) },
		func(str string) { record = str })

	timeReport("RockQA")

	if record != "RockQA : Ontime" {
		t.Error("Error")
	}
}

func TestOntime8AM(t *testing.T) {
	record := ""
	timeReport := NewTimeReport(
		func() time.Time { return time.Date(2017, 11, 16, 8, 00, 0, 0, time.UTC) },
		func(str string) { record = str })
	timeReport("Pack")

	if record != "Pack : Ontime" {
		t.Error("Error")
	}
}

func TestLateAfter9AM(t *testing.T) {
	record := ""
	timeReport := NewTimeReport(
		func() time.Time { return time.Date(2017, 11, 16, 9, 05, 0, 0, time.UTC) },
		func(str string) { record = str })

	timeReport("Nine")

	if record != "Nine : Late" {
		t.Error("Error")
	}

}
