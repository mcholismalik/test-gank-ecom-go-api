package main

import (
	"os"

	db "github.com/mcholismalik/test-gank-ecom-go-api/database"
	"github.com/mcholismalik/test-gank-ecom-go-api/database/migration"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/factory"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/http"
	"github.com/mcholismalik/test-gank-ecom-go-api/internal/middleware"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/elasticsearch"
	"github.com/mcholismalik/test-gank-ecom-go-api/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title gank-ecommerce-order
// @version 0.0.1
// @description This is a doc for gank-ecommerce-order.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3032
// @BasePath /
func main() {
	var PORT = os.Getenv("PORT")

	db.Init()
	migration.Init()
	elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
