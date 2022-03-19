package commands

import (
	"arriva-cli/pkg/api"
	"arriva-cli/pkg/auth"
	"fmt"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {

	timeStamp := auth.GetCurrentTimeStamp()
	tok := auth.GenerateToken(timeStamp)
	ds := api.FetchDepartureStations(timeStamp, tok)

	got := getStationByName("lJublJana aP", ds)
	if got == -1 {
		t.Error("Couldn't find the Ljubljana AP station")
	}

	date := time.Now()
	code := HandleSearch(ds, timeStamp, tok, "Kamnik metalka", "ljubljana aP", fmt.Sprintf("%s-%s", auth.FormatMonthDay(int(date.Month())), auth.FormatMonthDay(date.Day())))

	if code != 0 {
		t.Error("Search returned non 0 code.")
	}

}
