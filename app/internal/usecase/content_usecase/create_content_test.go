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

func TestCreateContentFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := psql.NewRepositoryContent(db)
	contentService := NewServiceContent(contentRepository)

	input := dto.CreateContentInput{
		Name: "Banner Depan Updated",
		IsPublished: 1,
		ContentTypeId: 1,
		Data: `[{"image":"/api/v1/asset/123123131/dd1b31f07b9e4244a7fbede515448c9f-img_avatar-2022-04-21.png", "description":"this is a description 5"},
        {"image":"/api/v1/images/123123131/05e9f8b886824801a15baa2066026ec8-img_avatar-2022-04-21.png","subscribtion":"content 1", "description":"this is a description"}]`,
		CreatedByName: "tom",
		CreatedById: 93908901890,
	}
	
	contentData := mapper.CreateContentMapper(input)
	newContentData, err := contentService.repository.SaveContent(contentData)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, newContentData.ID)
	assert.Equal(t, "Banner Depan Updated", newContentData.Name)
	assert.NotEqual(t, nil, newContentData.Data)
	assert.Equal(t, 1, newContentData.IsPublished)
	assert.Equal(t, 1, newContentData.ContentTypeId)
	assert.Equal(t, "tom", newContentData.CreatedByName);
	assert.Equal(t, 93908901890, newContentData.CreatedById)
}