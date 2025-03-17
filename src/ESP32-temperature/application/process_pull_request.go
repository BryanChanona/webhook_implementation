package application

import (
	"encoding/json"
	"webhook/src/ESP32-temperature/domain/value_objects"

	"log"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload value_objects.PullRequestEventPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403
	}

	if eventPayload.Action == "closed" {
		log.Printf("Evento Pull Request recibido: Rama Base=%s, Rama Default='%s',User='%s', Repositorio='%s'",
		eventPayload.PullRequest.Base.Ref, eventPayload.PullRequest.Head.Ref,
		eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)
	}

	return 200
}
