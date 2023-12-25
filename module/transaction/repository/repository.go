package repository

import (
	"context"
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/package/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// TransactionRepository
type ITransactionRepository interface {
	Create(context.Context, *model.Assets, *gorm.DB) (*int, error)
}

type TransactionRepository struct {
	fx.In
	Logger   *logger.LogRus
	Database *mysql.DB
}

// NewRepository :
func NewRepository(transactionRepository TransactionRepository) ITransactionRepository {
	return &transactionRepository
}

func (r *TransactionRepository) Create(ctx context.Context, reqData *model.Assets, tx *gorm.DB) (*int, error) {
	if err := tx.WithContext(ctx).Create(&reqData).Error; err != nil {
		r.Logger.Error(err)
		return nil, err
	}
	return &reqData.ID, nil
}
