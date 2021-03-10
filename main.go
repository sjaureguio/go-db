package main

import (
	"log"

	"github.com/sjaureguio/go-db/pkg/product"
	"github.com/sjaureguio/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	storageInvoiceHeader := storage.NewPsqlInvoiceHader(storage.Pool())
	serviceInvoiceHeader := product.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
	// product.NewService()
}
