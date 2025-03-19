package infraestructure

import (
	"log"
	"net/http"
	"webhook/src/Workflow/application"

	"github.com/gin-gonic/gin"
)

func HandleGitHubEvents(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("Nuevo evento: %s con ID: %s", eventType, deliveryID)

	rawData, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "leer datos"})
		return
	}

	var statusCode int

	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	case "workflow_run": 
		message := application.ProcessWorkflow(rawData)
		if message == "ERROR" {
			statusCode = 500
		} else {
			statusCode = post_discord_workflow(message)
		}
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"success": "Evento procesado con éxito"})
	case 500:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"success": "Normal"})
	}
}