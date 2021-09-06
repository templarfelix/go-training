package main

import (
	"99-project/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	gormlogger "gorm.io/gorm/logger"
)

type dbops struct {
	db *gorm.DB
}

func main() {

	// logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	zap.L().Info("init application")

	// db
	gormLogger := zapgorm2.New(zap.L()).LogMode(gormlogger.Info)
	dsn := "host=localhost user=root dbname=project port=26257 sslmode=disable TimeZone=GMT"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	checkErr(err)

	db := dbops{dbConnection}

	// repository
	userRepo := user.NewGormRepository(db.db)
	userService := user.NewService(userRepo)

	// server web
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// register handlers
	user.RegisterUserHandlers(e, userService)

	// Start server
	e.Logger.Fatal(e.Start(":9090"))

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
