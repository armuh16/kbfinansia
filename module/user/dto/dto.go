package dto

import (
	"errors"
	"fmt"
	"github.com/armuh16/kbfinansia/enum"
	"github.com/armuh16/kbfinansia/model"
	"github.com/armuh16/kbfinansia/static"
)

type FindRequest model.Users

// type FindRequestTenor model.Tenor
type FindRequestTenor struct {
	ID int
}

func (d *FindRequestTenor) Validate() error {
	if d.ID <= 0 {
		return fmt.Errorf(static.EmptyValue, "TenorID")
	}
	return nil
}

type CreateRequest struct {
	UserID       int
	Nik          int
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  string
	Salary       int
	UserKtp      string
	UserPhoto    string
	RoleID       enum.RoleType
}

func (d *CreateRequest) Validate() error {
	if d.UserID <= 0 {
		return fmt.Errorf(static.EmptyValue, "UserID")
	}
	if d.Nik <= 0 {
		return fmt.Errorf(static.EmptyValue, "Nik")
	}
	if d.FullName == "" {
		return fmt.Errorf(static.EmptyValue, "FullName")
	}
	if d.LegalName == "" {
		return fmt.Errorf(static.EmptyValue, "LegalName")
	}
	if d.PlaceOfBirth == "" {
		return fmt.Errorf(static.EmptyValue, "PlaceOfBirth")
	}
	if d.Salary <= 0 {
		return fmt.Errorf(static.MinValue, "Salary", 0)
	}
	if d.DateOfBirth == "" {
		return fmt.Errorf(static.MinValue, "DateOfBirth", 0)
	}
	if d.UserKtp == "" {
		return fmt.Errorf(static.EmptyValue, "UserKtp")
	}
	if d.UserPhoto == "" {
		return fmt.Errorf(static.EmptyValue, "UserPhoto")
	}
	if err := d.RoleID.IsValid(); err != nil {
		return err
	}
	if d.RoleID != enum.RoleTypeUser {
		return errors.New(static.Authorization)
	}
	return nil
}

type SetLimitsRequest struct {
	UserID int
	Limits []CreateRequestLimit
}

// Request Limit to upgrade
type CreateRequestLimit struct {
	UserID  int
	AdminID int
	Tenor   int
	Limit   int
	RoleID  enum.RoleType
}

func (d *CreateRequestLimit) Validate() error {
	if d.UserID <= 0 {
		return fmt.Errorf(static.EmptyValue, "UserID")
	}
	//if d.AdminID <= 0 {
	//	return fmt.Errorf(static.EmptyValue, "AdminID")
	//}
	if d.Tenor <= 0 {
		return fmt.Errorf(static.EmptyValue, "Tenor")
	}
	if d.Limit <= 0 {
		return fmt.Errorf(static.EmptyValue, "Limit")
	}
	if err := d.RoleID.IsValid(); err != nil {
		return err
	}
	if d.RoleID != enum.RoleTypeAdmin {
		return errors.New(static.Authorization)
	}
	return nil
}

type FindAllRequest struct {
	UserID int
}

func (d *FindAllRequest) Validate() error {
	if d.UserID <= 0 {
		return fmt.Errorf(static.EmptyValue, "UserID")
	}
	return nil
}
