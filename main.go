package main

import (
	"fmt"

	"github.com/cagatygedik-discord-bot-go/bot"
	"github.com/cagatygedik-discord-bot-go/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
