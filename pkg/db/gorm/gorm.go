package gorm

import (
	"gorm.io/gorm"
)

func (d DbConfig) GetDbConnection() *gorm.DB {
	return d.db
}
