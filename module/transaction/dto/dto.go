package dto

import (
	"errors"
	"fmt"
	"github.com/armuh16/kbfinansia/enum"
	"github.com/armuh16/kbfinansia/static"
)

type CreateOrderRequest struct {
	UserID         int
	ContractNumber int
	OnTheRoad      int
	AdminFee       int
	Installment    int
	Interest       float64
	AssetName      string
	RoleID         enum.RoleType
}

func (d *CreateOrderRequest) Validate() error {
	if d.UserID <= 0 {
		return fmt.Errorf(static.EmptyValue, "UserID")
	}
	if d.ContractNumber <= 0 {
		return fmt.Errorf(static.EmptyValue, "Nomor Kontrak")
	}
	if d.OnTheRoad <= 0 {
		return fmt.Errorf(static.EmptyValue, "On The Road")
	}
	if d.AdminFee <= 0 {
		return fmt.Errorf(static.EmptyValue, "Admin Fee")
	}
	if d.Installment <= 0 {
		return fmt.Errorf(static.EmptyValue, "Cicilan")
	}
	if d.Interest <= 0 {
		return fmt.Errorf(static.EmptyValue, "Bunga")
	}
	if d.AssetName == "" {
		return fmt.Errorf(static.EmptyValue, "Nama Asset")
	}
	if err := d.RoleID.IsValid(); err != nil {
		return err
	}
	if d.RoleID != enum.RoleTypeUser {
		return errors.New(static.Authorization)
	}
	return nil
}
