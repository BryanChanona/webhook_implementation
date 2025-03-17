package application

import (
	"fmt"
	"webhook/src/ESP32-temperature/domain"
)

type WebHookService struct {
	webhook domain.Webhook
}

func NewWebHookService(w domain.Webhook) *WebHookService {
	return &WebHookService{webhook: w}
}

func (s *WebHookService) Notify(ramaBase, repositorio, titulo, user, urlDetails string) error {
	message := fmt.Sprintf(
		"Nuevo pull request a la rama %s en el repositorio %s\nTítulo: %s\nAutor: %s\nDetalles: %s",
		ramaBase, repositorio, titulo, user, urlDetails,
	)
	return s.webhook.SendMessage(message)
}
