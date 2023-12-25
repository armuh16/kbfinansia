package logic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/module/user/dto"
	"github.com/armuh16/kbfinansia/module/user/repository"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/static"
	"github.com/armuh16/kbfinansia/utilities"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// UserLogic
type IUserLogic interface {
	Find(context.Context, *dto.FindRequest) (*model.Users, error)
	CreateLimit(context.Context, *dto.CreateRequestLimit, *gorm.DB) error
	Create(context.Context, *dto.CreateRequest, *gorm.DB) error
	FindAll(context.Context, *dto.FindAllRequest) ([]*model.Tenor, error)
	FindTenor(context.Context, *dto.FindRequestTenor) (*model.Tenor, error)
}

type UserLogic struct {
	fx.In
	Logger   *logger.LogRus
	UserRepo repository.IUserRepository
}

// NewLogic :
func NewLogic(userLogic UserLogic) IUserLogic {
	return &userLogic
}

// FindByID
func (l *UserLogic) Find(ctx context.Context, reqData *dto.FindRequest) (*model.Users, error) {
	user, err := l.UserRepo.Find(ctx, &model.Users{
		ID:    reqData.ID,
		Email: reqData.Email,
	})
	if err != nil {
		l.Logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, utilities.ErrorRequest(fmt.Errorf(static.DataNotFound, "user"), http.StatusNotFound)
		}
		return nil, utilities.ErrorRequest(err, http.StatusInternalServerError)
	}
	return user, nil
}

func (l *UserLogic) FindTenor(ctx context.Context, reqData *dto.FindRequestTenor) (*model.Tenor, error) {
	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	tenor, err := l.UserRepo.FindTenor(ctx, &model.Tenor{
		ID: reqData.ID,
	})
	if err != nil {
		l.Logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, utilities.ErrorRequest(fmt.Errorf(static.DataNotFound, "tenor"), http.StatusNotFound)
		}
		return nil, utilities.ErrorRequest(err, http.StatusInternalServerError)
	}
	return tenor, nil
}

func (l *UserLogic) Create(ctx context.Context, reqData *dto.CreateRequest, tx *gorm.DB) error {
	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	// Find user
	if _, err := l.UserRepo.Find(ctx, &model.Users{
		ID: reqData.UserID,
	}); err != nil {
		l.Logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			return utilities.ErrorRequest(fmt.Errorf(static.DataNotFound, "data"), http.StatusNotFound)
		}
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	if _, err := l.UserRepo.Create(ctx, &model.UserDetails{
		UserID:       reqData.UserID,
		Nik:          reqData.Nik,
		FullName:     reqData.FullName,
		LegalName:    reqData.LegalName,
		PlaceOfBirth: reqData.PlaceOfBirth,
		DateOfBirth:  reqData.DateOfBirth,
		Salary:       reqData.Salary,
		UserKtp:      reqData.UserKtp,
		UserPhoto:    reqData.UserPhoto,
	}, tx); err != nil {
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	return nil
}

func (l *UserLogic) CreateLimit(ctx context.Context, reqData *dto.CreateRequestLimit, tx *gorm.DB) error {
	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	// Find user
	if _, err := l.UserRepo.Find(ctx, &model.Users{
		ID: reqData.UserID,
	}); err != nil {
		l.Logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			return utilities.ErrorRequest(fmt.Errorf(static.DataNotFound, "data"), http.StatusNotFound)
		}
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	// assuming to approve limit by admin
	if _, err := l.UserRepo.CreateLimit(ctx, &model.Tenor{
		UserID: reqData.UserID,
		Tenor:  reqData.Tenor,
		Limit:  reqData.Limit,
	}, tx); err != nil {
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	return nil
}

// FindAll
func (l *UserLogic) FindAll(ctx context.Context, reqData *dto.FindAllRequest) ([]*model.Tenor, error) {
	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return nil, utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	user, err := l.UserRepo.FindAll(ctx, &model.Tenor{
		UserID: reqData.UserID,
	})
	if err != nil {
		l.Logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, utilities.ErrorRequest(fmt.Errorf(static.DataNotFound, "user"), http.StatusNotFound)
		}
		return nil, utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	return user, nil
}
