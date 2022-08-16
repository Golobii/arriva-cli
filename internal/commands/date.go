package commands

import (
	"fmt"
	"github.com/Golobii/arriva-cli/pkg/auth"
	"time"
)

func DisplayDate() {
	date := time.Now()

	fmt.Printf("%s-%s\n", auth.FormatMonthDay(int(date.Month())), auth.FormatMonthDay(date.Day()))
}
