package notifier

import "errors"

type INotifierFactory interface {
	SendMessage(message, to string) error
}

type NotifierFactory struct{}

func GetService(notifier Notifier) (INotifierFactory, error) {
	switch notifier {
	case Telegram:
		srv, err := NewTelegramService("default-telegram-token")
		if err != nil {
			return nil, err
		}
		return srv, nil
	case Whatsapp:
		srv, err := NewWhatsappService("default instance", "default token")
		if err != nil {
			return nil, err
		}
		return srv, nil
	case Sms:
		return nil, errors.New("sms service not implemented")
	default:
		return nil, errors.New("unknown service name")
	}
}
