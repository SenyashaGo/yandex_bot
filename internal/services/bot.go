package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

type Bot struct {
	token string
	api   *tgbotapi.BotAPI
}

var Users = make(map[int64]string)

//конструктор структуры (инициализация бота)

func NewBot() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	return &Bot{
		token: "",
		api:   bot,
	}, nil
}

func (bot *Bot) Polling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			bot.Commands(update)
			continue
		}
		if update.Message != nil {
			if _, ok := Users[update.Message.Chat.ID]; !ok {
				continue
			}
			Users[update.Message.Chat.ID] = update.Message.Text
			parseInt, err := strconv.ParseInt(os.Getenv("ADMINCHAT"), 10, 0)

			if err != nil {
				panic(err)
			}

			_, err = bot.api.Send(tgbotapi.NewMessage(parseInt, "ChatID: "+strconv.Itoa(int(update.Message.Chat.ID))+"\n"+
				"Name: "+update.Message.Chat.FirstName+" "+update.Message.Chat.LastName+"\nText: "+update.Message.Text))
			delete(Users, update.Message.Chat.ID)
			if err != nil {
				log.Println(err) // logrus
			}

		}
	}
}

func (bot *Bot) Commands(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		bot.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "воспользуйся командой /question , чтобы задать вопрос"))
		return
	case "question":
		bot.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Задай свой вопрос👇"))
		Users[update.Message.Chat.ID] = ""
		return
	case "help":
		bot.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Здесь будет реализована поддержка"))
	}
}
