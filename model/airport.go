package model

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
)

type Airport struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	IATA3     string `json:"iata"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

const (
	airportFileName string = "airports"
)

var airports map[string]Airport
var airportDataIsLoaded = false

func loadAirports() error {
	csvFile, err := os.Open("./data/" + airportFileName + ".csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	airports = make(map[string]Airport)
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

		if line[3] == "\\N" {
			continue
		}

		airports[line[3]] = Airport{
			Name:      line[0],
			City:      line[1],
			Country:   line[2],
			IATA3:     line[3],
			Latitude:  line[4],
			Longitude: line[5],
		}
	}

	airportDataIsLoaded = true
	return nil
}

func GetAirportsByRoute(originCity, originCountry, destinationCity, destinationCountry string) ([]string, []string, error) {
	if !airportDataIsLoaded {
		if err := loadAirports(); err != nil {
			return nil, nil, err
		}
	}

	var origins []string
	var destinations []string

	for _, airport := range airports {
		if airport.City == originCity && airport.Country == originCountry {
			origins = append(origins, airport.IATA3)
		} else if airport.City == destinationCity && airport.Country == destinationCountry {
			destinations = append(destinations, airport.IATA3)
		}
	}

	return origins, destinations, nil
}

func GetAirportByIATA(iata string) (*Airport, error) {
	if !airportDataIsLoaded {
		if err := loadAirports(); err != nil {
			return nil, err
		}
	}

	airport, ok := airports[iata]
	if !ok {
		return nil, errors.New("Airport not found")
	}
	return &airport, nil
}
