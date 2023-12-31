package route

import (
	"errors"
	"net/http"

	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/module/auth/dto"
	"github.com/armuh16/kbfinansia/module/auth/logic"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"
	"github.com/armuh16/kbfinansia/static"
	"github.com/armuh16/kbfinansia/utilities"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	fx.In
	Logic     logic.IAuthLogic
	EchoRoute *router.Router
	Logger    *logger.LogRus
	Db        *mysql.DB
}

func NewRoute(h handler, m ...echo.MiddlewareFunc) handler {
	h.Route(m...)
	return h
}

func (h *handler) Route(m ...echo.MiddlewareFunc) {
	auth := h.EchoRoute.Group("/v1/auth", m...)
	auth.POST("/login", h.Login)
}

// Login
func (h *handler) Login(c echo.Context) error {
	var reqData = new(dto.LoginRequest)

	if err := c.Bind(reqData); err != nil {
		h.Logger.Error(err)
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.BadRequest), http.StatusBadRequest),
		})
	}

	tx := h.Db.Gorm.Begin()
	resp, err := h.Logic.Login(c.Request().Context(), reqData, tx)
	if err != nil {
		h.Logger.Error(err)
		defer func() {
			tx.Rollback()
		}()
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: err,
		})
	}
	tx.Commit()

	return utilities.Response(c, &utilities.ResponseRequest{
		Code:   http.StatusOK,
		Status: static.Success,
		Data:   resp,
	})
}
