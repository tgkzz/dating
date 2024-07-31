package notifier

import (
	"errors"
	"fmt"
)

type TelegramService struct {
	Token string
}

func NewTelegramService(token string) (*TelegramService, error) {
	if token == "" {
		return nil, errors.New("no token")
	}

	return &TelegramService{Token: token}, nil
}

func (t *TelegramService) SendMessage(message, to string) error {
	fmt.Println("Sending message via Telegram to:", to, "Message:", message)
	return nil
}
