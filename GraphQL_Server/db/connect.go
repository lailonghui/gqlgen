package db

import (
	"github.com/jinzhu/gorm"
)

type Connect struct {
}

func (con *Connect) GetConnection() *gorm.DB {
	return Pop("postgres")
}
