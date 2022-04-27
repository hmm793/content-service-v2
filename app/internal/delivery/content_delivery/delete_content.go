package content_delivery

import (
	"content-service-v2/app/domain/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *contentHandler) DeleteContent(c *gin.Context) {
	input := dto.GetContentIdInput{}
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}


	currContent, err := h.contentService.FindContentById(input.ID)
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	hasil , _ := currContent.Data.MarshalJSON()
	var result []interface{}
    json.Unmarshal(hasil, &result)

	// Hapus File 
	for _, data := range result {
		url := data.(map[string]interface{})["image"]
		
		path := strings.Split(url.(string), "/")

		directoryPath := fmt.Sprintf(os.Getenv("assetPath")+"%s/%s",path[4], path[5])
		
		err := os.Remove(directoryPath)

		if err != nil {
			errorMessage := gin.H{
				"message": err.Error(),
			}
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	rowAffected,err := h.contentService.DeleteContent(input.ID)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	response := gin.H{
		"rowAffected" :rowAffected,
	}
	c.JSON(http.StatusOK, response)
}
