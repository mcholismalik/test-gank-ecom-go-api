package product

import (
	"errors"

	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/dto"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/factory"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/repository"
	res "github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/response"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.ProductGetRequest) (*dto.ProductGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.ProductGetByIDRequest) (*dto.ProductGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.ProductCreateRequest) (*dto.ProductCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.ProductUpdateRequest) (*dto.ProductUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.ProductDeleteRequest) (*dto.ProductDeleteResponse, error)
}

type service struct {
	Repository repository.Product
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ProductRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.ProductGetRequest) (*dto.ProductGetResponse, error) {
	var result *dto.ProductGetResponse
	var datas *[]model.ProductEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.ProductFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ProductGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.ProductGetByIDRequest) (*dto.ProductGetByIDResponse, error) {
	var result *dto.ProductGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ProductGetByIDResponse{
		ProductEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ProductCreateRequest) (*dto.ProductCreateResponse, error) {
	var result *dto.ProductCreateResponse
	var data *model.ProductEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.ProductEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ProductCreateResponse{
		ProductEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.ProductUpdateRequest) (*dto.ProductUpdateResponse, error) {
	var result *dto.ProductUpdateResponse
	var data *model.ProductEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.ProductEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ProductUpdateResponse{
		ProductEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.ProductDeleteRequest) (*dto.ProductDeleteResponse, error) {
	var result *dto.ProductDeleteResponse
	var data *model.ProductEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data, err = s.Repository.Delete(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ProductDeleteResponse{
		ProductEntityModel: *data,
	}

	return result, nil
}
