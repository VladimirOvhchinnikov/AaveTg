package routing

import (
	"TelegramWebhookReceiver/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type GoChiRouter struct {
	Logger  *zap.Logger
	Router  chi.Router
	Handler handlers.Handlers
}

func NewGoChiRouting(logger *zap.Logger, handler handlers.Handlers) *GoChiRouter {

	router := chi.NewRouter()
	router.Post("/telegram-webhook", handler.CommandHandler)

	return &GoChiRouter{
		Logger: logger,
		Router: router,
	}
}

func (gc *GoChiRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gc.Router.ServeHTTP(w, r)
}
