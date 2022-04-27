package fakers

import (
	"content-service-v2/app/domain/entity"

	"gorm.io/gorm"
)

func ContentTypeFaker(db *gorm.DB) []entity.Content_Type {
	data := []entity.Content_Type{
		{Name: "banner"},
		{Name: "section A"},
		{Name: "section B"},
		{Name: "section C"},
	}
	return data
}