package main

import (
	"fmt"
	"tgBotHolidays/internal"
	"tgBotHolidays/internal/storage/postgre"
)

func main() {

	// TODO инициализация конфига (viper) 1) Подключение к бд 2) Ключ к тг
	// Слой общения с базой
	// Слой с тг ботом

	// Вызов работы бота

	t, err := postgre.NewPgConnect()

	if err != nil {
		fmt.Println("ОШИБКА")
		panic(err)
	}

	tgBot := internal.NewBot()
	internal.StartBot(tgBot, t)
}
