package model

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Airline struct {
	TwoDigitCode   string `json:"-"`
	ThreeDigitCode string `json:"code"`
	Name           string `json:"name"`
	Country        string `json:"country"`
}

const (
	airlineFileName string = "airlines"
)

func GetAirlineByID(code string) (*Airline, error) {
	csvFile, err := os.Open("./data/" + routeFileName + ".csv")
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	isHeader := true
	for {
		line, err := reader.Read()
		if isHeader {
			isHeader = false
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if line[2] != code {
			continue
		}

		return &Airline{
			Name:           line[0],
			TwoDigitCode:   line[1],
			ThreeDigitCode: line[2],
			Country:        line[3],
		}, nil
	}

	return nil, nil
}

/*
Name					2 Digit Code	3 Digit Code	Country
Air China				CA				CCA				China
China Southern Airlines	CZ				CSN				China
Southwest Airlines		WN				SWA				United States
Turkish Airlines		TK				THY				Turkey
United Airlines			UA				UAL				United States
WestJet					WS				WJA				Canada
*/
