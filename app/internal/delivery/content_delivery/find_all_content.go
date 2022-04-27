package content_delivery

import (
	"content-service-v2/app/internal/delivery/content_delivery/formatter"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
//
func (h *contentHandler) FindAllContent(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	filter:= c.Query("filter")

	filterArr := filter
	contents_data, err := h.contentService.FindAllContent(limit, offset, filterArr)
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	};
	response := formatter.FormatGetAllContentResponse(contents_data)
	c.JSON(http.StatusOK, response)
}
//