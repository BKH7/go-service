package migrate

import (
	"gitlab.com/bunlert274/go-service/src/model"
	"gorm.io/gorm"
)

// Migrate database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
