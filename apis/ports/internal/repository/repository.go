package repository

import (
	"context"

	"ports/internal/models"
	"ports/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Repository encapsulates the logic to access ports from the data source.
type Repository interface {
	// Get returns the port with the specified port ID.
	Get(ctx context.Context, id int64) (*models.Port, error)

	// Query returns the list of all ports.
	Query(ctx context.Context) ([]*models.Port, error)

	// Create saves a new port in the storage.
	UpSert(ctx context.Context, port *models.Port) error

	// Create saves a new port in the storage.
	CreateBatch(ctx context.Context, ports []*models.Port) error

	// Delete removes the port with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists ports in database
type repository struct {
	db     *pgxpool.Pool
	logger log.Logger
}

// NewRepository creates a new port repository
func NewRepository(db *pgxpool.Pool, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the port with the specified ID from the database.
func (r repository) Get(ctx context.Context, id int64) (*models.Port, error) {
	var res models.Port

	if err := r.db.QueryRow(ctx, getPortByID, id).Scan(
		&res.Name,
		&res.RefName,
		&res.City,
		&res.Country,
		&res.Alias,
		&res.Regions,
		&res.Coordinates,
		&res.Province,
		&res.Timezone,
		&res.Unlocs,
		&res.Code,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &res, nil
}

// UpSert update the port if it already exists or create a new elsewhise
func (r repository) UpSert(ctx context.Context, port *models.Port) error {
	if _, err := r.db.Exec(ctx, upsertPort,
		port.Name,
		port.RefName,
		port.City,
		port.Country,
		port.Alias,
		port.Regions,
		port.Coordinates,
		port.Province,
		port.Timezone,
		port.Unlocs,
		port.Code,
	); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// CreateBatch saves a batch of new port records in the database.
func (r repository) CreateBatch(ctx context.Context, ports []*models.Port) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "repository.db.Begin")
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	batch := &pgx.Batch{}
	for _, port := range ports {
		batch.Queue(
			createPort,
			port.Name,
			port.RefName,
			port.City,
			port.Country,
			port.Alias,
			port.Regions,
			port.Coordinates,
			port.Province,
			port.Timezone,
			port.Unlocs,
			port.Code,
		)
	}

	batchResults := tx.SendBatch(ctx, batch)

	if _, err := batchResults.Exec(); err != nil {
		return errors.Wrap(err, "batchResults.Exec")
	}

	err = batchResults.Close()
	if err != nil {
		return errors.Wrap(err, "batchResults.Close")
	}

	return nil
}

// Delete deletes a port with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	if _, err := r.db.Exec(ctx, upsertPort, id); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// Query retrieves the port records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context) ([]*models.Port, error) {
	rows, err := r.db.Query(ctx, allPorts)
	if err != nil {
		return nil, errors.Wrap(err, "repository.db.Query")
	}
	defer rows.Close()
	res := []*models.Port{}

	for rows.Next() {
		var port models.Port
		if err := rows.Scan(
			&port.Id,
			&port.Name,
			&port.RefName,
			&port.City,
			&port.Country,
			&port.Alias,
			&port.Regions,
			&port.Coordinates,
			&port.Province,
			&port.Timezone,
			&port.Unlocs,
			&port.Code,
		); err != nil {
			return nil, errors.Wrap(err, "Scan")
		}
		res = append(res, &port)
	}
	return res, err
}
