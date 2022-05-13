package db

import (
	"host-monitoring/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatase() {
	connectionString := "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_DBNAME") + " port=" + os.Getenv("DATABASE_PORT") + " sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error dont connect to database: ", err)
	}

	DB.AutoMigrate(&models.Host{})
}
