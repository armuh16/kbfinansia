package logic

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/armuh16/kbfinansia/config"
	"github.com/armuh16/kbfinansia/module/auth/dto"
	userDto "github.com/armuh16/kbfinansia/module/user/dto"
	userLogic "github.com/armuh16/kbfinansia/module/user/logic"
	"github.com/armuh16/kbfinansia/static"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/armuh16/kbfinansia/package/jwt"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/utilities"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// AuthLogic
type IAuthLogic interface {
	Login(context.Context, *dto.LoginRequest, *gorm.DB) (*dto.Response, error)
}

type AuthLogic struct {
	fx.In
	Logger    *logger.LogRus
	UserLogic userLogic.IUserLogic
}

// NewLogic :
func NewLogic(AuthLogic AuthLogic) IAuthLogic {
	return &AuthLogic
}

// Login
func (l *AuthLogic) Login(ctx context.Context, reqData *dto.LoginRequest, tx *gorm.DB) (*dto.Response, error) {

	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	// Check exist user by email
	userDetail, err := l.UserLogic.Find(ctx, &userDto.FindRequest{
		Email: reqData.Email,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(reqData.Password)); err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(errors.New(static.InvalidAccessLogin), http.StatusBadRequest)
	}

	// Generate uuid for user jwt
	uuid, err := uuid.NewV4()
	if err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	// Generate access and refresh token
	token, err := jwt.RequestToken(ctx, jwt.ClaimData{
		UserID: userDetail.ID,
		UUID:   uuid.String(),
	}, config.Get().Secret, time.Now().Add(config.Get().ExpireAccessTokenDuration).Unix(), time.Now().Add(config.Get().ExpireRefreshTokenDuration).Unix())
	if err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	return &dto.Response{
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

// TODO
