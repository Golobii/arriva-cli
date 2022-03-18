package commands

import (
	ui "arriva-cli/internal/ui"
	"arriva-cli/pkg/api"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleSearch(ds api.DepartureStations, timeStamp string, tok string) int {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments supplied! Try: arriva-cli help")
		return -1
	}

	station1 := getStationByName(os.Args[1], ds)
	station2 := getStationByName(os.Args[2], ds)
	if station1 == -1 {
		fmt.Printf("Station %s doesn't exist.\n", os.Args[1])
		return -1
	} else if station2 == -1 {
		fmt.Printf("Station %s doesn't exist.\n", os.Args[2])
		return -1
	}

	if len(os.Args[3]) < 5 {
		fmt.Println("Invalid date format. Try: arriva-cli help")
		return -1
	}

	res := api.FetchDepartures(timeStamp, tok, strconv.Itoa(station2),
		strconv.Itoa(station1), fmt.Sprintf("%d-%s", time.Now().Year(), os.Args[3]))

	if res[0].Error != "0" {
		fmt.Println("No routes found.")
		return -1
	}
	ui.PrettyPrintDepartures(res, os.Args[1], os.Args[2], os.Args[3])
	return 0
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
