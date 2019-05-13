package api

import (
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
	origin := c.Query("origin")
	if origin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'origin' param not specified"})
		return
	}

	destination := c.Query("destination")
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

	routeMap, err := model.GetRouteMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting all routes: " + err.Error()})
		return
	}
	flight := calculateBestFlight(originAirports, destAirports, routeMap)
	if flight == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no route for this 'origin' and 'destination'"})
		return
	}
	c.JSON(http.StatusOK, flight)
}

func calculateBestFlight(originAirports, destAirports []string, routeMap model.RouteMap) []model.Connection {
	var bestFlight []model.Connection

	for _, origin := range originAirports {
		for _, dest := range destAirports {
			flight := findBestFlightForRoute(origin, dest, make(map[string]bool), routeMap)

			if flight == nil || len(flight) <= 0 {
				continue
			}

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

func findBestFlightForRoute(origin string, finalDestination string, previousOrigins map[string]bool, routeMap model.RouteMap) []model.Connection {
	routeMapDest, exists := routeMap[origin]
	if !exists {
		return nil
	}

	connection, exists := routeMapDest[finalDestination]
	if exists {
		return []model.Connection{connection}
	}

	var bestFlightForRoute []model.Connection

	for destination, connection := range routeMapDest {
		_, previousExists := previousOrigins[destination]
		if previousExists {
			continue
		}

		previousOrigins[origin] = true
		flight := findBestFlightForRoute(destination, finalDestination, previousOrigins, routeMap)
		if flight == nil || (bestFlightForRoute != nil && len(flight)+1 >= len(bestFlightForRoute)) {
			continue
		}
		flight = append([]model.Connection{connection}, flight...)
		bestFlightForRoute = flight
	}

	return bestFlightForRoute
}
