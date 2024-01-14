package utils

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Establish a database connection
func GetPostgresConnection(dbName string) *gorm.DB {
	connectionStr := fmt.Sprintf("user=postgres dbname=%s sslmode=disable", dbName)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := gormDb.DB()
	if err != nil {
		panic(err.Error())
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return gormDb
}

// Close a database connection
func ClosePostgresConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to exit database connection")
	}
	dbSQL.Close()
}
