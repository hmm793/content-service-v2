package formatter

import (
	"content-service-v2/app/domain/entity"
)

func FormatCreateContentResponse(content entity.Content) CreateContentResponseFormatter {
	createContentFormatter := CreateContentResponseFormatter{}
	createContentFormatter.ID = content.ID
	createContentFormatter.Name = content.Name
	createContentFormatter.CreatedByName = content.CreatedByName
	createContentFormatter.CreatedById = content.CreatedById
	createContentFormatter.IsPublished = content.IsPublished
	createContentFormatter.ContentTypeId = content.ContentTypeId
	createContentFormatter.Data = content.Data
	return createContentFormatter
}