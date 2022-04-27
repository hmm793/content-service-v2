package config

import (
	"content-service-v2/app/domain/entity"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres() (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`,
		os.Getenv("DB_HOST_DEV"),
		os.Getenv("DB_USER_DEV"),
		os.Getenv("DB_PASS_DEV"),
		os.Getenv("DB_SCHEMA_DEV"),
		os.Getenv("DB_PORT_DEV"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	dbConn.AutoMigrate(&entity.Content{})
	dbConn.AutoMigrate(&entity.Content_Type{})
	
	if err != nil {
		log.Println("Error connect to PostgreSQL: ", err.Error())
		return nil, nil, err
	}

	sqlDb, errDb := dbConn.DB()
	if errDb != nil {
		log.Println(errDb)
	} else {
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(1000)
	}

	log.Println("Postgres connection success")
	return dbConn, sqlDb, nil
}

func ConnectPostgresTest() (*gorm.DB, *sql.DB, error) {
	// Environment Variables
	err := godotenv.Load(`C:\Users\DRAGON\Documents\Golang\go\src\content-service-v2-2\app\environments\app.env`)
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`,
		os.Getenv("DB_HOST_DEV_TEST"),
		os.Getenv("DB_USER_DEV_TEST"),
		os.Getenv("DB_PASS_DEV_TEST"),
		os.Getenv("DB_SCHEMA_DEV_TEST"),
		os.Getenv("DB_PORT_DEV_TEST"),
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbConn.AutoMigrate(&entity.Content{})
	dbConn.AutoMigrate(&entity.Content_Type{})
	
	if err != nil {
		log.Println("Error connect to PostgreSQL Test: ", err.Error())
		return nil, nil, err
	}

	sqlDb, errDb := dbConn.DB()
	if errDb != nil {
		log.Println(errDb)
	} else {
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(1000)
	}

	log.Println("Postgres Test connection success")
	return dbConn, sqlDb, nil
}

