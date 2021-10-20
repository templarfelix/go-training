package main

import (
	"99-project/account"
	"99-project/transaction"
	"99-project/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	gormlogger "gorm.io/gorm/logger"
)

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

	// repository
	userRepo := user.NewGormRepository(dbConnection)
	accountRepo := account.NewGormRepository(dbConnection)
	transactionRepo := transaction.NewGormRepository(dbConnection)

	// service
	userService := user.NewService(userRepo)
	accountService := account.NewService(accountRepo)
	transactionService := transaction.NewService(transactionRepo)

	// server web
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// register handlers
	user.RegisterUserHandlers(e, userService)
	account.RegisterAccountHandlers(e, accountService)
	transaction.RegisterTransactionHandlers(e, transactionService)

	// Start server
	e.Logger.Fatal(e.Start(":9090"))

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
