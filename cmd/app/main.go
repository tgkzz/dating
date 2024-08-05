package main

import (
	"fmt"
	"github.com/tgkzz/dating/pkg/notifier"
)

func main() {
	not, err := notifier.CreateNotifierFactory()
	if err != nil {
		fmt.Println(err)
		return
	}

	telegramService, err := not.GetService(notifier.Telegram)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := telegramService.SendMessage("asdqwezxcttk", "me"); err != nil {
		fmt.Println(err)
		return
	}
}
