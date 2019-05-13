package model

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Route struct {
	AirlineID   string
	Origin      string
	Destination string
}
type RouteMap map[string]map[string]Connection

const (
	routeFileName string = "routes"
)

func GetAllRoutes() ([]Route, error) {
	csvFile, err := os.Open("./data/" + routeFileName + ".csv")
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var routes []Route

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
			return routes, err
		}

		routes = append(routes, Route{
			AirlineID:   line[0],
			Origin:      line[1],
			Destination: line[2],
		})
	}

	return routes, nil
}

func GetRouteMap() (RouteMap, error) {
	csvFile, err := os.Open("./data/" + routeFileName + ".csv")
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	routeMap := make(RouteMap)

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

		_, originExists := routeMap[line[1]]
		if !originExists {
			routeMap[line[1]] = make(map[string]Connection)
		}
		_, destExists := routeMap[line[1]][line[2]]
		if destExists {
			continue
		}

		airline, err := GetAirlineByID(line[0])
		if err != nil {
			continue
		}
		originAirport, err := GetAirportByIATA(line[1])
		if err != nil {
			continue
		}
		destAirport, err := GetAirportByIATA(line[2])
		if err != nil {
			continue
		}

		routeMap[line[1]][line[2]] = Connection{
			Airline:     airline,
			Origin:      originAirport,
			Destination: destAirport,
		}
	}

	return routeMap, nil
}
