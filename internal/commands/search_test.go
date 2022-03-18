package commands

import (
	"arriva-cli/pkg/api"
	"arriva-cli/pkg/auth"
	"testing"
)

func TestSearch(t *testing.T) {

	timeStamp := auth.GetCurrentTimeStamp()
	tok := auth.GenerateToken(timeStamp)
	ds := api.FetchDepartureStations(timeStamp, tok)

	got := getStationByName("lJublJana aP", ds)
	if got == -1 {
		t.Error("Couldn't find the Ljubljana AP station")
	}
}
