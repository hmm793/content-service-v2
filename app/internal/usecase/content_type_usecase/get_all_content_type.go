package content_type_usecase

import (
	"content-service-v2/app/domain/entity"
	"errors"
)

func (s *serviceContentType) GetAllContentType() ([]entity.Content_Type, error) {
	contentTypeDatas, err := s.repository.FindAllContentType()

	if err != nil {
		return contentTypeDatas, err
	}

	if len(contentTypeDatas) == 0 {
		return contentTypeDatas, errors.New("There is no data")
	}

	return contentTypeDatas, nil
}
