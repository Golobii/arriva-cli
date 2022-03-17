package main

import (
	"fmt"
	"time"
)

func formatMonth(month time.Month) string {
	if month < 10 {
		return fmt.Sprintf("0%d", month)
	}
    return fmt.Sprintf("%d", month)
}

func GetCurrentTimeStamp() string {
	date := time.Now()
	year, month, day := date.Date()
	hour := date.Hour()

	fDate := fmt.Sprintf("%d%s%d%d0000", year, formatMonth(month), day, hour)

	return fDate
}