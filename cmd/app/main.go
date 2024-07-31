package main

import (
	"fmt"
	"github.com/tgkzz/dating/pkg/notifier"
)

func main() {
	notifierService, err := notifier.GetService(notifier.Telegram)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "kamal"
	to := "kamal"
	var cfg interface{}

	if err := notifierService.SendMessage(msg, to, cfg); err != nil {
		fmt.Println(err)
		return
	}
}
