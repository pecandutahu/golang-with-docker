package db

import (
	"database/sql"
	"product/internal/domain"
	"product/internal/ports"
)

type ProductRepositoryDB struct {
	db *sql.DB
}

func NewProductRepositoryDB(db *sql.DB) ports.ProductRepository {
	return &ProductRepositoryDB{db: db}
}

func (r *ProductRepositoryDB) Save(product domain.Product) error {
	_, err := r.db.Exec("INSERT INTO products (product_code, product_name, stock) VALUES (?, ?, ?)", product.ProductCode, product.ProductName, product.Stock)
	return err
}

func (r *ProductRepositoryDB) FindByID(id int) (domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRow("SELECT product_id, product_code, product_name, stock FROM products WHERE product_id = ?", id).Scan(&product.ProductID, &product.ProductCode, &product.ProductName, &product.Stock)
	return product, err
}

func (r *ProductRepositoryDB) Update(product domain.Product) error {
	_, err := r.db.Exec("UPDATE products SET product_code = ?, product_name = ?, stock = ? WHERE product_id = ? ", product.ProductCode, product.ProductName, product.Stock, product.ProductID)
	return err
}

func (r *ProductRepositoryDB) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE product_id = ?", id)
	return err
}

func (r *ProductRepositoryDB) FindAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT product_id, product_code, product_name, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ProductID, &product.ProductCode, &product.ProductName, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
