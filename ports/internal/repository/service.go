package repository

import (
	"context"

	"ports/internal/models"
	"ports/log"
	"ports/utils"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// Service encapsulates usecase logic for albums.
type Service interface {
	Get(ctx context.Context, id string) (*Port, error)
	Query(ctx context.Context) ([]*Port, error)
	Create(ctx context.Context, input CreatePortRequest) (*Port, error)
	Update(ctx context.Context, id string, input UpdatePortRequest) (*Port, error)
	ParseJson(ctx context.Context) error
	Delete(ctx context.Context, id string) error
}

// Port represents the data about a Port.
type Port struct {
	models.Port
}

// CreatePortRequest represents a Port creation request.
type CreatePortRequest struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

// Validate validates the CreatePortRequest fields.
func (m CreatePortRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

// UpdatePortRequest represents an Port update request.
type UpdatePortRequest struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

// Validate validates the UpdatePortRequest fields.
func (m UpdatePortRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new album service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the album with the specified the album ID.
func (s service) Get(ctx context.Context, id string) (*Port, error) {
	port, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.repo.Get")
	}
	return &Port{port}, nil
}

// Create creates a new Port.
func (s service) Create(ctx context.Context, req CreatePortRequest) (*Port, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "req.Validate")
	}
	id := models.GenerateID()
	err := s.repo.Create(ctx, models.Port{
		Id:          id,
		Name:        req.Name,
		City:        req.City,
		Country:     req.Country,
		Alias:       req.Alias,
		Regions:     req.Regions,
		Coordinates: req.Coordinates,
		Province:    req.Province,
		Timezone:    req.Timezone,
		Unlocs:      req.Unlocs,
		Code:        req.Code,
	})
	if err != nil {
		return nil, errors.Wrap(err, "service.repo.Create")
	}
	return s.Get(ctx, id)
}

// Update updates the Port with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdatePortRequest) (*Port, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "req.Validate")
	}

	port, err := s.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.Get")
	}
	port.Name = req.Name
	port.City = req.City
	port.Country = req.Country
	port.Alias = req.Alias
	port.Regions = req.Regions
	port.Coordinates = req.Coordinates
	port.Province = req.Province
	port.Timezone = req.Timezone
	port.Unlocs = req.Unlocs
	port.Code = req.Code

	if err := s.repo.Update(ctx, port.Port); err != nil {
		return nil, errors.Wrap(err, "service.repo.Update")
	}
	return port, nil
}

// Delete deletes the Port with the specified ID.
func (s service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "service.repo.Delete")
	}
	return nil
}

// Query returns all Ports
func (s service) Query(ctx context.Context) ([]*Port, error) {
	p, err := s.repo.Query(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.repo.Query")
	}

	ports := []*Port{}
	for i := range p {
		ports = append(ports, &Port{p[i]})
	}
	return ports, nil
}

// Delete deletes the Port with the specified ID.
func (s service) ParseJson(ctx context.Context) error {
	p := models.PortsMap{}
	if err := utils.ReadJson("./ports.json", p); err != nil {
		return errors.Wrap(err, "utils.ReadJson")
	}
	ports := p.ToPorts()
	if err := s.repo.CreateBatch(ctx, ports); err != nil {
		return errors.Wrap(err, "service.repo.CreateBatch")
	}
	return nil
}
