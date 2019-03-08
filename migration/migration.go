package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/keiya01/myblog/fields"
)

func Set(db *gorm.DB) {
	db.AutoMigrate(&fields.Blog{})
}
