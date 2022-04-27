package content_delivery

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/internal/delivery/content_delivery/formatter"
	"content-service-v2/app/internal/delivery/helper"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *contentHandler) UpdateContent(c *gin.Context) {
	var inputId dto.GetContentIdInput
	err := c.ShouldBindUri(&inputId)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	
	inputData := dto.UpdateContentInput{}
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// Buat Validasi Image
	currContent, err := h.contentService.FindContentById(inputId.ID)
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

	var arrValidation []interface{}
	for _, data := range result {
		url := data.(map[string]interface{})["image"]
		arrValidation = append(arrValidation, url)
	}
	
	// Dari Input
	var arrInputData []interface{}
	for _, data := range inputData.Data.([]interface{}) {
		url := data.(map[string]interface{})["image"]
		arrInputData = append(arrInputData, url)
	}

	for _, newImage := range arrInputData {
		if contains(arrValidation, newImage) {
			arrValidation = remove(arrValidation,newImage)
		}
	}

	for _, data := range arrValidation {
		path := strings.Split(data.(string), "/")

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


	for _, data := range inputData.Data.([]interface{}) {
		url := data.(map[string]interface{})["image"]
		path := strings.Split(url.(string), "/")
		if path[3] != "asset" {
			path[3] = "asset"
			fixedPath := strings.Join(path, "/")
			
			data.(map[string]interface{})["image"] = fixedPath

			// Mindahkan file ke file existing
			// Lokasi lama
			oldLocation := "app/temp/" + path[4] + "/" + path[5]
			
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
	}

	contentUpdated, err := h.contentService.UpdateContentById(inputId.ID, inputData)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatContentResponse := formatter.FormatUpdateContentResponse(contentUpdated)

	c.JSON(http.StatusOK, formatContentResponse)
}


func remove(items []interface{}, item interface{}) []interface{} {
    var newitems []interface{}

    for _, i := range items {
        if i != item {
            newitems = append(newitems, i)
        }
    }

    return newitems
}


func contains(s []interface{}, e interface{}) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}