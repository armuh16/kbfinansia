package logic

import (
	"context"
	"fmt"
	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/module/transaction/dto"
	"github.com/armuh16/kbfinansia/module/transaction/repository"
	userDto "github.com/armuh16/kbfinansia/module/user/dto"
	userLogic "github.com/armuh16/kbfinansia/module/user/logic"
	userRepo "github.com/armuh16/kbfinansia/module/user/repository"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/utilities"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"math"
	"net/http"
)

// TransactionLogic
type ITransactionLogic interface {
	CreateOrder(context.Context, *dto.CreateOrderRequest, *gorm.DB) error
}

type TransactionLogic struct {
	fx.In
	Logger          *logger.LogRus
	UserLogic       userLogic.IUserLogic
	UserRepo        userRepo.IUserRepository
	TransactionRepo repository.ITransactionRepository
}

// NewLogic :
func NewLogic(transactionLogic TransactionLogic) ITransactionLogic {
	return &transactionLogic
}

// Create
func (l *TransactionLogic) CreateOrder(ctx context.Context, reqData *dto.CreateOrderRequest, tx *gorm.DB) error {
	// Validate request data
	if err := reqData.Validate(); err != nil {
		l.Logger.Error(err)
		return utilities.ErrorRequest(err, http.StatusBadRequest)
	}

	// Find detail user
	userDetail, err := l.UserLogic.Find(ctx, &userDto.FindRequest{
		ID: reqData.UserID,
	})
	if err != nil {
		l.Logger.Error(err)
		return err
	}

	tenorDetail, err := l.UserLogic.FindTenor(ctx, &userDto.FindRequestTenor{
		ID: reqData.TenorID,
	})
	if err != nil {
		l.Logger.Error(err)
		return err
	}

	if tenorDetail.Limit < reqData.OnTheRoad {
		return utilities.ErrorRequest(fmt.Errorf("limit tidak mencukupi"), http.StatusBadRequest)
	}

	tenorDetail.Limit -= reqData.OnTheRoad

	// Update the tenor with the new limit
	if err := l.UserRepo.UpdateTenorLimit(ctx, tenorDetail, tx); err != nil {
		l.Logger.Error(err)
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	if reqData.Installment != tenorDetail.Tenor {
		return utilities.ErrorRequest(fmt.Errorf("tenor tidak ditemukan"), http.StatusBadRequest)
	}

	adminFeePerInstallment := float64(reqData.AdminFee) / float64(reqData.Installment)
	subtotal := float64(reqData.OnTheRoad) + adminFeePerInstallment

	interestAmount := subtotal * (reqData.Interest / 100)

	grandTotal := subtotal + interestAmount

	reqData.GrandTotal = int(math.Round(grandTotal))

	if _, err := l.TransactionRepo.Create(ctx, &model.Assets{
		UserID:         userDetail.ID,
		TenorID:        tenorDetail.ID,
		ContractNumber: reqData.ContractNumber,
		OnTheRoad:      reqData.OnTheRoad,
		AdminFee:       reqData.AdminFee,
		Installment:    reqData.Installment,
		Interest:       reqData.Interest,
		AssetName:      reqData.AssetName,
		GrandTotal:     reqData.GrandTotal,
	}, tx); err != nil {
		return utilities.ErrorRequest(err, http.StatusInternalServerError)
	}

	return nil
}
