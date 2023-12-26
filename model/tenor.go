package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/armuh16/kbfinansia/static"
	"gorm.io/gorm"
	"time"
)

type Tenor struct {
	ID     int `json:"-" gorm:"primaryKey"`
	UserID int
	Tenor  int
	Limit  int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`

	User       *Users       `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`
	UserDetail *UserDetails `json:",omitempty" gorm:"<-:false;foreignKey:UserID;references:ID;"`
	Asset      *Assets      `json:",omitempty" gorm:"<-:false;foreignKey:TenorID;references:ID;"`
}

func (j Tenor) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Tenor) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(static.SomethingWrong)
	}

	result := Tenor{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = result

	return nil
}
