package content_usecase

import (
	"content-service-v2/app/domain/repository"
)

type serviceContent struct {
	repository repository.RepositoryContent
}

func NewServiceContent(repository repository.RepositoryContent) *serviceContent {
	return &serviceContent{repository}
}
