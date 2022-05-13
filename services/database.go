package services

import (
	"host-monitoring/db"
	"host-monitoring/models"
)

func SearchAllHosts() []models.Host {
	var hosts []models.Host
	db.DB.Find(&hosts)
	return hosts
}
