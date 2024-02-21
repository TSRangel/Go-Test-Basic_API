package product

import (
	"errors"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
	"time"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID = errors.New("id is invalid")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice = errors.New("invalid price")
)

type Product struct {
	ID tools.ID `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := tools.ParseID(p.ID.String());err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error){
	product := &Product{
		ID: tools.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	return product, err
}