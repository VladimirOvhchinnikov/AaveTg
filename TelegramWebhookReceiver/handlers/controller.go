package handlers

import (
	"TelegramWebhookReceiver/models"
	"TelegramWebhookReceiver/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Handlers struct {
	Logger  *zap.Logger
	Update  Updates
	Usecase usecase.Case
}

type Updates struct {
	models.ButtonStandard
	models.UserInfo
}

// NewHandlers создает экземпляр Handlers.
func NewHandlers(logger *zap.Logger, usecase usecase.Case) *Handlers {
	return &Handlers{
		Logger: logger,
	}
}

// CommandHandler обрабатывает входящие команды.
func (h *Handlers) CommandHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&h.Update)
	if err != nil {
		h.Logger.Error("Ошибка в декодировании ", zap.Error(err))
		h.sendResponse(w, "Ошибка в декодировании: "+err.Error(), http.StatusBadRequest)
		return
	}

	if h.Update.Message.Text != "" {
		err := h.messageHandler(w, r)
		h.Update = Updates{}
		if err != nil {
			h.Logger.Error(err.Error())
			h.sendResponse(w, "Не правильный путь", http.StatusBadRequest)
			return
		}
	}

	if h.Update.CallbackQuery.Data != "" {
		err := h.callbackQueryHandler(w, r)
		h.Update = Updates{}
		if err != nil {
			h.Logger.Error(err.Error())
			h.sendResponse(w, "Не правильный путь", http.StatusBadRequest)
			return
		}
	}

}

// messageHandler обрабатывает сообщения от пользователей.
func (h *Handlers) messageHandler(w http.ResponseWriter, r *http.Request) error {

	switch h.Update.Message.Text {
	case "/start":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			h.startHandler(w, r)
		}

	default:
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			fmt.Println(h.Update.Message.Text)
			return errors.New("Путь messageHandler не найден")
		}
	}
	return nil
}

// callbackQueryHandler обрабатывает нажатия на кнопки.
func (h *Handlers) callbackQueryHandler(w http.ResponseWriter, r *http.Request) error {
	switch h.Update.CallbackQuery.Data {
	case "Start":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			h.option1Handler(w, r)
		}
	case "Subscribe":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			h.Subscribe(w, r)
		}

	case "Back":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			h.startHandlerToButton(w, r)
		}
	case "200points":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			
		}
	default:
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			return errors.New("Путь callbackQueryHandler не найден")
		}
	}

	return nil
}

// StartHandler обрабатывает команду /start.
func (h *Handlers) startHandler(w http.ResponseWriter, r *http.Request) {

	h.Usecase.StartCase(h.Update.Message.Chat.ID)

}

func (h *Handlers) startHandlerToButton(w http.ResponseWriter, r *http.Request) {
	h.Usecase.StartCase(h.Update.CallbackQuery.Message.Chat.ID)
}

// option1Handler обрабатывает событие нажатия на кнопку "option1".
func (h *Handlers) option1Handler(w http.ResponseWriter, r *http.Request) {

	//h.Usecase.StartButton(h.Update.ButtonStandard.CallbackQuery.Message.Chat.ID)
}

// Subscribe обрабатывает события нажатия на кнопку подписки
func (h *Handlers) Subscribe(w http.ResponseWriter, r *http.Request) {
	h.Usecase.StartSubscribe(h.Update.ButtonStandard.CallbackQuery.Message.Chat.ID)
}

// ApiResponse структура ответа API.
type ApiResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// sendResponse отправляет ответ клиенту.
func (h *Handlers) sendResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ApiResponse{
		Message: message,
		Success: statusCode >= 200 && statusCode < 300,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.Logger.Error("Не получилось создать ответ", zap.Error(err))
	}
}

// notFoundTellegramEndPoint отправляет ответ о том, что конечная точка не найдена.
func (h *Handlers) notFoundTellegramEndPoint(w http.ResponseWriter) {
	h.sendResponse(w, "Операция не совершено удачно", http.StatusNotFound)
}
