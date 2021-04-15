package clock

import "time"

const (
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"
)

func FormatDate(when time.Time) string {
	return when.Format(DateLayout)
}

func FormatDateTime(when time.Time) string {
	return when.Format(DateTimeLayout)
}
