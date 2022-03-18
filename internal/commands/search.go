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

func HandleSearch(ds api.DepartureStations, timeStamp string, tok string) {
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
	ui.PrettyPrintDepartures(res, os.Args[1], os.Args[2], os.Args[3])

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
