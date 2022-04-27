package formatter

import "content-service-v2/app/domain/entity"

//
func FormatGetContentResponse(content entity.ScanedContent) GetContentByIdResponseFormatter {
	getContentFormatter := GetContentByIdResponseFormatter{}
	getContentFormatter.ID = content.ID//
	getContentFormatter.Name = content.Name//
	getContentFormatter.Data = content.Data//
	getContentFormatter.IsPublished = content.IsPublished//
	getContentFormatter.ContentTypeId = content.ContentTypeId//
	getContentFormatter.CreatedByName = content.CreatedByName //
	getContentFormatter.UpdatedByName = content.UpdatedByName
	getContentFormatter.CreatedById = content.CreatedById
	getContentFormatter.UpdatedById = content.UpdatedById
	getContentFormatter.CreatedAt = content.CreatedAt
	getContentFormatter.UpdatedAt = content.UpdatedAt
	return getContentFormatter
}
//