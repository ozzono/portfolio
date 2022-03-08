package repository

import (
	"context"

	"ports/internal/models"
	"ports/pkg/log"

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

	// ParsedJson verifies if the json file was parsed and added to the database.
	ParsedJson(ctx context.Context) (bool, error)

	// SetParsed sets the json control as parsed avoiding multiple parsing.
	SetParsed(ctx context.Context) error

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
	var (
		res            models.Port
		strAlias       string
		strRegions     string
		strCoordinates string
		strUnlocs      string
	)

	if err := r.db.QueryRow(ctx, getPortByID, id).Scan(
		&res.Id,
		&res.Name,
		&res.RefName,
		&res.City,
		&res.Country,
		&strAlias,
		&strRegions,
		&strCoordinates,
		&res.Province,
		&res.Timezone,
		&strUnlocs,
		&res.Code,
	); err != nil {
		return nil, errors.Wrap(err, "repository.db.QueryRow")
	}

	models.FromString(&strAlias, &res.Alias)
	models.FromString(&strRegions, &res.Regions)
	models.FromString(&strCoordinates, &res.Coordinates)
	models.FromString(&strUnlocs, &res.Unlocs)

	return &res, nil
}

// Get reads the port with the specified ID from the database.
func (r repository) ParsedJson(ctx context.Context) (bool, error) {
	parsed := false

	if err := r.db.QueryRow(ctx, jsonParsed).Scan(
		&parsed,
	); errors.Is(pgx.ErrNoRows, err) {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "repository.db.QueryRow.Scan")
	}

	return parsed, nil
}

// UpSert update the port if it already exists or create a new elsewhise
func (r repository) UpSert(ctx context.Context, port *models.Port) error {
	if _, err := r.db.Exec(ctx, upsertPort,
		port.Id,
		port.Name,
		port.RefName,
		port.City,
		port.Country,
		models.ToString(port.Alias),
		models.ToString(port.Regions),
		models.ToString(port.Coordinates),
		port.Province,
		port.Timezone,
		models.ToString(port.Unlocs),
		port.Code,
	); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// SetParsed sets the json control as parsed avoiding multiple parsing.
func (r repository) SetParsed(ctx context.Context) error {
	if _, err := r.db.Exec(ctx, setParsed,
		true,
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
			port.Name,                         //name
			port.RefName,                      //ref_name
			port.City,                         //city
			port.Country,                      //country
			models.ToString(port.Alias),       //alias
			models.ToString(port.Regions),     //regions
			models.ToString(port.Coordinates), //coordinates
			port.Province,                     //province
			port.Timezone,                     //timezone
			models.ToString(port.Unlocs),      //unlocs
			port.Code,                         //code
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
	if _, err := r.db.Exec(ctx, delPort, id); err != nil {
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
		var (
			port           models.Port
			strAlias       string
			strRegions     string
			strCoordinates string
			strUnlocs      string
		)
		if err := rows.Scan(
			&port.Id,
			&port.Name,
			&port.RefName,
			&port.City,
			&port.Country,
			&strAlias,
			&strRegions,
			&strCoordinates,
			&port.Province,
			&port.Timezone,
			&strUnlocs,
			&port.Code,
		); err != nil {
			return nil, errors.Wrap(err, "Scan")
		}

		if strAlias != "" {
			models.FromString(&strAlias, &port.Alias)
		}

		if strRegions != "" {
			models.FromString(&strRegions, &port.Regions)
		}

		if strCoordinates != "" {
			models.FromString(&strCoordinates, &port.Coordinates)
		}

		if strUnlocs != "" {
			models.FromString(&strUnlocs, &port.Unlocs)
		}
		res = append(res, &port)
	}
	return res, err
}
