package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/req"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/res"
)

var ErrNotFound = errors.New("not found")

type ServiceInterface interface {
	CreateProduct(req CreateProductRequest) (*Product, error)
	GetProduct(id uint) (*Product, error)
	UpdateProduct(input UpdateProductInput) (*Product, error)
	DeleteProduct(id uint) error
	ListProducts() ([]GetProductResponse, error)
}

type ProductHandlerDeps struct {
	ServiceInterface
	AuthMiddleware func(http.Handler) http.Handler
}

type productHandler struct {
	ServiceInterface
}

func NewProductHandler(r *http.ServeMux, deps ProductHandlerDeps) {
	handler := &productHandler{
		ServiceInterface: deps.ServiceInterface,
	}
	authed := deps.AuthMiddleware
	r.Handle("POST /products", authed(handler.createProduct()))
	r.Handle("GET /products/{id}", authed(handler.getProduct()))
	r.Handle("PUT /products/{id}", authed(handler.updateProduct()))
	r.Handle("DELETE /products/{id}", authed(handler.deleteProduct()))
	r.Handle("GET /products", authed(handler.listProducts()))
}

func (h *productHandler) createProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateProductRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdProduct, err := h.CreateProduct(*body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusCreated, createdProduct)
	}
}

func (h *productHandler) getProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		product, err := h.GetProduct(uint(idUint))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product == nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		res.JSON(w, http.StatusOK, product)
	}
}
func (h *productHandler) updateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[UpdateProductRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedProduct, err := h.UpdateProduct(UpdateProductInput{
			ID:          uint(idUint),
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusOK, updatedProduct)
	}
}
func (h *productHandler) deleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		err = h.DeleteProduct(uint(idUint))
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusOK, nil)
	}
}
func (h *productHandler) listProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.ListProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusOK, products)
	}
}
