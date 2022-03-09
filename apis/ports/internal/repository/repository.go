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

	// GetByCode returns the port with the specified port code.
	GetByCode(ctx context.Context, code string) (*models.Port, error)

	// Query returns the list of all ports.
	Query(ctx context.Context) ([]*models.Port, error)

	// UpSert update the port if it already exists or create a new elsewhise
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
	r.logger.Infof("fetching port from id: %d", id)
	var (
		res            models.Port
		strAlias       string
		strRegions     string
		strCoordinates string
		strUnlocs      string
	)

	if err := r.db.QueryRow(ctx, getPortByID, id).Scan(
		&res.Id,         // 01
		&res.Name,       // 02
		&res.RefName,    // 03
		&res.City,       // 04
		&res.Country,    // 05
		&strAlias,       // 06
		&strRegions,     // 07
		&strCoordinates, // 08
		&res.Province,   // 09
		&res.Timezone,   // 10
		&strUnlocs,      // 11
		&res.Code,       // 12
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
func (r repository) GetByCode(ctx context.Context, code string) (*models.Port, error) {
	r.logger.Infof("fetching port from id: %s", code)
	var (
		res            models.Port
		strAlias       string
		strRegions     string
		strCoordinates string
		strUnlocs      string
	)

	if err := r.db.QueryRow(ctx, getPortByCode, code).Scan(
		&res.Id,         // 01
		&res.Name,       // 02
		&res.RefName,    // 03
		&res.City,       // 04
		&res.Country,    // 05
		&strAlias,       // 06
		&strRegions,     // 07
		&strCoordinates, // 08
		&res.Province,   // 09
		&res.Timezone,   // 10
		&strUnlocs,      // 11
		&res.Code,       // 12
	); err != nil {
		return nil, errors.Wrap(err, "repository.db.QueryRow")
	}

	models.FromString(&strAlias, &res.Alias)
	models.FromString(&strRegions, &res.Regions)
	models.FromString(&strCoordinates, &res.Coordinates)
	models.FromString(&strUnlocs, &res.Unlocs)

	return &res, nil
}

// Parsed data from json file into database
func (r repository) ParsedJson(ctx context.Context) (bool, error) {
	r.logger.Infof("parsing json data")
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
	r.logger.Infof("upserting %s port", port.Name)
	if _, err := r.db.Exec(ctx, upsertPort,
		port.Name,                         // 01
		port.RefName,                      // 02
		port.City,                         // 03
		port.Country,                      // 04
		models.ToString(port.Alias),       // 05
		models.ToString(port.Regions),     // 06
		models.ToString(port.Coordinates), // 07
		port.Province,                     // 08
		port.Timezone,                     // 09
		models.ToString(port.Unlocs),      // 10
		port.Code,                         // 11
	); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// SetParsed sets the json control as parsed avoiding multiple parsing.
func (r repository) SetParsed(ctx context.Context) error {
	r.logger.Info("setting json control as parsed")
	if _, err := r.db.Exec(ctx, setParsed,
		true,
	); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// CreateBatch saves a batch of new port records in the database.
func (r repository) CreateBatch(ctx context.Context, ports []*models.Port) error {
	r.logger.Infof("batch inserting %d ports", len(ports))

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
			upsertPort,
			port.Name,                         // 01
			port.RefName,                      // 02
			port.City,                         // 03
			port.Country,                      // 04
			models.ToString(port.Alias),       // 05
			models.ToString(port.Regions),     // 06
			models.ToString(port.Coordinates), // 07
			port.Province,                     // 08
			port.Timezone,                     // 09
			models.ToString(port.Unlocs),      // 10
			port.Code,                         // 11
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
	r.logger.Infof("deleting port id: %s", id)
	if _, err := r.db.Exec(ctx, delPort, id); err != nil {
		return errors.Wrap(err, "repository.db.Exec")
	}
	return nil
}

// Query retrieves the port records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context) ([]*models.Port, error) {
	r.logger.Info("fetching all ports")
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
	r.logger.Infof("returning %d ports", len(res))
	return res, err
}
