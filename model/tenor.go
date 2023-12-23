package model

import (
	"gorm.io/gorm"
	"time"
)

type Tenor struct {
	ID      int
	UserID  int
	AdminID int
	Tenor   int
	Limit   int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`

	Admin *Users `json:",omitempty" gorm:"<-:false;foreignKey:AdminID;references:ID;"`
	User  *Users `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`
}
