package formatter

import "content-service-v2/app/domain/entity"

func FormatContentTypesResponse(contentTypes []entity.Content_Type) []ContentTypeFormatter {
	var contentTypesFormatter []ContentTypeFormatter
	for _, contentType := range contentTypes {
		contentTypesFormatter = append(contentTypesFormatter, FormatContentTypeResponse(contentType))
	}
	return contentTypesFormatter
}
