package formatter

import "content-service-v2/app/domain/entity"

func FormatUpdateContentResponse(content entity.Content) UpdateContentByIdResponseFormatter {
	updateContentFormatter := UpdateContentByIdResponseFormatter{}
	updateContentFormatter.ID = content.ID
	updateContentFormatter.Name = content.Name
	updateContentFormatter.Data = content.Data
	updateContentFormatter.IsPublished = content.IsPublished
	updateContentFormatter.ContentTypeId = content.ContentTypeId
	updateContentFormatter.UpdatedByName = content.UpdatedByName
	updateContentFormatter.UpdatedById = content.UpdatedById
	return updateContentFormatter
}
