package repository

import (
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/package/logger"
	"go.uber.org/fx"
)

// TransactionRepository
type ITransactionRepository interface {
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
