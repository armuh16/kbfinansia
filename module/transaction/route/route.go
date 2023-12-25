package route

import (
	"errors"
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/module/transaction/dto"
	"github.com/armuh16/kbfinansia/module/transaction/logic"
	"github.com/armuh16/kbfinansia/package/jwt"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"
	"github.com/armuh16/kbfinansia/static"
	"github.com/armuh16/kbfinansia/utilities"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type Handler struct {
	fx.In
	Logic     logic.ITransactionLogic
	EchoRoute *router.Router
	Logger    *logger.LogRus
	Db        *mysql.DB
}

func NewRoute(h Handler, m ...echo.MiddlewareFunc) Handler {
	h.Route(m...)
	return h
}

func (h *Handler) Route(m ...echo.MiddlewareFunc) {
	transaction := h.EchoRoute.Group("/v1/transaction", m...)
	transaction.POST("/order", h.CreateOrder, h.EchoRoute.Authentication)
}

// CreateOrder
func (h *Handler) CreateOrder(c echo.Context) error {
	var reqData = new(dto.CreateOrderRequest)

	if err := c.Bind(reqData); err != nil {
		h.Logger.Error(err)
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.BadRequest), http.StatusBadRequest),
		})
	}

	data, ok := c.Request().Context().Value(jwt.InternalClaimData{}).(jwt.InternalClaimData)
	if !ok {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusUnauthorized),
		})
	}

	reqData.UserID = data.UserID
	reqData.RoleID = data.Role

	tx := h.Db.Gorm.Begin()
	if err := h.Logic.CreateOrder(c.Request().Context(), reqData, tx); err != nil {
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
	})
}
