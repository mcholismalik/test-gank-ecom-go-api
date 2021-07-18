package main

import (
	"os"

	db "github.com/mcholismalik/test-gank-ecom-go-api/services/auth/database"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/database/migration"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/factory"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/http"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/internal/middleware"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/pkg/elasticsearch"
	"github.com/mcholismalik/test-gank-ecom-go-api/services/auth/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title gank-ecommerce-auth
// @version 0.0.1
// @description This is a doc for gank-ecommerce-auth.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
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
