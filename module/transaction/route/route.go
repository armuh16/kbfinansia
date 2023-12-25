package route

import (
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/module/transaction/logic"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	fx.In
	Logic     logic.ITransactionLogic
	EchoRoute *router.Router
	Logger    *logger.LogRus
	Db        *mysql.DB
}

func NewRoute(h handler, m ...echo.MiddlewareFunc) handler {
	h.Route(m...)
	return h
}

func (h *handler) Route(m ...echo.MiddlewareFunc) {
	//transaction := h.EchoRoute.Group("/v1/transaction", m...)
	//transaction.POST("/detail", h.Create, h.EchoRoute.Authentication)
}
