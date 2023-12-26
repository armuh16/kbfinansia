package model

import (
	"time"

	"github.com/armuh16/kbfinansia/enum"
	"gorm.io/gorm"
)

type Users struct {
	ID       int `json:"-" gorm:"primaryKey"`
	Name     string
	Email    string
	Password string        `json:"-"`
	Role     enum.RoleType `json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
