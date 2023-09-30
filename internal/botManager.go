package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
	"tgBotHolidays"
	"tgBotHolidays/internal/message"
)

var gBot *tgbotapi.BotAPI

// TODO какой тут должен быть интерфейс и какие методы поддерживать
type WorkWithUser interface {
	AddUser() string
	GetUser() string
	CheckUser() string
	RegUser() string
}

func NewBot() *tgbotapi.BotAPI {
	var tgToken string

	if tgToken = os.Getenv("tg-token"); tgToken == "" {
		panic("token is empty")
	}

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}

// TODO Почему тут если сделать db *WorkWithUser в main это не работает?
func StartBot(bot *tgbotapi.BotAPI, db WorkWithUser) {
	gBot = bot // TODO Можно ли так делать?
	log.Printf("Authorized on account %s", gBot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = tgBotHolidays.UpdateConfigTimeout

	updates := gBot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			checkCommand(update, db)
		}
	}
}

// "/start test"
// checkCommand проверяет введенную в бот команду
func checkCommand(update tgbotapi.Update, db WorkWithUser) {
	command := strings.ToLower(strings.Split(update.Message.Text, " ")[0])

	switch command {
	case "/start":
		startMessage(update)
	case "/help":
		helpMessage(update)
	case "/register":
		regUser(update, db)
	case "/testAdmin":
		testAdmin(update, db)
	}

}

// функция для всяких тестов
func testAdmin(update tgbotapi.Update, db WorkWithUser) {

	db.CheckUser()
	db.AddUser()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.RegUser)
	gBot.Send(msg)
}

func regUser(update tgbotapi.Update, db WorkWithUser) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.RegUser)
	db.RegUser()
	gBot.Send(msg)
}

func helpMessage(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.HelpMessage)
	gBot.Send(msg)
}

func startMessage(update tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	//TODO переделать сообщение
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, message.HelpMessage)
	if !checkUser() {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, message.StartMessage)
	}

	gBot.Send(msg)

}

// Проверяем есть ли этот юзер в системе
func checkUser() bool {
	//TODO добавить проверку пользователя
	return false
}

func checkRegister() {
	//TODO сделать проверку регистрации
}

// connectToBot осуществляет коннект к тг и проверку подключение
func connectToBot() {

}
