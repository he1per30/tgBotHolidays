package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
	"tgBotHolidays"
	"tgBotHolidays/internal/message"
)

var gBot2 *tgbotapi.BotAPI

type addUser interface {
}

func init() {
	var tgToken string

	if tgToken = os.Getenv("tg-token"); tgToken == "" {
		panic("token is empty")
	}

	var err error

	if gBot2, err = tgbotapi.NewBotAPI(tgToken); err != nil {
		log.Panic(err)
	}

	gBot2.Debug = true

}

func StartBot() {
	log.Printf("Authorized on account %s", gBot2.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = tgBotHolidays.UpdateConfigTimeout

	updates := gBot2.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			checkCommand(update.Message.Text, update)
		}
	}
}

// "/start test"
// checkCommand проверяет введенную в бот команду
func checkCommand(message string, update tgbotapi.Update) {
	command := strings.ToLower(strings.Split(message, " ")[0])

	switch command {
	case "/start":
		startMessage(update)
	case "/help":
		helpMessage(update)
	case "/register":
		regUser(update)
	}

}

func regUser(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.HelpMessage)
	gBot2.Send(msg)
}

func helpMessage(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.HelpMessage)
	gBot2.Send(msg)
}

func startMessage(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.StartMessage)
	gBot2.Send(msg)
}

func checkRegister() {
	//TODO сделать проверку регистрации
}

// connectToBot осуществляет коннект к тг и проверку подключение
func connectToBot() {

}
