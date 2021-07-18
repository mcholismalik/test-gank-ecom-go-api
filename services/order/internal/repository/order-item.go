package repository

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"

	"gorm.io/gorm"
)

type OrderItem interface {
	BulkCreate(ctx *abstraction.Context, e *[]model.OrderItemEntityModel) (*[]model.OrderItemEntityModel, error)
	DeleteByOrderID(ctx *abstraction.Context, orderID *int) (*[]model.OrderItemEntityModel, error)
}

type orderItem struct {
	abstraction.Repository
}

func NewOrderItem(db *gorm.DB) *orderItem {
	return &orderItem{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *orderItem) BulkCreate(ctx *abstraction.Context, e *[]model.OrderItemEntityModel) (*[]model.OrderItemEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(e).First(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *orderItem) DeleteByOrderID(ctx *abstraction.Context, orderID *int) (*[]model.OrderItemEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data []model.OrderItemEntityModel
	err := conn.Where("order_id = ?", orderID).Find(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Where("order_id = ?", orderID).Delete(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
