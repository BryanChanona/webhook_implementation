package domain

type Webhook interface {
	SendMessage(msg string) error
}