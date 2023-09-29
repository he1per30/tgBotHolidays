package main

import (
	"tgBotHolidays/internal"
)

func main() {

	// TODO инициализация конфига (viper) 1) Подключение к бд 2) Ключ к тг
	// Слой общения с базой
	// Слой с тг ботом

	// Вызов работы бота
	internal.StartBot()
}
