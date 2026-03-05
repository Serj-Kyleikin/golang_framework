package utils

import (
	"fmt"
	"time"
)

func ParseMonthUTC(s string) (time.Time, error) {
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date, expected MM-YYYY")
	}
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC), nil
}

func FormatMonthUTC(t time.Time) string {
	t = t.UTC()
	return fmt.Sprintf("%02d-%04d", int(t.Month()), t.Year())
}
