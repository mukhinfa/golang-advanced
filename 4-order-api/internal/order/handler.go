package order

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/middleware"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/req"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/res"
)

type OrderService interface {
	CreateOrder(phone string, req CreateOrderRequest) (*Order, error)
	GetOrder(id uint) (*Order, error)
	GetUserOrders(phone string, limit, offset int) ([]Order, error)
}

type OrderHandlerDeps struct {
	OrderService   OrderService
	AuthMiddleware func(http.Handler) http.Handler
}
type orderHandler struct {
	OrderService
}

func NewHandler(r *http.ServeMux, deps OrderHandlerDeps) {
	handler := &orderHandler{
		OrderService: deps.OrderService,
	}
	authed := deps.AuthMiddleware
	r.Handle("POST /order", authed(handler.create()))
	r.Handle("GET /order/{id}", authed(handler.getByID()))
	r.Handle("GET /my-orders", authed(handler.getMy()))
}

func (h *orderHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			return
		}
		phone := middleware.GetPhone(r)
		if phone == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		createdOrder, err := h.OrderService.CreateOrder(*phone, *body) // to do
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusCreated, createdOrder)

	}
}
func (h *orderHandler) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}
		order, err := h.GetOrder(uint(idUint))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSON(w, http.StatusOK, order)
	}
}
func (h *orderHandler) getMy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		limit, err := getParam("limit", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if limit > 100 || limit <= 0 {
			http.Error(w, "incorrect limit", http.StatusBadRequest)
			return
		}
		offset, err := getParam("offset", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		phone := middleware.GetPhone(r)
		if phone == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		orders, err := h.GetUserOrders(r.Context().Value(middleware.PhoneCtxKey).(string), limit, offset) //to do
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusOK, orders)
	}
}

func getParam(name string, r *http.Request) (int, error) {
	res, err := strconv.Atoi(r.URL.Query().Get(name))
	if err != nil {
		return 0, fmt.Errorf("invalid %s", name)
	}
	return res, nil
}
