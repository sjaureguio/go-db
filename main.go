package main

import (
	"log"

	"github.com/sjaureguio/go-db/pkg/product"
	"github.com/sjaureguio/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	// m := &product.Model{
	// 	Name:         "Curso de BD con go",
	// 	Price:        77,
	// 	Observations: "on fire",
	// }

	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v", m)

	// ms, err := serviceProduct.GetAll()

	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }
	// fmt.Println(ms)

	// m, err := serviceProduct.GetByID(2)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	log.Print("No hay un producto con ese id")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// m := &product.Model{
	// 	ID:    20,
	// 	Name:  "CURSO DE GO",
	// 	Price: 50,
	// }

	// err := serviceProduct.Update(m)

	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// m := &product.Model{
	// 	ID:    20,
	// 	Name:  "CURSO DE GO",
	// 	Price: 50,
	// }

	err := serviceProduct.Delete(3)

	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}

}
