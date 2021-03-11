## Importación de paquete usados
``` go
import (
    "github.com/sjaureguio/go-db/pkg/invoiceitem"
    "github.com/sjaureguio/go-db/pkg/product"
    "github.com/sjaureguio/go-db/pkg/invoiceheader" 
)
```

## Migrar tabla de product

``` go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
    log.Fatalf("product.Migrate: %v", err)
}
```

## Migrar tabla de invoiceheader
``` go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
    log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

## Migrar tabla de invoiceitem
``` go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
    log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

