package model

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/date"

	"gorm.io/gorm"
)

type OrderItemEntity struct {
	Qty int `json:"qty" validate:"required"`
}

type OrderItemFilter struct {
}

type OrderItemEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	OrderItemEntity

	// relations
	OrderId   int                `json:"order_id"`
	ProductId int                `json:"product_id"`
	Product   ProductEntityModel `json:"product" gorm:"foreignKey:ProductId"`

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
