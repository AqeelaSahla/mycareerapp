package mysql

import (
	"mycareerapp/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	if err := db.AutoMigrate(
		entity.Article{},
		entity.User{},
		entity.Jobpulse{},
	); err != nil {
		return err
	}

	return nil
}
