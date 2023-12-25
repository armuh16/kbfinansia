package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/package/logger"

	"go.uber.org/fx"
)

// UserRepository
type IUserRepository interface {
	Find(context.Context, *model.Users) (*model.Users, error)
	Create(context.Context, *model.UserDetails, *gorm.DB) (*int, error)
	CreateLimit(context.Context, *model.Tenor, *gorm.DB) (*int, error)
	FindAll(context.Context, *model.Tenor) ([]*model.Tenor, error)
}

type UserRepository struct {
	fx.In
	Logger   *logger.LogRus
	Database *mysql.DB
}

// NewRepository :
func NewRepository(userRepository UserRepository) IUserRepository {
	return &userRepository
}

// Find
func (l *UserRepository) Find(ctx context.Context, reqData *model.Users) (*model.Users, error) {
	product := new(model.Users)

	if err := l.Database.Gorm.WithContext(ctx).
		Where(&model.Users{
			ID:    reqData.ID,
			Email: reqData.Email,
		}).First(&product).Error; err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return product, nil
}

func (r *UserRepository) Create(ctx context.Context, reqData *model.UserDetails, tx *gorm.DB) (*int, error) {
	if err := tx.WithContext(ctx).Create(reqData).Error; err != nil {
		r.Logger.Error(err)
		return nil, err
	}
	return &reqData.ID, nil
}

func (r *UserRepository) CreateLimit(ctx context.Context, reqData *model.Tenor, tx *gorm.DB) (*int, error) {
	if err := tx.WithContext(ctx).Create(reqData).Error; err != nil {
		r.Logger.Error(err)
		return nil, err
	}
	return &reqData.ID, nil
}

// FindAll
func (l *UserRepository) FindAll(ctx context.Context, reqData *model.Tenor) ([]*model.Tenor, error) {
	tenor := []*model.Tenor{}

	if err := l.Database.Gorm.WithContext(ctx).Model(&model.Tenor{}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Where("user_id = ?", reqData.UserID).
		Order("id desc").
		Find(&tenor).
		Error; err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	return tenor, nil
}
