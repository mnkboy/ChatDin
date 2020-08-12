package migratefunctions

import (
	"github.com/jinzhu/gorm"
)

//DropTables : borramos tablas si existen
func DropTables(db *gorm.DB) {
	db.DropTableIfExists(user)
}
