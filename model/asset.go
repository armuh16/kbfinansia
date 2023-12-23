package model

import (
	"time"

	"gorm.io/gorm"
)

type Assets struct {
	ID             int
	UserID         int
	ContractNumber int
	OnTheRoad      int
	AdminFee       int
	Installment    int
	Interest       float64
	AssetName      string
	//Limit          []Tenor `json:"Limit" gorm:"foreignKey:AssetID;references:ID;"`
	User *Users `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
