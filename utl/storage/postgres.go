package storage

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDB(url string) *gorm.DB {
	sqlDB, err := sql.Open("pgx", url) // protocol://user:password@host:port/database
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	return gormDB
}
