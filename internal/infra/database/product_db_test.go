package database

import (
	"fmt"
	"testing"

	"math/rand"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/product"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"

	"github.com/stretchr/testify/assert"
)

func TestCreatNewProduct(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnectionWithDataSource("product_test.db")
	as.Nil(err)
	defer db.Close()
	err = tools.CreateProductsTable(db)
	as.Nil(err)
	productConnection := NewProductConnection(db)
	newProduct, err := product.NewProduct("product 1", 10.0)
	as.Nil(err)
	err = productConnection.Create(newProduct)
	as.Nil(err)
	searchedProduct, err := productConnection.FindByID(newProduct.ID.String())
	as.Nil(err)
	as.Equal(newProduct.Price, searchedProduct.Price)
}

func TestFindAll(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnectionWithDataSource("products_test.db")
	as.Nil(err)
	defer db.Close()
	err = tools.CreateProductsTable(db)
	as.Nil(err)
	productConnection :=NewProductConnection(db)
	for i := 1; i <= 25; i ++ {
		newProduct, err := product.NewProduct(fmt.Sprintf("Product %d", i), (rand.Float64() * 100))
		as.Nil(err)
		err = productConnection.Create(newProduct)
		as.Nil(err)	
	}
	products, err := productConnection.FindAll(1, 10, "asc")
	as.Nil(err)
	as.Len(products, 10)
	as.Equal("Product 1", products[0].Name)
	as.Equal("Product 10", products[9].Name)
}

func TestFindByID(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnectionWithDataSource("product_test.db")
	as.Nil(err)
	defer db.Close()
	err = tools.CreateProductsTable(db)
	as.Nil(err)
	productConnection := NewProductConnection(db)
	newProduct, err := product.NewProduct("Product 1", 10.0)
	as.Nil(err)
	err = productConnection.Create(newProduct)
	as.Nil(err)
	searchedProduct, err := productConnection.FindByID(newProduct.ID.String())
	as.Nil(err)
	as.Equal(newProduct.ID, searchedProduct.ID)
	as.Equal(newProduct.Name, searchedProduct.Name)
	as.Equal(newProduct.Price, searchedProduct.Price)
}

func TestUpdateProduct(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnectionWithDataSource("product_test.db")
	as.Nil(err)
	defer db.Close()
	err = tools.CreateProductsTable(db)
	as.Nil(err)
	productConnection := NewProductConnection(db)
	newProduct, err := product.NewProduct("Product 1", 10.0)
	as.Nil(err)
	err = productConnection.Create(newProduct)
	as.Nil(err)
	newProduct.Price = 20.0
	searchedProduct, err := productConnection.FindByID(newProduct.ID.String())
	as.Nil(err)
	err = productConnection.Update(newProduct)
	as.Nil(err)
	as.NotEqual(newProduct.Price, searchedProduct.Price)
}

func TestDeleteProduct(t *testing.T) {
	as := assert.New(t)
	db, err := tools.DBConnectionWithDataSource("product_test.db")
	as.Nil(err)
	defer db.Close()
	err = tools.CreateProductsTable(db)
	as.Nil(err)
	productConnection := NewProductConnection(db)
	newProduct, err := product.NewProduct("Product 1", 10.0)
	as.Nil(err)
	err = productConnection.Create(newProduct)
	as.Nil(err)
	searchedProductBeforeDeletion, err := productConnection.FindByID(newProduct.ID.String())
	as.Nil(err)
	as.NotEmpty(searchedProductBeforeDeletion)
	err = productConnection.Delete(newProduct.ID.String())
	as.Nil(err)
	searchedProductAfterDeletion, err := productConnection.FindByID(newProduct.ID.String())
	as.NotNil(err)
	as.Empty(searchedProductAfterDeletion)
}