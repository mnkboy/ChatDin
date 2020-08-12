package migratefunctions

import (
	"github.com/jinzhu/gorm"
)

//CreateTables : creamos tablas si existen
func CreateTables(db *gorm.DB) {
	db.CreateTable(user)
}
