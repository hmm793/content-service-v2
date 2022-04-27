package content_usecase

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/domain/entity"
	"content-service-v2/app/internal/usecase/content_usecase/mapper"
	"errors"
)

func (s *serviceContent) UpdateContentById(id int, input dto.UpdateContentInput) (entity.Content, error) {
	tempContentData, err := s.repository.FindContentById(id)

	contentData := mapper.ScannedContentToContent(tempContentData)

	if err != nil {
		return contentData, err
	}

	if contentData.ID == 0 {
		return contentData, errors.New("Content Not Found")
	}

	mapperData := mapper.UpdateContentMapper(input, contentData)
	contentDataUpdated, err := s.repository.UpdateContent(mapperData)

	if err != nil {
		return contentData, err
	}

	return contentDataUpdated, nil
}
