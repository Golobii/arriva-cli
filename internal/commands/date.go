package commands

import (
	"arriva-cli/pkg/auth"
	"fmt"
	"time"
)

func DisplayDate() {
	date := time.Now()

	fmt.Printf("%s-%s\n", auth.FormatMonthDay(int(date.Month())), auth.FormatMonthDay(date.Day()))
}
