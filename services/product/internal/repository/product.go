package repository

import (
	"fmt"

	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"

	"gorm.io/gorm"
)

type Product interface {
	Find(ctx *abstraction.Context, m *model.ProductFilterModel, p *abstraction.Pagination) (*[]model.ProductEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.ProductEntityModel, error)
	Create(ctx *abstraction.Context, e *model.ProductEntity) (*model.ProductEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.ProductEntity) (*model.ProductEntityModel, error)
	Delete(ctx *abstraction.Context, id *int) (*model.ProductEntityModel, error)
}

type product struct {
	abstraction.Repository
}

func NewProduct(db *gorm.DB) *product {
	return &product{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *product) Find(ctx *abstraction.Context, m *model.ProductFilterModel, p *abstraction.Pagination) (*[]model.ProductEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.ProductEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.ProductEntityModel{})

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

func (r *product) FindByID(ctx *abstraction.Context, id *int) (*model.ProductEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.ProductEntityModel
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *product) Create(ctx *abstraction.Context, e *model.ProductEntity) (*model.ProductEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.ProductEntityModel
	data.Context = ctx
	data.ProductEntity = *e
	err := conn.Create(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *product) Update(ctx *abstraction.Context, id *int, e *model.ProductEntity) (*model.ProductEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.ProductEntityModel
	data.Context = ctx
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	data.ProductEntity = *e
	err = conn.Model(&data).Updates(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *product) Delete(ctx *abstraction.Context, id *int) (*model.ProductEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.ProductEntityModel
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
