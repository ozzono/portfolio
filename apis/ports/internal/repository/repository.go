package repository

import (
	"context"

	"ports/internal/models"
	"ports/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Repository encapsulates the logic to access ports from the data source.
type Repository interface {
	// Get returns the port with the specified port ID.
	Get(ctx context.Context, id string) (models.Port, error)

	// Query returns the list of all ports.
	Query(ctx context.Context) ([]models.Port, error)

	// Create saves a new port in the storage.
	Create(ctx context.Context, port models.Port) error

	// Create saves a new port in the storage.
	CreateBatch(ctx context.Context, ports []models.Port) error

	// Update updates the port with given ID in the storage.
	Update(ctx context.Context, port models.Port) error

	// Delete removes the port with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists ports in database
type repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new port repository
func NewRepository(db *gorm.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the port with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (models.Port, error) {
	var port models.Port
	result := r.db.WithContext(ctx).First(&port, "id = ?", id)
	return port, result.Error
}

// Create saves a new port record in the database.
func (r repository) Create(ctx context.Context, port models.Port) error {
	if result := r.db.WithContext(ctx).Create(&port); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateBatch saves a batch of new port records in the database.
func (r repository) CreateBatch(ctx context.Context, ports []models.Port) error {
	r.logger.Infof("creating %d batch registries", len(ports))
	result := r.db.WithContext(ctx).Table("all_ports").CreateInBatches(ports, len(ports))
	if result.Error != nil {
		return errors.Wrap(result.Error, "CreateBatch err -")
	}
	return nil
}

// Update saves the changes to a port in the database.
func (r repository) Update(ctx context.Context, port models.Port) error {
	if result := r.db.WithContext(ctx).Save(&port); result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a port with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	port, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	if result := r.db.WithContext(ctx).Delete(&port); result.Error != nil {
		return result.Error
	}
	return nil
}

// Query retrieves the port records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context) ([]models.Port, error) {
	var ports []models.Port
	result := r.db.WithContext(ctx).Find(&ports)
	return ports, result.Error
}
