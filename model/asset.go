package model

import (
	"time"

	"gorm.io/gorm"
)

type Assets struct {
	ID             int
	UserID         int
	TenorID        int
	ContractNumber int
	OnTheRoad      int
	AdminFee       int
	Installment    int
	Interest       float64
	AssetName      string
	GrandTotal     int
	Tenor          *Tenor `json:",omitempty" gorm:"<-:false;foreignKey:TenorID;references:ID;"`
	User           *Users `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
