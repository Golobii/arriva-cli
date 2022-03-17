package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"arriva-cli/pkg/api"
	"arriva-cli/pkg/auth"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("arrivaCLI - arriva CLI Tool")
		return
	}

	timeStamp := auth.GetCurrentTimeStamp()
	tok := auth.GenerateToken(timeStamp)
	ds := api.FetchDepartureStations(timeStamp, tok)

	switch os.Args[1] {
	case "list":
		for _, station := range ds[0].DepartureStations {
			fmt.Println(station.PosNaz)
		}
		return
	case "date":
		fmt.Println("Not implemented yet.")
		return
	default:
		res := api.FetchDepartures(timeStamp, tok, strconv.Itoa(GetStationByName(os.Args[1], ds)),
			strconv.Itoa(GetStationByName(os.Args[2], ds)), fmt.Sprintf("%d-%s", time.Now().Year(), os.Args[3]))
		PrettyPrintDepartures(res)
		break
	}

}

// TODO: clean this file

func PrettyPrintDepartures(departures api.DeparturesResponse) {
	fmt.Println("2022-03-16")
	for _, departure := range departures[0].Departures {
		fmt.Printf("*- %s -- %s\n", "Mengeš Lovec", departure.RodIodh) // TODO: add right names
		fmt.Println("│  ...")
		fmt.Printf("*- %s -- %s\n", "Kamnik Metalka", departure.RodIpri)
		fmt.Printf("Trajanje: %s\n\n", formatTime(departure.RodCas))
	}
}

func GetStationByName(name string, ds api.DepartureStations) int {

	for _, station := range ds[0].DepartureStations {
		if station.PosNaz == name {
			return station.JposIjpp
		}
	}

	return -1
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
