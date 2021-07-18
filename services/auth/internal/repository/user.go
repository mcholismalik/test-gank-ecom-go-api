package repository

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/abstraction"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntity) (*model.UserEntityModel, error)
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, e *model.UserEntity) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	data.UserEntity = *e
	err := conn.Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
