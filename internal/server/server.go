package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура нашего сервера.
type Server struct {
	logger *log.Logger
	server *http.Server
}

// Функция NewServer принимает log.Logger и возвращает экземпляр структуры сервера.
func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.MainHandler)
	router.HandleFunc("/upload", handlers.HttpParcerHandler)

	// Структура сервера из задания
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: server,
	}
}

func (s *Server) Start() error {
	s.logger.Println("Сервер запущен на порту 8080")
	return s.server.ListenAndServe()
}
