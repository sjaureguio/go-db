package product

import "time"

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Models slice of model
type Models []*Model

// Storage interface that must implemente a db dtorage
type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll(Models) error
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service have contain logic product
type Service struct {
	storage Storage
}

// NewService return a pinter of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
