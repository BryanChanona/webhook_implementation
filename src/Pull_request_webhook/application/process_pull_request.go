package application

import (
	"encoding/json"
	"log"
	"webhook/src/Pull_request_webhook/domain/value_objects"
)

func ProcessPullRequestEvent(rawData []byte) string {
	var eventPayload value_objects.PullRequestEventPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload: %v", err)
		return "ERROR"
	}
	log.Printf("Pull Request Event: %v", eventPayload.Action)

	base := eventPayload.PullRequest.Base.Ref
	titulo := eventPayload.PullRequest.Title
	repoFullName := eventPayload.Repository.FullName
	user := eventPayload.PullRequest.User.Login
	urlPullRequest := eventPayload.PullRequest.HTMLURL
	merged := eventPayload.PullRequest.Merged
	action := eventPayload.Action
	return GenerateMessageToDiscord(action, base, titulo, repoFullName, user, urlPullRequest, merged)
}
