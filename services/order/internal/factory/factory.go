package factory

import (
	"github.com/mcholismalik/test-gank-ecom-go-api/database"

	"github.com/mcholismalik/test-gank-ecom-go-api/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db                  *gorm.DB
	OrderRepository     repository.Order
	OrderItemRepository repository.OrderItem
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("SAMPLE1")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.OrderRepository = repository.NewOrder(f.Db)
	f.OrderItemRepository = repository.NewOrderItem(f.Db)
}
