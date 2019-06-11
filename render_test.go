package main

import (
	"testing"
	"time"
)

var (
	checkFridayTestData = []struct {
		weekDay  time.Weekday
		isFriday bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, true},
		{6, false},
	}
)

func TestCheckFriday(t *testing.T) {
	for _, test := range checkFridayTestData {
		if checkFriday(test.weekDay) != test.isFriday {
			t.Error("checkFriday function returned incorrect result")
		}
	}
}
