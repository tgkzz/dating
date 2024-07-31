package notifier

import "errors"

type WhatsappService struct {
	InstanceId string
	AuthToken  string
}

func NewWhatsappService(instanceId, authToken string) (*WhatsappService, error) {
	if instanceId == "" || authToken == "" {
		return nil, errors.New("instanceId or authToken is missed")
	}

	return &WhatsappService{AuthToken: authToken, InstanceId: instanceId}, nil
}

func (ws *WhatsappService) SendMessage(message, to string) error {
	panic("implement me")
}
