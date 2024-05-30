package telegramapi

import (
	"bytes"
	"net/http"

	"go.uber.org/zap"
)

type TelegramAPI struct {
	Logger *zap.Logger
}

func (t *TelegramAPI) SendResponsePhoto(jsonStruct []byte) {
	url := "https://api.telegram.org/bot7118393014:AAGzwGxS4XY4yGEa2dPBOt2GVu6QR7nOUdI/sendPhoto"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStruct))
	if err != nil {
		t.Logger.Error("Ошибка отправки ответа", zap.Error(err))
	}
	defer resp.Body.Close()
}

func (t *TelegramAPI) SendResponseMessage(jsonStruct []byte) {
	url := "https://api.telegram.org/bot7118393014:AAGzwGxS4XY4yGEa2dPBOt2GVu6QR7nOUdI/sendMessage"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStruct))
	if err != nil {
		t.Logger.Error("Ошибка отправки ответа", zap.Error(err))
	}
	defer resp.Body.Close()
}
