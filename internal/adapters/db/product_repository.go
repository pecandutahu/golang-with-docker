package db

import (
	"product/internal/domain"
	"product/internal/ports"

	"gorm.io/gorm"
)

type ProductRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	return &ProductRepositoryDB{db: db}
}

func (r *ProductRepositoryDB) Save(product domain.Product) error {
	return r.db.Create(&product).Error
	// _, err := r.db.Exec("INSERT INTO products (product_code, product_name, stock) VALUES (?, ?, ?)", product.ProductCode, product.ProductName, product.Stock)
	// return err
}

func (r *ProductRepositoryDB) FindByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return product, err
	// err := r.db.QueryRow("SELECT product_id, product_code, product_name, stock FROM products WHERE product_id = ?", id).Scan(&product.ProductID, &product.ProductCode, &product.ProductName, &product.Stock)
	// return product, err
}

func (r *ProductRepositoryDB) Update(product domain.Product) error {
	return r.db.Save(&product).Error
	// _, err := r.db.Exec("UPDATE products SET product_code = ?, product_name = ?, stock = ? WHERE product_id = ? ", product.ProductCode, product.ProductName, product.Stock, product.ProductID)
	// return err
}

func (r *ProductRepositoryDB) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
	// _, err := r.db.Exec("DELETE FROM products WHERE product_id = ?", id)
	// return err
}

func (r *ProductRepositoryDB) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
	// rows, err := r.db.Query("SELECT product_id, product_code, product_name, stock FROM products")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var products []domain.Product

	// for rows.Next() {
	// 	var product domain.Product
	// 	if err := rows.Scan(&product.ProductID, &product.ProductCode, &product.ProductName, &product.Stock); err != nil {
	// 		return nil, err
	// 	}
	// 	products = append(products, product)
	// }
	// return products, nil
}
