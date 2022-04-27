package formatter

import (
	"content-service-v2/app/domain/entity"
)

//
func FormatGetAllContentResponseForEachContent(content entity.ScanedContent) GetAllContentResponseFormatter {
	getAllContentFormatter := GetAllContentResponseFormatter{}
	getAllContentFormatter.ID = content.ID
	getAllContentFormatter.Name = content.Name
	getAllContentFormatter.CreatedByName = content.CreatedByName
	getAllContentFormatter.CreatedById = content.CreatedById
	getAllContentFormatter.IsPublished = content.IsPublished
	getAllContentFormatter.Data = content.Data
	getAllContentFormatter.ContentTypeId = content.ContentTypeId
	getAllContentFormatter.UpdatedByName = content.UpdatedByName
	getAllContentFormatter.UpdatedById = content.UpdatedById
	getAllContentFormatter.CreatedAt = content.CreatedAt
	getAllContentFormatter.UpdatedAt = content.UpdatedAt
	getAllContentFormatter.ContentTypeName = content.ContentType.Name
	return getAllContentFormatter
}
//

//
func FormatGetAllContentResponse(contentData []entity.ScanedContent) []GetAllContentResponseFormatter {
	var contentsFormatter []GetAllContentResponseFormatter
	for _, content := range contentData {
		contentsFormatter = append(contentsFormatter, FormatGetAllContentResponseForEachContent(content))
	}
	return contentsFormatter
}
//