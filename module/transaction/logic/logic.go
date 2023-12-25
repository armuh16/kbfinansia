package logic

import (
	"github.com/armuh16/kbfinansia/module/transaction/repository"
	"github.com/armuh16/kbfinansia/package/logger"
	"go.uber.org/fx"
)

// TransactionLogic
type ITransactionLogic interface {
}

type TransactionLogic struct {
	fx.In
	Logger          *logger.LogRus
	TransactionRepo repository.ITransactionRepository
}

// NewLogic :
func NewLogic(transactionLogic TransactionLogic) ITransactionLogic {
	return &transactionLogic
}
