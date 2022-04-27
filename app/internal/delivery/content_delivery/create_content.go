package content_delivery

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/internal/delivery/content_delivery/formatter"
	"content-service-v2/app/internal/delivery/helper"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *contentHandler) CreateContent(c *gin.Context) {
	input := dto.CreateContentInput{}
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	for _, data := range input.Data.([]interface{}) {
		url := data.(map[string]interface{})["image"]
		path := strings.Split(url.(string), "/")
		path[3] = "asset"
		fixedPath := strings.Join(path, "/")
		
		data.(map[string]interface{})["image"] = fixedPath
	}


	content_data, err := h.contentService.CreateContent(input)
	
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	for _, data := range input.Data.([]interface{}) {
		url := data.(map[string]interface{})["image"]
		path := strings.Split(url.(string), "/")
		
		// Mindahkan file ke file existing
		// Lokasi lama
		oldLocation := os.Getenv("tempPath") + path[4] + "/" + path[5]
		
		// Cek apakah folder dengan nama user id sudah ada klo belum buat baru
		newPath := fmt.Sprintf(os.Getenv("assetPath") + "%s",path[4])
		_, err = os.Stat(newPath)

		if os.IsNotExist(err) {
        	if err := os.MkdirAll(newPath, os.ModePerm); err != nil {
				errorMessage := gin.H{
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, errorMessage)
				return
			}
    	}
		
		// Lokasi Baru
		newLocation := os.Getenv("assetPath") + path[4] + "/" + path[5]
		err = os.Rename(oldLocation, newLocation)
		if err != nil {
			errorMessage := gin.H{
				"message": err.Error(),
			}
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	formatCreateContentResponse := formatter.FormatCreateContentResponse(content_data)
	c.JSON(http.StatusOK, formatCreateContentResponse)
}
