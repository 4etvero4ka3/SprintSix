package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		logger.Println("Обработка запроса на /")
		res.Write([]byte("Добро пожаловать!"))
	})
	mux.HandleFunc("/alphabet", func(res http.ResponseWriter, req *http.Request) {
		logger.Println("Обработка запроса на /alphabet")
		res.Write([]byte("ОК"))
	})
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}
