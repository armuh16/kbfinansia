package route

import (
	"errors"
	"github.com/armuh16/kbfinansia/enum"
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
	user := h.EchoRoute.Group("/v1/user", m...)
	user.POST("/detail", h.Create, h.EchoRoute.Authentication)
	user.POST("/limit", h.CreateLimit, h.EchoRoute.Authentication)
	user.GET("/limit", h.FindAll, h.EchoRoute.Authentication)
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

	if data.Role != enum.RoleTypeUser {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusForbidden),
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
	var reqData = new(dto.CreateRequestLimit)

	data, ok := c.Request().Context().Value(jwt.InternalClaimData{}).(jwt.InternalClaimData)
	if !ok {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusUnauthorized),
		})
	}

	// Ensure that the role is Admin
	if data.Role != enum.RoleTypeAdmin {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusForbidden),
		})
	}

	reqData.AdminID = data.UserID
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

// FindAll
func (h *handler) FindAll(c echo.Context) error {
	var reqData = new(dto.FindAllRequest)

	data, ok := c.Request().Context().Value(jwt.InternalClaimData{}).(jwt.InternalClaimData)
	if !ok || data.Role != enum.RoleTypeUser {
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusUnauthorized),
		})
	}

	//if data.Role != enum.RoleTypeUser {
	//	return utilities.Response(c, &utilities.ResponseRequest{
	//		Error: utilities.ErrorRequest(errors.New(static.Authorization), http.StatusForbidden),
	//	})
	//}

	reqData.UserID = data.UserID

	resp, err := h.Logic.FindAll(c.Request().Context(), reqData)
	if err != nil {
		h.Logger.Error(err)
		return utilities.Response(c, &utilities.ResponseRequest{
			Error: err,
		})
	}

	return utilities.Response(c, &utilities.ResponseRequest{
		Code:   http.StatusOK,
		Status: static.Success,
		Data:   resp,
	})
}
