package main

import "github.com/sjaureguio/go-db/storage"

func main() {
	storage.NewPostgresDB()

	// storageProduct := storage.NewPsqlProduct(storage.Pool())

	// product.NewService()
}
