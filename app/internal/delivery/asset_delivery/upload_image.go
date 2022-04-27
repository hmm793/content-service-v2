package asset_delivery

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *assetHandler) UploadAsset(c *gin.Context) {
	file, err := c.FormFile("fileName")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return
	}

	idUser := c.PostForm("userId")
	if idUser == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message" :"Id User is empty"})
		return
	}
	
	path := fmt.Sprintf("app/temp/%s",idUser)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
        if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
    }
	
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	pathForSave := fmt.Sprintf("app/temp/%s/%s-%s-%s.%s",idUser, uuid, strings.Split(file.Filename, ".")[0], time.Now().Format("2006-01-02"),  strings.Split(file.Filename, ".")[1])

	fmt.Println("pathForSave",pathForSave)
	err = c.SaveUploadedFile(file, pathForSave)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image" : fmt.Sprintf("%s/images/%s/%s-%s-%s.%s","/api/v1",idUser,uuid, strings.Split(file.Filename, ".")[0], time.Now().Format("2006-01-02"),  strings.Split(file.Filename, ".")[1] ),
	})
}
