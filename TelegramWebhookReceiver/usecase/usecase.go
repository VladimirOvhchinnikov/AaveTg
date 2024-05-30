package usecase

import (
	telegramapi "TelegramWebhookReceiver/infrastructure/telegramAPi"
	"TelegramWebhookReceiver/models"
	"encoding/json"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

type Case struct {
	Logger   *zap.Logger
	telegram telegramapi.TelegramAPI
}

func NewCase(logger *zap.Logger, tg telegramapi.TelegramAPI) *Case {
	return &Case{
		Logger:   logger,
		telegram: telegramapi.TelegramAPI{},
	}
}

func (c *Case) StartCase(chatID int64) error {
	var jsonStruct models.StartButton
	jsonFile, err := os.Open("/home/vladimir/Рабочий стол/BotTelegram/TelegramWebhookReceiver/models/startButton.JSON")
	if err != nil {
		c.Logger.Error("Чтение файла прошло не удачно", zap.Error(err))
		return err
	}
	defer jsonFile.Close()

	byteJsonFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		c.Logger.Error("Перевод в байты не удался", zap.Error(err))
		return err
	}

	err = json.Unmarshal(byteJsonFile, &jsonStruct)
	if err != nil {
		c.Logger.Error("Проблема с анмаршелингом", zap.Error(err))
		return err
	}

	jsonStruct.ChatID = chatID
	jsonStruct.Caption = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lectus risus, finibus ornare vestibulum et, feugiat quis dui. Vivamus sit amet dolor et magna facilisis rhoncus. Curabitur maximus est sed porta scelerisque. Sed suscipit, arcu volutpat feugiat posuere, eros nisi tristique nibh, mollis vehicula lectus tortor eu purus."

	// сделать отдельный пласт,который будет обращаться не в бизнес логике
	jsonStruct.Photo = "https://postimg.cc/YGZR4K5B"

	JsonOk, err := json.Marshal(jsonStruct)
	if err != nil {
		c.Logger.Error("Ошибка сборки структуруы", zap.Error(err))
		return err
	}

	c.telegram.Logger = c.Logger
	c.telegram.SendResponsePhoto(JsonOk)
	c.telegram = telegramapi.TelegramAPI{}

	return nil
}

func (c *Case) StartSubscribe(chatID int64) error {
	var jsonStruct models.SubscribeModel
	jsonFile, err := os.Open("/home/vladimir/Рабочий стол/BotTelegram/TelegramWebhookReceiver/models/SubscribeButton.json")
	if err != nil {
		c.Logger.Error("Чтение файла прошло не удачно", zap.Error(err))
		return err
	}
	defer jsonFile.Close()

	byteJsonFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		c.Logger.Error("Перевод в байты не удался", zap.Error(err))
		return err
	}

	err = json.Unmarshal(byteJsonFile, &jsonStruct)
	if err != nil {
		c.Logger.Error("Проблема с анмаршелингом", zap.Error(err))
		return err
	}

	jsonStruct.ChatID = int(chatID)

	JsonOk, err := json.Marshal(jsonStruct)
	if err != nil {
		c.Logger.Error("Ошибка сборки структуруы", zap.Error(err))
		return err
	}

	c.telegram.Logger = c.Logger
	c.telegram.SendResponseMessage(JsonOk)

	c.telegram = telegramapi.TelegramAPI{}

	return err
}
