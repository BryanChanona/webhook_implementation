package main

import (
	"os"
	infrastructurePull "webhook/src/Pull_request_webhook/infraestructure"
	infraestructureWork "webhook/src/Workflow/infraestructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine := gin.Default()

	infrastructurePull.Routes(engine)
	infraestructureWork.Routes(engine)

	engine.Run(":" + port)
}
