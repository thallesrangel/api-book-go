package migrations

import (
	"github/thallesrangel/api-book-go/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
