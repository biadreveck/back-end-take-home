package main

import (
	"back-end-take-home/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.CreateRoutes(router.Group("/"))
	router.Run()
}
