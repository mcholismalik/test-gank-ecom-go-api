package model

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/date"

	"gorm.io/gorm"
)

type ProductEntity struct {
	Code string `json:"code" validate:"required" gorm:"index:idx_product_key,unique"`
	Name string `json:"name" validate:"required"`
	Tag  string `json:"tag" validate:"required"`
}

type ProductFilter struct {
	Name *string `query:"name" filter:"ILIKE"`
	Tag  *string `query:"tag" filter:"ILIKE"`
}

type ProductEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	ProductEntity

	// relations
	OrderItems []OrderItemEntityModel `json:"order_items" gorm:"foreignKey:ProductId"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type ProductFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	ProductFilter
}

func (ProductEntityModel) TableName() string {
	return "products"
}

func (m *ProductEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ProductEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
