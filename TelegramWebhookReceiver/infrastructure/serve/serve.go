package serve

import (
	"net/http"

	"go.uber.org/zap"
)

type ServerHTTP struct {
	Logger  *zap.Logger
	Address string
	Handler http.Handler
}

func NewServerHTTP(logger *zap.Logger, address string, handler http.Handler) *ServerHTTP {
	return &ServerHTTP{
		Logger:  logger,
		Address: address,
		Handler: handler,
	}
}

func (sh *ServerHTTP) Start() error {
	sh.Logger.Info("Запуск сервера на адресе", zap.String("address", sh.Address))
	err := http.ListenAndServe(sh.Address, sh.Handler)
	if err != nil {
		sh.Logger.Error("Ошибка в старте HTTP сервера", zap.Error(err))
		return err
	}

	sh.Logger.Info("Сервер успешно запущен")
	return nil
}

func (sh *ServerHTTP) Stop() error {
	return nil
}

func (sh *ServerHTTP) Restart() error {
	return nil
}

func (sh *ServerHTTP) SetRouter() error {
	return nil
}

func (sh *ServerHTTP) Configure(options map[string]interface{}) error {
	return nil
}
