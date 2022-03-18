package main

import (
	"fmt"
	"os"

	"arriva-cli/internal/commands"
	"arriva-cli/pkg/api"
	"arriva-cli/pkg/auth"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("arrivaCLI - arriva CLI Tool")
		return
	} else if os.Args[1] == "help" {
		commands.PrintHelp()
		return
	} else if os.Args[1] == "date" {
		commands.DisplayDate()
		return
	}

	timeStamp := auth.GetCurrentTimeStamp()
	tok := auth.GenerateToken(timeStamp)
	ds := api.FetchDepartureStations(timeStamp, tok)

	if os.Args[1] == "list" {
		for _, station := range ds[0].DepartureStations {
			fmt.Println(station.PosNaz)
		}
		return
	} else {
		commands.HandleSearch(ds, timeStamp, tok)
		return
	}

}
