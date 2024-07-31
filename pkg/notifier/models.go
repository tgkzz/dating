package notifier

type Notifier string

const (
	Telegram Notifier = "Telegram"
	Whatsapp Notifier = "Whatsapp"
	Sms      Notifier = "Sms"
)

func (n Notifier) GetName() string {
	return string(n)
}
