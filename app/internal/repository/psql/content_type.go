package psql

import (
	"content-service-v2/app/domain/entity"

	"gorm.io/gorm"
)
type repositoryContentType struct {
	db *gorm.DB
}

func NewRepositoryContentType(db *gorm.DB) *repositoryContentType {
	return &repositoryContentType{db}
}

func (r *repositoryContentType) FindAllContentType() ([]entity.Content_Type, error) {
	contentTypeDatas := []entity.Content_Type{}
	err := r.db.Find(&contentTypeDatas).Error
	if err != nil {
		return contentTypeDatas, err
	}
	return contentTypeDatas, nil
}

