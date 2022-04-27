package content_usecase

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/domain/entity"
	"content-service-v2/app/internal/usecase/content_usecase/mapper"
)

//
func (s *serviceContent) CreateContent(input dto.CreateContentInput) (entity.Content, error) {
	contentData := mapper.CreateContentMapper(input)
	newContentData, err := s.repository.SaveContent(contentData)
	if err != nil {
		return newContentData, err
	}
	return newContentData, nil
}
//