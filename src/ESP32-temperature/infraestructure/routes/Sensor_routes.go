package routes

import (
	"webhook/src/ESP32-temperature/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/sensors")

	postSensorController := dependencies.GetPostSensorController()

	routes.POST("/", postSensorController.Execute) 

}

