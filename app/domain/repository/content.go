package repository

import (
	"content-service-v2/app/domain/entity"
)

type RepositoryContent interface {
	SaveContent(content entity.Content) (entity.Content, error)//
	FindContentById(id int) (entity.ScanedContent, error)//
	FindAll() ([]entity.ScanedContent, error)//
	FindAllWithLimit(limit int) ([]entity.ScanedContent, error)//
	
	UpdateContent(content entity.Content) (entity.Content, error)//

	DeleteContentById(content entity.Content) (int64,error)//
}
