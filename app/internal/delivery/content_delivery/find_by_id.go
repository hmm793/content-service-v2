package content_delivery

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/internal/delivery/content_delivery/formatter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *contentHandler) FindById(c *gin.Context) {
	input := dto.GetContentIdInput{}
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	content_data, err := h.contentService.FindContentById(input.ID)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatGetContentByIdResponse := formatter.FormatGetContentResponse(content_data)

	c.JSON(http.StatusOK, formatGetContentByIdResponse)
}
