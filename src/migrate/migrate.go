package migrate

import "gorm.io/gorm"

// Migrate database
func Migrate(db *gorm.DB) {
	db.AutoMigrate()
}
