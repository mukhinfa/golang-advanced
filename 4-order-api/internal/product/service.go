package product

import (
	"fmt"
	"log"
)

type ProductRepositoryInterface interface {
	Create(product *Product) (*Product, error)
	GetByID(id uint) (*Product, error)
	GetByName(name string) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id uint) error
	List() ([]Product, error)
}

type UpdateProductInput struct {
	ID          uint
	Name        *string
	Description *string
	Images      *[]string
}

type Service struct {
	ProductRepositoryInterface
}

func NewService(repo ProductRepositoryInterface) *Service {
	return &Service{
		ProductRepositoryInterface: repo,
	}
}

func (s *Service) CreateProduct(req CreateProductRequest) (*Product, error) {
	var product *Product
	product = NewProduct(req.Name, req.Description, req.Images)

	existingProduct, _ := s.GetByName(product.Name)
	if existingProduct != nil {
		return nil, fmt.Errorf("product with name %s already exists", product.Name)
	}

	createdProduct, err := s.Create(product)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, err
	}
	return createdProduct, nil
}

func (s *Service) GetProduct(id uint) (*Product, error) {
	product, err := s.GetByID(id)
	if err != nil {
		log.Printf("Error getting product with ID %d: %v", id, err)
		return nil, err
	}
	return product, nil
}

func (s *Service) UpdateProduct(input UpdateProductInput) (*Product, error) {
	product, err := s.GetByID(input.ID)
	if err != nil {
		log.Printf("Error getting product with ID %d: %v", input.ID, err)
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("product with ID %d not found", input.ID)
	}

	// Update the product fields
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Images != nil {
		product.Images = *input.Images
	}

	updatedProduct, err := s.Update(product)
	if err != nil {
		log.Printf("Error updating product with ID %d: %v", input.ID, err)
		return nil, err
	}
	return updatedProduct, nil
}

func (s *Service) DeleteProduct(id uint) error {
	_, err := s.GetByID(id)
	if err != nil {
		return ErrNotFound
	}
	if err := s.Delete(id); err != nil {
		log.Printf("Error deleting product with ID %d: %v", id, err)
		return err
	}
	return nil
}

func (s *Service) ListProducts() ([]GetProductResponse, error) {
	products, err := s.List()
	if err != nil {
		log.Printf("Error listing products: %v", err)
		return nil, err
	}

	var responses []GetProductResponse
	for _, product := range products {
		responses = append(responses, GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Images:      product.Images,
		})
	}
	return responses, nil
}
