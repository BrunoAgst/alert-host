package main

import (
	"host-monitoring/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}
	hosts := []services.Host{
		{URL: "https://www.google.com.br"},
		{URL: "https://www.mercadolivre.com.br"},
		{URL: "https://www.globo.com"},
	}

	services.MonitorHosts(hosts)

}
