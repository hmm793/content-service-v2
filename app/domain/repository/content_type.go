package repository

import (
	"content-service-v2/app/domain/entity"
)

type RepositoryContentType interface {
	FindAllContentType() ([]entity.Content_Type, error)
}
