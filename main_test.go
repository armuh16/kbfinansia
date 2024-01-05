package main_test

import (
	"context"
	"github.com/armuh16/kbfinansia/config"
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/enum"
	"github.com/armuh16/kbfinansia/module"
	transactionRoute "github.com/armuh16/kbfinansia/module/transaction/route"
	"github.com/armuh16/kbfinansia/package/jwt"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	config.SetConfig()
	fx.New(
		fx.Provide(router.NewRouter),
		fx.Provide(mysql.NewMysql),
		fx.Provide(logger.NewLogRus),
		module.BundleRepository,
		module.BundleLogic,
		module.BundleRoute,
		fx.Invoke(NewRouteTest),
	).Start(context.Background())
}

type RouteTest struct {
	fx.In
	TransactionHandler transactionRoute.Handler
}

var r RouteTest

func NewRouteTest(routeTest RouteTest) {
	r = routeTest
}

func TestTransaction(t *testing.T) {
	// Initialize dependencies here (router, logger, etc.)

	t.Run("FailCreateOrderInvalidData", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
			"tenorID": 1,
			"contractNumber" : 123456789,
			"onTheRoad": 100000000,
    		"adminFee": 1000000,
			"installment": 6,
    		"interest": 10,
    		"assetName": "Mobil"
        }`))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)

		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, jwt.InternalClaimData{}, jwt.InternalClaimData{
			UserID: 2,
			Role:   enum.RoleTypeUser,
		})
		c.SetRequest(c.Request().WithContext(ctx))

		if assert.NoError(t, r.TransactionHandler.CreateOrder(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("SuccessCreateOrder", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
			"tenorID": 2,
			"contractNumber" : 123456789,
			"onTheRoad": 100000000,
    		"adminFee": 1000000,
			"installment": 6,
    		"interest": 10,
    		"assetName": "Mobil"
        }`))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)

		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, jwt.InternalClaimData{}, jwt.InternalClaimData{
			UserID: 2,
			Role:   enum.RoleTypeUser,
		})
		c.SetRequest(c.Request().WithContext(ctx))

		if assert.NoError(t, r.TransactionHandler.CreateOrder(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
