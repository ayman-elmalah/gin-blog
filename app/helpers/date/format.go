package date

import (
	"fmt"
	"time"
)

func FormatAsDate(t time.Time) string {
	return fmt.Sprintf("%d/%02d/%02d", t.Year(), t.Month(), t.Day())
}

func FormatAsDateTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}
