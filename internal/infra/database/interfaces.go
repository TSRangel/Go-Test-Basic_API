package database

import (
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/user"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/product"
)

type UserInterface interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}

type ProductInterface interface {
	Create(product *product.Product) error
	FindAll(page, limit int, sort string) ([]product.Product, error)
	FindByID(id string) (*product.Product, error)
	Update(product *product.Product) error
	Delete(id string) error
}