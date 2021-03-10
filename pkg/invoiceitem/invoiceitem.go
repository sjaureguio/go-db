package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Models slice of model
type Models []*Model

// Storage interface that must implemente a db dtorage
type Storage interface {
	Migrate() error
}

// Service have contain logic invoiceitem
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
