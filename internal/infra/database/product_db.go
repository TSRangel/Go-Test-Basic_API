package database

import (
	"database/sql"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/product"
)

type ProductConnection struct {
	DB *sql.DB
}

func NewProductConnection(db *sql.DB) *ProductConnection{
	return &ProductConnection{DB: db}
}

func (p *ProductConnection) Create(product *product.Product) error {
	stmt, err := p.DB.Prepare("insert into products (id, name, price, created_at) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name,product.Price, product.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductConnection) FindAll(page, limit int,  sort string) ([]product.Product, error) {
	if sort == "" || sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var products []product.Product

	if page != 0 && limit != 0 {
		stmt, err := p.DB.Prepare("select id, name, price, created_at from products order by created_at " + sort + " limit ? offset ?")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		rows, err := stmt.Query(limit, (page -1))
		if err != nil {
			return nil, err
		}
		
		for rows.Next() {
			var product product.Product

			err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	} else {
		stmt, err := p.DB.Prepare("select id, name, price, created_at from products order by created_at " + sort)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		rows, err := stmt.Query()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var product product.Product
			err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}
	return products, nil
}

func (p *ProductConnection) FindByID(id string) (*product.Product, error) {
	stmt, err := p.DB.Prepare("select id, name, price, created_at from products where id = ?")
	if  err != nil {
		return nil, err
	}
	defer stmt.Close()
	var foundedProduct product.Product
	err = stmt.QueryRow(id).Scan(&foundedProduct.ID, &foundedProduct.Name, &foundedProduct.Price, &foundedProduct.CreatedAt)
	if err != nil  {
		return nil, err
	}
	return &foundedProduct, nil
}

func (p *ProductConnection) Update(product *product.Product) error {
	stmt, err := p.DB.Prepare("update products set name = ?, price = ?, created_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.CreatedAt, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductConnection) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	stmt, err := p.DB.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}