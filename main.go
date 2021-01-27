package main

import (
	"log"
	"net/http"

	conf "innovasive/go-clean-template/config"
	localMidW "innovasive/go-clean-template/middleware"
	userHandler "innovasive/go-clean-template/service/user/http"
	userRepository "innovasive/go-clean-template/service/user/repository"
	userUsecase "innovasive/go-clean-template/service/user/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func sqlDB() *sqlx.DB {
	var connstr = conf.GetEnv("PSQL_DATABASE_URL", "postgres://127.0.0.1:5432")
	db, err := conf.NewPsqlConnection(connstr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	psqlDB := sqlDB()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		sqlDB()

		return c.String(http.StatusOK, "Hello, World!")
	})
	initMiddleware := localMidW.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userRepo := userRepository.NewPsqlUserRepository(psqlDB)

	userUC := userUsecase.NewUserUsecase(userRepo)

	userHandler.NewUserHandler(e, initMiddleware, userUC)

	e.Logger.Fatal(e.Start(":7000"))
}
