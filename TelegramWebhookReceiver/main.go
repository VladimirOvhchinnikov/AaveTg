package main

import (
	"TelegramWebhookReceiver/handlers"
	"TelegramWebhookReceiver/infrastructure/routing"
	"TelegramWebhookReceiver/infrastructure/serve"
	telegramapi "TelegramWebhookReceiver/infrastructure/telegramAPi"
	"TelegramWebhookReceiver/usecase"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Ошибка запсука логера")
		os.Exit(1)
	}
	tg := telegramapi.TelegramAPI{}
	usecase := usecase.NewCase(logger, tg)
	handlers := handlers.NewHandlers(logger, *usecase)
	router := routing.NewGoChiRouting(logger, *handlers)
	server := serve.NewServerHTTP(logger, ":8080", router)

	if err := server.Start(); err != nil {
		logger.Error("Ошибка при запуске сервера", zap.Error(err))
		os.Exit(1)
	}
}
