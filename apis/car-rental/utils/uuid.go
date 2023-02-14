package utils

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInvalidUUID = errors.New("invalid uuid")
)

type UUID struct {
	uuid.UUID
}

func (uuid UUID) Valid() error {
	if uuid.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}
	if err := uuid.Parse(uuid.String()); err != nil {
		return errors.Wrap(err, "uuid.Parse")
	}
	return nil
}
