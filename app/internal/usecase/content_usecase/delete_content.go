package content_usecase

import (
	"content-service-v2/app/internal/usecase/content_usecase/mapper"
	"errors"
)

func (s *serviceContent) DeleteContent(id int) (int64,error) {
	content, err := s.repository.FindContentById(id)

	if err != nil {
		return 0, err
	}
	
	if content.ID == 0 {
		return 0, errors.New("Content Not Found")
	}

	fixedContent := mapper.ScannedContentToContent(content)

	deleteByName := "andi"
	deleteById := 20283782743839
	mapperData := mapper.DeleteContentMapper(deleteByName, deleteById, fixedContent)
	
	rowAffected, err := s.repository.DeleteContentById(mapperData)
	if err != nil {
		return rowAffected,err
	}
	
	return rowAffected, nil
}
