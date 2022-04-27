package formatter

import "content-service-v2/app/domain/entity"

func FormatContentTypeResponse(contentType entity.Content_Type) ContentTypeFormatter {
	contentTypeFormatter := ContentTypeFormatter{}
	contentTypeFormatter.ID = contentType.ID
	contentTypeFormatter.Name = contentType.Name

	return contentTypeFormatter
}
