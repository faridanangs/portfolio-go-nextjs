package apps

import (
	"fmt"
	"os"
	"portfolio/helpers"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	helpers.PanicIfError(err, "Error at godotenv.load /apps")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	dns := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s TimeZone=%s sslmode=disable", dbUser, dbName, dbPass, dbHost, dbPort, dbTimeZone)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	helpers.PanicIfError(err, "Error at connect to database /apps")

	sqlDB, err := db.DB()
	helpers.PanicIfError(err, "Error at config sqlDB /apps")

	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)

	return db
}
