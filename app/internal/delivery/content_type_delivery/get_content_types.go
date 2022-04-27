package content_type_delivery

import (
	"content-service-v2/app/internal/delivery/content_type_delivery/formatter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *contentTypeHandler) GetContentTypes(c *gin.Context) {
	content_type_datas, err := h.service.GetAllContentType()
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	response := formatter.FormatContentTypesResponse(content_type_datas)
	c.JSON(http.StatusOK, response)
}