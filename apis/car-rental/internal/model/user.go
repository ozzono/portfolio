package model

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type User struct {
	UUID    uuid.UUID `json:"uuid,-" gorm:"primaryKey"`
	Name    string
	Contact string

	// it's a string, although containing `number` in its name
	PhoneNumber string
}

func (u User) Copy() *User {
	return &User{
		UUID:        u.UUID,
		Name:        u.Name,
		Contact:     u.Contact,
		PhoneNumber: u.PhoneNumber,
	}
}

func (u User) LogText() string {
	return fmt.Sprintf(`
	UUID --------- %v
	Name --------- %v
	Contact ------ %v
	PhoneNumber -- %v
	`,
		u.UUID,
		u.Name,
		u.Contact,
		u.PhoneNumber,
	)
}
