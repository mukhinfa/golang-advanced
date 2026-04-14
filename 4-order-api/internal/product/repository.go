package product

import (
	"fmt"

	"gorm.io/gorm/clause"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

type productRepository struct {
	Database *db.Db
}

func NewRepository(db *db.Db) *productRepository {
	return &productRepository{
		Database: db,
	}
}

func (r *productRepository) Create(product *Product) (*Product, error) {
	result := r.Database.Create(product)
	if result.Error != nil {
		return nil, fmt.Errorf("product.Repository.Create failed to create product: %w", result.Error)
	}
	return product, nil
}

func (r *productRepository) GetByID(id uint) (*Product, error) {
	var product Product
	result := r.Database.First(&product, id)
	if result.Error != nil {
		return nil, fmt.Errorf("product.Repository.GetByID failed to get product with ID %d: %w", id, result.Error)
	}
	return &product, nil
}

func (r *productRepository) GetByName(name string) (*Product, error) {
	var product Product
	result := r.Database.First(&product, "name = ?", name)
	if result.Error != nil {
		return nil, fmt.Errorf("product.Repository.GetByName failed to get product with name %s: %w", name, result.Error)
	}
	return &product, nil
}

func (r *productRepository) Update(product *Product) (*Product, error) {
	result := r.Database.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, fmt.Errorf("product.Repository.Update failed to update product with ID %d: %w", product.ID, result.Error)
	}
	return product, nil
}

func (r *productRepository) Delete(id uint) error {
	result := r.Database.Delete(&Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("product.Repository.Delete failed to delete product with ID %d: %w", id, result.Error)
	}
	return nil
}

func (r *productRepository) List() ([]Product, error) {
	var products []Product
	result := r.Database.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("product.Repository.List failed to list products: %w", result.Error)
	}
	return products, nil
}
