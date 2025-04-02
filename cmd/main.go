package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags) // создается логер

	srv := server.NewServer(logger) // Создается сервер с помощью функции NewServer из пакета server

	// Запускается сервер
	if err := srv.Start(); err != nil {
		// Если возникла ошибка, выводим её в лог и завершаем программу
		logger.Fatal("Error starting server:", err)
	}

}

// http://localhost:8080/ - для проверки
