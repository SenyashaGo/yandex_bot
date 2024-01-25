package main

import "github.com/SenyashaGo/yandex_bot/internal/services"

func main() {

	bot, err := services.NewBot()
	if err != nil {
		panic(err)
	}
	bot.Polling()
}
