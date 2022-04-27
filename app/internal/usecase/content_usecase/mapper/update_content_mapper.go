package mapper

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/domain/entity"
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func UpdateContentMapper(input dto.UpdateContentInput, contentData entity.Content) entity.Content{
	inputData, _ := json.Marshal(input.Data)

	contentData.Data = postgres.Jsonb{
		RawMessage: json.RawMessage(inputData),
	}
	contentData.IsPublished = input.IsPublished
	contentData.Name = input.Name
	contentData.ContentTypeId = input.ContentTypeId
	contentData.UpdatedByName = input.UpdatedByName
	contentData.UpdatedById = input.UpdatedById
	contentData.UpdatedAt = time.Now()
	return contentData
}