package repository

import (
	"context"
	"strconv"

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
	UpSert(ctx context.Context, id string, input UpSertPortRequest) (*Port, error)
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

// UpSertPortRequest represents an Port upsert request.
type UpSertPortRequest struct {
	Id          *int          `json:"id"`
	Name        string        `json:"name"`
	RefName     string        `json:"ref_name"`
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
func (m UpSertPortRequest) Validate() error {
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

	uID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id; must be int")
	}

	port, err := s.repo.Get(ctx, int64(uID))
	if err != nil {
		return nil, errors.Wrap(err, "service.repo.Get")
	}
	return &Port{*port}, nil
}

// UpSert update the port if it already exists or create a new elsewhise
func (s service) UpSert(ctx context.Context, id string, req UpSertPortRequest) (*Port, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "req.Validate")
	}

	uID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id; must be int")
	}

	port := &models.Port{
		Id:          &uID,
		Name:        req.Name,
		RefName:     req.RefName,
		City:        req.City,
		Country:     req.Country,
		Alias:       req.Alias,
		Regions:     req.Regions,
		Coordinates: req.Coordinates,
		Province:    req.Province,
		Timezone:    req.Timezone,
		Unlocs:      req.Unlocs,
		Code:        req.Code,
	}

	if err := s.repo.UpSert(ctx, port); err != nil {
		return nil, errors.Wrap(err, "service.repo.Update")
	}
	return &Port{*port}, nil
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
		ports = append(ports, &Port{*p[i]})
	}
	return ports, nil
}

// Delete deletes the Port with the specified ID.
func (s service) ParseJson(ctx context.Context) error {
	p := map[string]models.Port{}
	if err := utils.ReadJson("ports.json", &p); err != nil {
		return errors.Wrap(err, "utils.ReadJson")
	}
	s.logger.Debug(p)
	if p != nil {
		ports := models.MapToPorts(p)
		if err := s.repo.CreateBatch(ctx, ports); err != nil {
			return errors.Wrap(err, "service.repo.CreateBatch")
		}
	}
	return nil
}
