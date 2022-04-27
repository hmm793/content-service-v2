package mapper

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/domain/entity"
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

//
func CreateContentMapper(input dto.CreateContentInput) entity.Content{
	inputData, _ := json.Marshal(input.Data)

	var t time.Time
	contentData := entity.Content{}
	
	contentData.Data = postgres.Jsonb{
		RawMessage: json.RawMessage(inputData),
	}
	contentData.Name = input.Name
	contentData.IsPublished = input.IsPublished
	contentData.CreatedByName = input.CreatedByName
	contentData.ContentTypeId = input.ContentTypeId
	contentData.CreatedById = input.CreatedById
	contentData.CreatedAt = time.Now()
	contentData.UpdatedAt = t
	contentData.DeletedAt = t
	return contentData
}
//