package commands

import (
	"fmt"
	ui "github.com/Golobii/arriva-cli/internal/ui"
	"github.com/Golobii/arriva-cli/pkg/api"
	"strconv"
	"strings"
	"time"
)

func HandleSearch(ds api.DepartureStations, timeStamp string, tok string, startStation string, endStation string, date string) int {

	station1 := getStationByName(startStation, ds)
	station2 := getStationByName(endStation, ds)
	if station1 == -1 {
		fmt.Printf("Station %s doesn't exist.\n", startStation)
		return -1
	} else if station2 == -1 {
		fmt.Printf("Station %s doesn't exist.\n", endStation)
		return -1
	}

	if len(date) < 5 {
		fmt.Println("Invalid date format. Try: arriva-cli help")
		return -1
	}

	res := api.FetchDepartures(timeStamp, tok, strconv.Itoa(station2),
		strconv.Itoa(station1), fmt.Sprintf("%d-%s", time.Now().Year(), date))

	if res[0].Error != "0" {
		fmt.Println("No routes found.")
		return -1
	}
	ui.PrettyPrintDepartures(res, startStation, endStation, date)
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
