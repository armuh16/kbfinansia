package route

import (
	"errors"
	"github.com/armuh16/kbfinansia/package/jwt"
	"net/http"

	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/module/user/dto"
	"github.com/armuh16/kbfinansia/module/user/logic"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"
	"github.com/armuh16/kbfinansia/static"
	"github.com/armuh16/kbfinansia/utilities"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	fx.In
	Logic     logic.IUserLogic
	EchoRoute *router.Router
	Logger    *logger.LogRus
	Db        *mysql.DB
}

func NewRoute(h handler, m ...echo.MiddlewareFunc) handler {
	h.Route(m...)
	return h
}

func (h *handler) Route(m ...echo.MiddlewareFunc) {
	auth := h.EchoRoute.Group("/v1/user", m...)
	auth.POST("/detail", h.Create, h.EchoRoute.Authentication)
	auth.POST("/limit", h.CreateLimit, h.EchoRoute.Authentication)
}

// Create
func (h *handler) Create(c echo.Context) error {
	var reqData = new(dto.CreateRequest)

	data, ok := c.Request().Context().Value(jwt.InternalClaimData{}).(jwt.InternalClaimData)
	if !ok {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusUnauthorized),
		})
	}

	reqData.UserID = data.UserID
	reqData.RoleID = data.Role

	if err := c.Bind(reqData); err != nil {
		h.Logger.Error(err)
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.BadRequest), http.StatusBadRequest),
		})
	}

	tx := h.Db.Gorm.Begin()
	if err := h.Logic.Create(c.Request().Context(), reqData, tx); err != nil {
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
		Code:   http.StatusCreated,
		Status: static.Success,
	})
}

// Create
func (h *handler) CreateLimit(c echo.Context) error {
	var reqData = new(dto.UpdateLimit)

	data, ok := c.Request().Context().Value(jwt.InternalClaimData{}).(jwt.InternalClaimData)
	if !ok {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusUnauthorized),
		})
	}

	//reqData.UserID = data.UserID
	reqData.AdminID = data.AdminID
	reqData.RoleID = data.Role

	if err := c.Bind(reqData); err != nil {
		h.Logger.Error(err)
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.BadRequest), http.StatusBadRequest),
		})
	}

	tx := h.Db.Gorm.Begin()
	if err := h.Logic.CreateLimit(c.Request().Context(), reqData, tx); err != nil {
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
		Code:   http.StatusCreated,
		Status: static.Success,
	})
}
