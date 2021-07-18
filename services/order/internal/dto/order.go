package dto

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"
	res "github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/response"
)

// Get
type OrderGetRequest struct {
	abstraction.Pagination
	model.OrderFilterModel
}
type OrderGetResponse struct {
	Datas          []model.OrderEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type OrderGetResponseDoc struct {
	Body struct {
		Meta res.Meta                 `json:"meta"`
		Data []model.OrderEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type OrderGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type OrderGetByIDResponse struct {
	model.OrderEntityModel
}
type OrderGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data OrderGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type OrderCreateRequest struct {
	OrderItems []struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	} `json:"order_items"`
}
type OrderCreateResponse struct {
	model.OrderEntityModel
}
type OrderCreateResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data OrderCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type OrderUpdateRequest struct {
	ID         int `param:"id" validate:"required,numeric"`
	OrderItems []struct {
		ProductID int `json:"product_id"`
		Qty       int `json:"qty"`
	} `json:"order_items"`
}
type OrderUpdateResponse struct {
	model.OrderEntityModel
}
type OrderUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data OrderUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type OrderDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type OrderDeleteResponse struct {
	model.OrderEntityModel
}
type OrderDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data OrderDeleteResponse `json:"data"`
	} `json:"body"`
}
