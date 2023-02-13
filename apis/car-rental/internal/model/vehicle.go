package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Vehicle ...
type Vehicle struct {
	UUID         uuid.UUID `json:"uuid,-" gorm:"primaryKey"`
	Model        string
	LicensePlate string `gorm:"size:255"`
	State        string `gorm:"size:255"`
	Archived     bool   `gorm:"default:false"`
	Available    bool
	Year         int16
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (v Vehicle) Copy() *Vehicle {
	return &Vehicle{
		UUID:         v.UUID,
		LicensePlate: v.LicensePlate,
		State:        v.State,
		Archived:     v.Archived,
		Available:    v.Available,
		Year:         v.Year,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
}

func (v Vehicle) LogText() string {
	return fmt.Sprintf(`
	UUID ---------- %v
	Model --------- %v
	LicensePlate -- %v
	State --------- %v
	Archived ------ %v
	Available ----- %v
	Year ---------- %v
	CreatedAt ----- %v
	UpdatedAt ----- %v
	`,
		v.UUID,
		v.Model,
		v.LicensePlate,
		v.State,
		v.Archived,
		v.Available,
		v.Year,
		v.CreatedAt,
		v.UpdatedAt,
	)
}
