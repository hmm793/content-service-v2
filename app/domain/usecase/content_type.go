package usecase

import (
	"content-service-v2/app/domain/entity"
)

type ServiceContentType interface {
	GetAllContentType() ([]entity.Content_Type, error)
}
