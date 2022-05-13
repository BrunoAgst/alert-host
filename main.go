package main

import (
	"host-monitoring/db"
	"host-monitoring/services"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
		os.Exit(-1)
	}

	db.ConnectDatase()
	for {
		hosts := services.SearchAllHosts()
		services.MonitorHosts(hosts)
		time.Sleep(1 * time.Hour)
	}
}
