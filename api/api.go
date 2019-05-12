package api

import (
	"fmt"
	"net/http"
	"strings"

	"back-end-take-home/model"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.RouterGroup) {
	search := router.Group("/flight")
	{
		search.GET("", getBestFlight)
	}
}

func getBestFlight(c *gin.Context) {
	origin := strings.Replace(c.Query("origin"), " ", "", -1)
	if origin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'origin' param not specified"})
		return
	}

	destination := strings.Replace(c.Query("destination"), " ", "", -1)
	if destination == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'destination' param not specified"})
		return
	}

	if origin == destination {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no route because 'origin' and 'destination' are the same"})
		return
	}

	originArr := strings.Split(origin, ",")
	if len(originArr) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected format for 'origin' param"})
		return
	}

	destArr := strings.Split(destination, ",")
	if len(destArr) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected format for 'destination' param"})
		return
	}

	originAirports, destAirports, err := model.GetAirportsByRoute(originArr[0], originArr[1], destArr[0], destArr[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting origin and destination airports: " + err.Error()})
		return
	}
	if len(originAirports) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no airport for the specified 'origin'"})
		return
	}
	if len(destAirports) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no airport for the specified 'destination'"})
		return
	}

	routes, err := model.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting all routes: " + err.Error()})
		return
	}

	flight := calculateBestFlight(originAirports, destAirports, routes)
	if flight == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no route for this 'origin' and 'destination'"})
		return
	}
	c.JSON(http.StatusOK, flight)
}

func calculateBestFlight(originAirports, destAirports []string, routes []model.Route) []model.Connection {
	fmt.Println("calculateBestFlight")
	var bestFlight []model.Connection
	for _, origin := range originAirports {
		for _, dest := range destAirports {
			var previousOrigins []string
			flight := findBestFlightForRoute(origin, dest, previousOrigins, routes, 0)

			if flight == nil || len(flight) <= 0 {
				continue
			}

			fmt.Printf("calculateBestFlight loop result: %+v\n", flight)

			if len(bestFlight) == 0 || len(bestFlight) > len(flight) {
				bestFlight = flight
				if len(bestFlight) == 1 {
					break
				}
			}
		}
	}

	return bestFlight
}

func findBestFlightForRoute(origin string, finalDestination string, previousOrigins []string, routes []model.Route, bestFlightLength int) []model.Connection {
	fmt.Printf("findBestFlightForRoute %s %s %v \n", origin, finalDestination, previousOrigins)
	if bestFlightLength > 0 && bestFlightLength <= len(previousOrigins) {
		return nil
	}

	for _, route := range routes {
		if route.Origin == origin {
			if !destinationIsValid(previousOrigins, route.Destination) {
				continue
			}

			if route.Destination == finalDestination {
				connection, err := generateConnection(route)
				if err != nil {
					return nil
				}
				return []model.Connection{*connection}
			} else {
				flight := findBestFlightForRoute(route.Destination, finalDestination, append(previousOrigins, route.Origin), routes, bestFlightLength)
				if flight == nil {
					continue
				}
				connection, err := generateConnection(route)
				if err != nil {
					return nil
				}
				flight = append([]model.Connection{*connection}, flight...)
				return flight
			}
		}
	}

	return nil
}

func destinationIsValid(origins []string, destination string) bool {
	for _, origin := range origins {
		if origin == destination {
			return false
		}
	}
	return true
}

func generateConnection(route model.Route) (*model.Connection, error) {
	// airline, err := model.GetAirlineByID(route.AirlineID)
	// if err != nil {
	// 	return nil, err
	// }

	// originAirport, err := model.GetAirportByIATA(route.Origin)
	// if err != nil {
	// 	return nil, err
	// }

	// destAirport, err := model.GetAirportByIATA(route.Destination)
	// if err != nil {
	// 	return nil, err
	// }

	// return &model.Connection{
	// 	Airline:     airline,
	// 	Origin:      originAirport,
	// 	Destination: destAirport,
	// }, nil
	return &model.Connection{
		Airline:     &model.Airline{ThreeDigitCode: route.AirlineID},
		Origin:      &model.Airport{IATA3: route.Origin},
		Destination: &model.Airport{IATA3: route.Destination},
	}, nil
}

func generateConnections(routes []model.Route) ([]model.Connection, error) {
	var connections []model.Connection

	for _, route := range routes {
		airline, err := model.GetAirlineByID(route.AirlineID)
		if err != nil {
			return nil, err
		}

		originAirport, err := model.GetAirportByIATA(route.Origin)
		if err != nil {
			return nil, err
		}

		destAirport, err := model.GetAirportByIATA(route.Destination)
		if err != nil {
			return nil, err
		}

		connections = append(connections, model.Connection{
			Airline:     airline,
			Origin:      originAirport,
			Destination: destAirport,
		})
	}

	return connections, nil
}
