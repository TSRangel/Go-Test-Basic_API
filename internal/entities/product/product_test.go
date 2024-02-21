package product

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNewProduct(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("Product 1", 10.0)
	as.Nil(err)
	as.Equal("Product 1", product.Name)
	as.Equal(10, product.Price)
}

func TestAllErrors(t *testing.T) {
	as := assert.New(t)
	type product struct {
		name string
		price float64
		err error
	}
	
	var products = []product{
		{name: "", price: 10.0, err: ErrNameIsRequired},
		{name: "Product 1", price: 0.0, err: ErrPriceIsRequired},
		{name: "Product 1", price: -10.0, err: ErrInvalidPrice},
	}
	for _, product := range products {
		newProduct, err := NewProduct(product.name, product.price)
		as.Error(err, product.err)
		as.NotNil(newProduct)
	}
}