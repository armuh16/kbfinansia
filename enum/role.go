package enum

import (
	"fmt"

	"github.com/armuh16/kbfinansia/static"
)

type RoleType int

const (
	RoleTypeAdmin RoleType = 1
	RoleTypeUser  RoleType = 2
)

func (t RoleType) String() string {
	switch t {
	case RoleTypeAdmin:
		return "Administrator"
	case RoleTypeUser:
		return "Users"
	default:
		return "Unknown"
	}
}

func (t RoleType) IsValid() error {
	switch t {
	case RoleTypeAdmin, RoleTypeUser:
		return nil
	}
	return fmt.Errorf(static.DataNotFound, "Role")
}
