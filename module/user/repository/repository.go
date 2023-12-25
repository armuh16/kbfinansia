package repository

import (
	"context"
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/package/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go.uber.org/fx"
)

// UserRepository
type IUserRepository interface {
	Find(context.Context, *model.Users) (*model.Users, error)
	Create(context.Context, *model.UserDetails, *gorm.DB) (*int, error)
	CreateLimit(context.Context, *model.Tenor, *gorm.DB) (*int, error)
	FindAll(context.Context, *model.Tenor) ([]*model.Tenor, error)
	FindTenor(context.Context, *model.Tenor) (*model.Tenor, error)
	UpdateTenorLimit(context.Context, *model.Tenor, *gorm.DB) error
	FindTenorWithLock(context.Context, int, *gorm.DB) (*model.Tenor, error)
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
	user := new(model.Users)

	if err := l.Database.Gorm.WithContext(ctx).
		Where(&model.Users{
			ID:    reqData.ID,
			Email: reqData.Email,
		}).First(&user).Error; err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return user, nil
}

// Find tenor
func (l *UserRepository) FindTenor(ctx context.Context, reqData *model.Tenor) (*model.Tenor, error) {
	tenor := new(model.Tenor)

	if err := l.Database.Gorm.WithContext(ctx).
		Where("id = ?", reqData.ID).
		First(&tenor).Error; err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return tenor, nil
}

func (l *UserRepository) FindTenorWithLock(ctx context.Context, tenorID int, tx *gorm.DB) (*model.Tenor, error) {
	var tenor model.Tenor
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", tenorID).First(&tenor).Error; err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return &tenor, nil
}

// Update Tenor
func (r *UserRepository) UpdateTenorLimit(ctx context.Context, reqData *model.Tenor, tx *gorm.DB) error {
	if err := tx.WithContext(ctx).Model(reqData).Update("limit", reqData.Limit).Error; err != nil {
		r.Logger.Error(err)
		return err
	}
	return nil
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
