package usecase

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/domain/entity"
)

type ServiceContent interface {
	CreateContent(input dto.CreateContentInput) (entity.Content, error)//
	FindContentById(id int) (entity.ScanedContent, error)//
	FindAllContent(limit int, offset int, filterArr interface{}) ([]entity.ScanedContent, error)//
	UpdateContentById(id int, input dto.UpdateContentInput) (entity.Content, error)//

	DeleteContent(id int) (int64,error)
}
