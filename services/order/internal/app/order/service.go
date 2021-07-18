package order

import (
	"errors"
	"fmt"

	"github.com/mcholismalik/test-gank-ecom-go-api/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/dto"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/factory"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/model"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/repository"
	res "github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/response"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/trxmanager"
	"github.com/teris-io/shortid"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.OrderGetRequest) (*dto.OrderGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.OrderGetByIDRequest) (*dto.OrderGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.OrderCreateRequest) (*dto.OrderCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.OrderUpdateRequest) (*dto.OrderUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.OrderDeleteRequest) (*dto.OrderDeleteResponse, error)
}

type service struct {
	Repository          repository.Order
	OrderItemRepository repository.OrderItem
	Db                  *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.OrderRepository
	orderItemRepository := f.OrderItemRepository
	db := f.Db
	return &service{repository, orderItemRepository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.OrderGetRequest) (*dto.OrderGetResponse, error) {
	var result *dto.OrderGetResponse
	var datas *[]model.OrderEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.OrderFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.OrderGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.OrderGetByIDRequest) (*dto.OrderGetByIDResponse, error) {
	var result *dto.OrderGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.OrderGetByIDResponse{
		OrderEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.OrderCreateRequest) (*dto.OrderCreateResponse, error) {
	var result *dto.OrderCreateResponse
	var data *model.OrderEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// insert order
		var order model.OrderEntityModel
		order = model.OrderEntityModel{
			Context: ctx,
			UserId:  ctx.Auth.ID,
			OrderEntity: model.OrderEntity{
				Code: fmt.Sprintf("ORD-%s", shortid.MustGenerate()),
			},
		}

		data, err = s.Repository.Create(ctx, &order)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		// insert order items
		var orderItems []model.OrderItemEntityModel
		for _, v := range payload.OrderItems {
			orderItem := model.OrderItemEntityModel{
				Context:   ctx,
				OrderId:   data.ID,
				ProductId: v.ProductID,
				OrderItemEntity: model.OrderItemEntity{
					Qty: v.Qty,
				},
			}
			orderItems = append(orderItems, orderItem)
		}

		_, err = s.OrderItemRepository.BulkCreate(ctx, &orderItems)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		// get order
		data, err = s.Repository.FindByID(ctx, &data.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.OrderCreateResponse{
		OrderEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.OrderUpdateRequest) (*dto.OrderUpdateResponse, error) {
	var result *dto.OrderUpdateResponse
	var data *model.OrderEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// check order
		data, err = s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		// delete order items
		_, err = s.OrderItemRepository.DeleteByOrderID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		// insert order items
		var orderItems []model.OrderItemEntityModel
		for _, v := range payload.OrderItems {
			orderItem := model.OrderItemEntityModel{
				Context:   ctx,
				OrderId:   data.ID,
				ProductId: v.ProductID,
				OrderItemEntity: model.OrderItemEntity{
					Qty: v.Qty,
				},
			}
			orderItems = append(orderItems, orderItem)
		}

		_, err = s.OrderItemRepository.BulkCreate(ctx, &orderItems)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		// get new order
		data, err = s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.OrderUpdateResponse{
		OrderEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.OrderDeleteRequest) (*dto.OrderDeleteResponse, error) {
	var result *dto.OrderDeleteResponse
	var data *model.OrderEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// check order
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		// delete order items
		_, err = s.OrderItemRepository.DeleteByOrderID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		// delete order
		data, err = s.Repository.Delete(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.OrderDeleteResponse{
		OrderEntityModel: *data,
	}

	return result, nil
}
