package main

import (
	"webhook/src/ESP32-temperature/infraestructure/dependencies"
	"webhook/src/ESP32-temperature/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	defer dependencies.CloseDB()

	r := gin.Default()

	

	routes.Routes(r)

	r.Run(":8080")
}
