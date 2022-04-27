package content_usecase

import (
	"content-service-v2/app/domain/entity"
	"errors"
)

//
func (s *serviceContent) FindContentById(id int) (entity.ScanedContent, error) {
	content, err := s.repository.FindContentById(id)
	if err != nil {
		return content, err
	}
	if content.ID == 0 {
		return content, errors.New("Content Not Found")
	}
	return content, nil
}
//