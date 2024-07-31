package notifier

import (
	"context"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"time"
)

const (
	TelegramBaseUrl = "https://api.telegram.org/bot"
)

const (
	JsonType           = "application/json"
	FormDataType       = "multipart/form-data"
	FormUrlencodedType = "application/x-www-form-urlencoded"

	TelegramSecretHeader = "X-Telegram-Bot-Api-Secret-Token"

	Method = "method"
)

type (
	SendMessageReq struct {
		ChatId string `json:"chat_id"`
		Text   string `json:"text"`
	}
)

// TODO: add another methods
type TelegramService struct {
	client *req.Client
}

// TODO: add req to set telegram webhook
func NewTelegramService(token string) (*TelegramService, error) {
	if token == "" {
		return nil, errors.New("no token")
	}

	client := req.C().SetBaseURL(TelegramBaseUrl + token)

	method := "getMe"

	getMeResp, err := req.C().R().SetPathParam(Method, method).Get(fmt.Sprintf("/{%s}", Method))

	if err != nil {
		return nil, fmt.Errorf("telegram returned error while logging %s", err.Error())
	}

	if getMeResp.IsErrorState() {
		return nil, errors.New("telegram send error while checking token")
	}

	return &TelegramService{client: client}, nil
}

func (t *TelegramService) SendMessage(message, to string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	method := "sendMessage"

	body := SendMessageReq{Text: message, ChatId: to}

	resp, err := t.client.R().
		SetContext(ctx).
		SetPathParam(Method, method).
		SetBody(&body).
		SetContentType(JsonType).
		Get(fmt.Sprintf("/{%s}", Method))
	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		return errors.New("telegram send error while sending error")
	}

	return nil
}
