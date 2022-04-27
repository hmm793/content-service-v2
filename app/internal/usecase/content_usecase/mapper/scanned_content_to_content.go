package mapper

import (
	"content-service-v2/app/domain/entity"
	"encoding/json"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func ScannedContentToContent(input entity.ScanedContent) entity.Content {
	mappedContent := entity.Content{}
	mappedContent.ID = input.ID
	mappedContent.Name = input.Name

	inputData, _ := json.Marshal(input.Data)
	
	mappedContent.Data = postgres.Jsonb{
		RawMessage: json.RawMessage(inputData),
	}
	mappedContent.IsPublished = input.IsPublished
	mappedContent.CreatedByName = input.CreatedByName
	mappedContent.DeletedByName = input.DeletedByName
	mappedContent.UpdatedByName = input.UpdatedByName
	mappedContent.CreatedById = input.CreatedById
	mappedContent.UpdatedById = input.UpdatedById
	mappedContent.DeletedById = input.DeletedById
	mappedContent.ContentTypeId = input.ContentTypeId
	mappedContent.CreatedAt = input.CreatedAt
	mappedContent.UpdatedAt = input.UpdatedAt
	mappedContent.DeletedAt = input.DeletedAt
	mappedContent.ContentType = input.ContentType
	return mappedContent 
}