package notifier

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type INotifierFactory interface {
	SendMessage(message, to string) error
	SetWebhook(url string) error
}

type NotifierFactory struct{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while reading config %s", err)
	}
}

func CreateNotifierFactory() (NotifierFactory, error) {
	return NotifierFactory{}, nil
}

func (n *NotifierFactory) GetService(notifier Notifier) (INotifierFactory, error) {
	switch notifier {
	case Telegram:
		return n.createTelegramService()
	case Whatsapp:
		return n.createWhatsappService()
	case Sms:
		return nil, errors.New("sms service not implemented")
	default:
		return nil, errors.New("unknown service name")
	}
}

func (n *NotifierFactory) createTelegramService() (INotifierFactory, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	if telegramToken == "" {
		return nil, errors.New("telegram token is not set")
	}

	srv, err := NewTelegramService(telegramToken)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (n *NotifierFactory) createWhatsappService() (INotifierFactory, error) {
	instanceId, authToken := os.Getenv("WHATSAPP_INSTANCE"), os.Getenv("WHATSAPP_TOKEN")

	if instanceId == "" || authToken == "" {
		return nil, errors.New("whatsapp instanceId or authToken is missed")
	}

	srv, err := NewWhatsappService(instanceId, authToken)
	if err != nil {
		return nil, err
	}

	return srv, nil
}
