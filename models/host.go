package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	URL string
}

var Hosts []Host
