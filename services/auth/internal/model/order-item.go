package model

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/pkg/util/date"

	"gorm.io/gorm"
)

type OrderItemEntity struct {
	Code string `json:"code" validate:"required" gorm:"index:idx_order_item_key,unique"`
}

type OrderItemFilter struct {
}

type OrderItemEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	OrderItemEntity

	// relations
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type OrderItemFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	OrderItemFilter
}

func (OrderItemEntityModel) TableName() string {
	return "order_items"
}

func (m *OrderItemEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *OrderItemEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
