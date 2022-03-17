package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUri = "https://prometws.alpetour.si/"

type DeparturesResponse []struct {
	Departures []struct {
		SpodSif int     `json:"SPOD_SIF"`
		RegIsif string  `json:"REG_ISIF"`
		RprSif  string  `json:"RPR_SIF"`
		RprNaz  string  `json:"RPR_NAZ"`
		OvrSif  string  `json:"OVR_SIF"`
		RodIodh string  `json:"ROD_IODH"`
		RodIpri string  `json:"ROD_IPRI"`
		RodCas  int     `json:"ROD_CAS"`
		RodPer  string  `json:"ROD_PER"`
		RodKm   int     `json:"ROD_KM"`
		RodOpo  string  `json:"ROD_OPO"`
		VzclCen float64 `json:"VZCL_CEN"`
		VvlnZl  int     `json:"VVLN_ZL"`
		RodZapz int     `json:"ROD_ZAPZ"`
		RodZapk int     `json:"ROD_ZAPK"`
	} `json:"Departures"`
	Error    string `json:"Error"`
	ErrorMsg string `json:"ErrorMsg"`
}

type DepartureStations []struct {
	DepartureStations []struct {
		JposIjpp int    `json:"JPOS_IJPP"`
		PosNaz   string `json:"POS_NAZ"`
	} `json:"DepartureStations"`
	Error    string `json:"Error"`
	ErrorMsg string `json:"ErrorMsg"`
}

func addAuthParams(timeStamp string, token string) string {
	return fmt.Sprintf("cTimeStamp=%s&cToken=%s", timeStamp, token)
}

func FetchDepartureStations(timeStamp string, token string) DepartureStations {
	url := fmt.Sprintf("%sWS_ArrivaSLO_TimeTable_DepartureStations.aspx?&%s&json=1", baseUri, addAuthParams(timeStamp, token))

	body := fetch(url)

	var parsedJson DepartureStations
	if err := json.Unmarshal(body, &parsedJson); err != nil {
		fmt.Println(err)
	}
	return parsedJson

}

func FetchDepartures(timeStamp string, token string, IJPPZ string, IJPPK string, date string) DeparturesResponse {
	url := fmt.Sprintf("%sWS_ArrivaSLO_TimeTable_TimeTableDepartures.aspx?&%s&JPOS_IJPPZ=%s&JPOS_IJPPK=%s&VZVK_DAT=%s&json=1",
		baseUri, addAuthParams(timeStamp, token), IJPPZ, IJPPK, date)

	body := fetch(url)

	var parsedJson DeparturesResponse
	if err := json.Unmarshal(body, &parsedJson); err != nil {
		fmt.Println("Can't unmarshal JSON")
	}

	return parsedJson
}

func fetch(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body
}
