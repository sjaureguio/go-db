package storage

import (
	"database/sql"
	"fmt"

	"github.com/sjaureguio/go-db/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHaeder = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoices_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES ($1) RETURNING id, created_at`
)

// PsqlInvoiceHader used for work with postgres - invoiceheader
type PsqlInvoiceHader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of PsqlInvoiceHader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHader {
	return &PsqlInvoiceHader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHaeder)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de invoiceHeader ejecutada correctamente")

	return nil
}

// CreateTx implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
