package seeders

import (
	"content-service-v2/app/domain/entity"
	"content-service-v2/app/internal/config/fakers"

	"gorm.io/gorm"
)

type ContentTypeSeeder struct {
	data []entity.Content_Type
}

func CreateContentTypeSeeder(db *gorm.DB) ContentTypeSeeder {
	return ContentTypeSeeder{
		data : fakers.ContentTypeFaker(db),
	}
}

func ContentTypeDBSeed(db *gorm.DB) error {
	for _, datum := range CreateContentTypeSeeder(db).data {
		err := db.Debug().Create(&datum).Error
		if err != nil {
			return err
		}
	}
	return nil
}