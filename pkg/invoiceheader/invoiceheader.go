package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models slice of model
type Models []*Model

// Storage interface that must implemente a db dtorage
type Storage interface {
	Migrate() error
}

// Service have contain logic invoiceheader
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
