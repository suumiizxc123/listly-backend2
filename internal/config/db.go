package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "postgresql://citizix_user:S3cret@103.48.116.100:5432/citizix_db"
	// dsn := "postgresql://postgres:mik%23123@192.168.0.103:5432/postgres"
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
