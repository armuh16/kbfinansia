package model

import (
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	ID           int
	UserID       int
	Nik          int
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  time.Time
	Salary       int
	UserKtp      string // storing as url
	UserPhoto    string // storing as url

	// Relation
	Asset []Assets `json:"assets" gorm:"foreignKey:UserID;references:ID;"`
	User  *Users   `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
