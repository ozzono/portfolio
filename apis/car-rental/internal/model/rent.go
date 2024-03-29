package model

import (
	"car-rental/utils"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var (
	RegularFee     float64 = 1
	OvertimeFee    float64 = 1.5
	InvalidUUIDErr         = errors.New("invalid uuid")
)

type Rent struct {
	UUID        utils.UUID `json:"uuid,-" gorm:"primaryKey"`
	UserUUID    utils.UUID `gorm:"foreignKey"`
	VehicleUUID utils.UUID `gorm:"foreignKey"`
	Cost        float64
	Refundable  float64
	CreatedAt   time.Time
	PickUpAt    *time.Time
	PickedAt    *time.Time
	DropOffAt   *time.Time
	DroppedAt   *time.Time
	CanceledAt  *time.Time

	// scheduled - when PickUpAt != nil
	//
	// active - when PickedAt != nil
	//
	// delayed - when time.Now >  DropOffAt
	//
	// inactive - when DroppedAt != nil
	//
	// canceled - when CanceledAt !=nil
	Status string
}

func (r Rent) Copy() *Rent {
	return &Rent{
		UUID:        r.UUID,
		UserUUID:    r.UserUUID,
		VehicleUUID: r.VehicleUUID,
		Cost:        r.Cost,
		CreatedAt:   r.CreatedAt,
		PickUpAt:    r.PickUpAt,
		PickedAt:    r.PickedAt,
		DropOffAt:   r.DropOffAt,
		DroppedAt:   r.DroppedAt,
	}
}

func (r Rent) LogText() string {
	return fmt.Sprintf(`
	UUID --------- %v        
	UserUUID ----- %v    
	VehicleUUID -- %v 
	Status ------- %v      
	Cost --------- %v        
	Refundable --- %v        
	CreatedAt ---- %v   
	PickUpAt ----- %v    
	DropOffAt ---- %v   
	DroppedAt ---- %v   
	CanceledAt --- %v   
`,
		r.UUID,
		r.UserUUID,
		r.VehicleUUID,
		r.Status,
		r.Cost,
		r.Refundable,
		r.CreatedAt,
		r.PickUpAt,
		r.DropOffAt,
		r.DroppedAt,
		r.CanceledAt,
	)
}

func (r *Rent) Valid() error {
	if err := r.UserUUID.Valid(); err != nil {
		return errors.Wrap(err, "invalid UserUUID")
	}
	if err := r.VehicleUUID.Valid(); err != nil {
		return errors.Wrap(err, "invalid VehicleUUID")
	}
	return nil
}
