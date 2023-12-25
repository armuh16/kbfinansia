package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/armuh16/kbfinansia/static"
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	ID           int `json:"-" gorm:"primaryKey"`
	UserID       int `gorm:"index"`
	Nik          int
	FullName     string
	LegalName    string
	PlaceOfBirth string
	DateOfBirth  string
	Salary       int
	UserKtp      string // storing as url
	UserPhoto    string // storing as url

	// Relation
	//Asset []Assets `json:"Assets" gorm:"foreignKey:UserDetailID;references:ID;"`
	User  *Users `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`
	Tenor *Tenor `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (j UserDetails) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *UserDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(static.SomethingWrong)
	}

	result := UserDetails{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = result

	return nil
}
