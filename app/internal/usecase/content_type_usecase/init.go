package content_type_usecase

import (
	"content-service-v2/app/domain/repository"
)

type serviceContentType struct {
	repository repository.RepositoryContentType
}

func NewServiceContentType(repository repository.RepositoryContentType) *serviceContentType {
	return &serviceContentType{repository}
}
