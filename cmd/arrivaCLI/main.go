package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	ui "arriva-cli/internal/ui"
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
	case "help":
		fmt.Println("Not implemented yet.")
		return
	default:
		handleSearch(ds, timeStamp, tok)
		break
	}

}

func handleSearch(ds api.DepartureStations, timeStamp string, tok string) {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments supplied! Try: arrivaCLI help")
		os.Exit(1)
	}
	res := api.FetchDepartures(timeStamp, tok, strconv.Itoa(getStationByName(os.Args[1], ds)),
		strconv.Itoa(getStationByName(os.Args[2], ds)), fmt.Sprintf("%d-%s", time.Now().Year(), os.Args[3]))

	if res[0].Error != "0" {
		fmt.Println("No routes found.")
		os.Exit(0)
	}
	ui.PrettyPrintDepartures(res, os.Args[1], os.Args[2])

}

func getStationByName(name string, ds api.DepartureStations) int {

	name = strings.ToLower(name)
	for _, station := range ds[0].DepartureStations {
		if strings.ToLower(station.PosNaz) == name {
			return station.JposIjpp
		}
	}

	return -1
}
