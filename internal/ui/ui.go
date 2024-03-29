package console

import (
	"fmt"
	"github.com/Golobii/arriva-cli/pkg/api"
	"strings"
	"time"
)

func formatName(name string) string {
	return strings.Title(strings.ToLower(name))
}

func PrettyPrintDepartures(departures api.DeparturesResponse, startStation string, endStation string, date string) {
	fmt.Printf("%d-%s\n", time.Now().Year(), date)

	startStation = formatName(startStation)
	endStation = formatName(endStation)

	for _, departure := range departures[0].Departures {
		fmt.Printf("*- %s -- %s\n", startStation, departure.RodIodh)
		fmt.Println("│  ...")
		fmt.Printf("*- %s -- %s\n", endStation, departure.RodIpri)
		fmt.Printf("Trajanje: %s\n\n", formatTime(departure.RodCas))
	}
}

// Formats the length of a bus ride from int to str (7 => 00:07)
func formatTime(time int) string {
	hours := time / 60
	minutes := time - hours*60

	if hours < 10 {
		return fmt.Sprintf("0%d:%d", hours, minutes)
	}
	return fmt.Sprintf("%d:%d", hours, minutes)
}
