package model

import (
	"bufio"
	"encoding/csv"
	"errors"
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

var airlines map[string]Airline
var airlineDataIsLoaded = false

func loadAirlines() error {
	csvFile, err := os.Open("./data/" + airlineFileName + ".csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	airlines = make(map[string]Airline)
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
			return err
		}

		airlines[line[1]] = Airline{
			Name:           line[0],
			TwoDigitCode:   line[1],
			ThreeDigitCode: line[2],
			Country:        line[3],
		}
	}

	airlineDataIsLoaded = true
	return nil
}

func GetAirlineByID(id string) (*Airline, error) {
	if !airlineDataIsLoaded {
		if err := loadAirlines(); err != nil {
			return nil, err
		}
	}

	airline, ok := airlines[id]
	if !ok {
		return nil, errors.New("Airline not found")
	}
	return &airline, nil
}
