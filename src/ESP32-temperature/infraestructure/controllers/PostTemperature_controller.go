package controllers

import (
	"net/http"
	"webhook/src/ESP32-temperature/application"
	"webhook/src/ESP32-temperature/domain"

	"github.com/gin-gonic/gin"
)
type PostSensorController struct {
	useCaseCreate *application.PostSensor
}

func NewPostSensorController(useCaseCreate *application.PostSensor) *PostSensorController{
	return &PostSensorController{useCaseCreate: useCaseCreate}
}
func (createPayment *PostSensorController) Execute(c *gin.Context){
	var sensor domain.Sensor

	
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createPayment.useCaseCreate.Execute(sensor.Menssage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusCreated, gin.H{"message": "pago registrado"})
}

