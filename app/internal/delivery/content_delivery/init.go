package content_delivery

import (
	"content-service-v2/app/domain/usecase"
)

type contentHandler struct {
	contentService usecase.ServiceContent
}

func NewContentHandler(contentService usecase.ServiceContent) *contentHandler {
	return &contentHandler{contentService}
}
