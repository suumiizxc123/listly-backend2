package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "postgresql://doadmin:AVNS_OPeOEVEkaDXd_Me7kYf@db-cargo-do-user-16113953-0.c.db.ondigitalocean.com:25060/defaultdb"
	// dsn := "postgresql://postgres:mik%23123@192.168.0.103:5432/postgres"
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
