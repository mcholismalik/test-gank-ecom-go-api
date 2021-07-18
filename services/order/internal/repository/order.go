package repository

import (
	"fmt"

	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"

	"gorm.io/gorm"
)

type Order interface {
	Find(ctx *abstraction.Context, m *model.OrderFilterModel, p *abstraction.Pagination) (*[]model.OrderEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error)
	Create(ctx *abstraction.Context, e *model.OrderEntityModel) (*model.OrderEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.OrderEntity) (*model.OrderEntityModel, error)
	Delete(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error)
}

type order struct {
	abstraction.Repository
}

func NewOrder(db *gorm.DB) *order {
	return &order{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *order) Find(ctx *abstraction.Context, m *model.OrderFilterModel, p *abstraction.Pagination) (*[]model.OrderEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.OrderEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.OrderEntityModel{})

	// filter
	query = r.Filter(ctx, query, *m)

	// sort
	if p.Sort != nil && p.SortBy != nil {
		sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
		query = query.Order(sort)
	}

	// pagination
	page := 1
	pageSize := 10
	if p.Page == nil {
		p.Page = &page
	}
	if p.PageSize == nil {
		p.PageSize = &pageSize
	}

	// association
	query = query.Preload("OrderItems.Product")

	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *order) FindByID(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.OrderEntityModel

	// association
	conn = conn.Preload("OrderItems.Product")

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *order) Create(ctx *abstraction.Context, e *model.OrderEntityModel) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *order) Update(ctx *abstraction.Context, id *int, e *model.OrderEntity) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.OrderEntityModel
	data.Context = ctx
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	data.OrderEntity = *e
	err = conn.Model(&data).Updates(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *order) Delete(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.OrderEntityModel
	data.Context = ctx
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Where("id = ?", id).Delete(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
