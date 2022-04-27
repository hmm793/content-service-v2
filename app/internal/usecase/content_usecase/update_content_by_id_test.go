package content_usecase

import (
	"content-service-v2/app/domain/dto"
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/repository/psql"
	"content-service-v2/app/internal/usecase/content_usecase/mapper"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateContentByIdFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()
	
	id := 3
	input := dto.UpdateContentInput{
		Name: "Banner Depan Updated",
		UpdatedByName: "john",
		IsPublished: 1,
		ContentTypeId: 1,
		UpdatedById : 283478293,
		Data: `[{"image":"/api/v1/asset/123123131/dd1b31f07b9e4244a7fbede515448c9f-img_avatar-2022-04-21.png", "description":"this is a description 5"},
        {"image":"/api/v1/images/123123131/05e9f8b886824801a15baa2066026ec8-img_avatar-2022-04-21.png","subscribtion":"content 1", "description":"this is a description"}]`,
	}
	
	contentRepository := psql.NewRepositoryContent(db)
	contentService := NewServiceContent(contentRepository)

	tempContentData, err := contentService.repository.FindContentById(id)
	assert.Equal(t, nil, err)

	contentData := mapper.ScannedContentToContent(tempContentData)
	mapperData := mapper.UpdateContentMapper(input, contentData)
	contentDataUpdated, err := contentService.repository.UpdateContent(mapperData)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, contentDataUpdated.ID)
	assert.NotEqual(t, nil, contentDataUpdated.Data)
	assert.Equal(t, "Banner Depan Updated", contentDataUpdated.Name)
	assert.Equal(t, 1, contentDataUpdated.IsPublished)
	assert.Equal(t, 1, contentDataUpdated.ContentTypeId)
	assert.Equal(t, "john", contentDataUpdated.UpdatedByName);
	assert.Equal(t, 283478293, contentDataUpdated.UpdatedById)
}