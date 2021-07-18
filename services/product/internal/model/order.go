package model

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/date"

	"gorm.io/gorm"
)

type OrderEntity struct {
	Code string `json:"code" validate:"required" gorm:"index:idx_order_key,unique"`
}

type OrderFilter struct {
	UserId *int `query:"user_id"`
}

type OrderEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	OrderEntity

	// relations
	UserId     int                    `json:"user_id"`
	OrderItems []OrderItemEntityModel `json:"order_items" gorm:"foreignKey:OrderId"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type OrderFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	OrderFilter
}

func (OrderEntityModel) TableName() string {
	return "orders"
}

func (m *OrderEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *OrderEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
