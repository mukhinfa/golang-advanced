package product

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description,omitempty"`
	Images      []string `json:"images,omitempty"`
}

type GetProductResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Images      []string `json:"images,omitempty"`
}

type ListProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}
type UpdateProductRequest struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Images      *[]string `json:"images,omitempty"`
}
