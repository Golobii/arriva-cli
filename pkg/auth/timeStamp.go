package auth

import (
	"fmt"
	"time"
)

func FormatMonthDay(month int) string {
	if month < 10 {
		return fmt.Sprintf("0%d", month)
	}
	return fmt.Sprintf("%d", month)
}

func GetCurrentTimeStamp() string {
	date := time.Now()
	year, month, day := date.Date()
	hour := date.Hour()

	fDate := fmt.Sprintf("%d%s%s%d0000", year, FormatMonthDay(int(month)), FormatMonthDay(day), hour)

	return fDate
}
