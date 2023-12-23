package model

import (
	"time"

	"github.com/armuh16/kbfinansia/enum"
	"gorm.io/gorm"
)

type Users struct {
	ID         int
	Name       string
	Email      string
	Password   string        `json:"-"`
	Role       enum.RoleType `json:"-"`
	UserDetail UserDetails   `json:"UserDetail" gorm:"foreignKey:UserID;references:ID;"`
	Limit      []Tenor       `json:"-" gorm:"foreignKey:UserID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
