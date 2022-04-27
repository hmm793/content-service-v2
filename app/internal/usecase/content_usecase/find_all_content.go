package content_usecase

import (
	"content-service-v2/app/domain/entity"
	"errors"
)

func (s *serviceContent) FindAllContent(limit int, offset int, filterArr interface{}) ([]entity.ScanedContent, error) {
	// var newFilterArr formatter.Filter
	// var filterArr1 = `{"contentTypeId":1}`
	// err := json.Unmarshal([]byte(filterArr1.(string)), &newFilterArr)
    // if err != nil {
    //     fmt.Println("error:", err)
    // }

	
	if limit != 0 {
		contents, err := s.repository.FindAllWithLimit(limit)
		if err != nil {
			return contents, err
		}
		if len(contents) == 0 {
			return contents, errors.New("Content Not Found")
		}
		if offset != 0 {
			if offset > len(contents) {
				return contents, errors.New("Offset is out of range")
			}
			return contents[offset:], nil
		}
		return contents, nil
	} else {
		contents, err := s.repository.FindAll()
		if err != nil {
			return contents, err
		}
		if len(contents) == 0 {
			return contents, errors.New("Content Not Found")
		}
		return contents, nil
	}
}