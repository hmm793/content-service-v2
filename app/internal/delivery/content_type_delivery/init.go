package content_type_delivery

import (
	"content-service-v2/app/domain/usecase"
)

type contentTypeHandler struct {
	service usecase.ServiceContentType
}

func NewContentTypeHandler(service usecase.ServiceContentType) *contentTypeHandler {
	return &contentTypeHandler{service}
}
