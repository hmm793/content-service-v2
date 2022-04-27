package psql

import (
	"content-service-v2/app/domain/entity"
	"content-service-v2/app/internal/config"
	"encoding/json"
	"log"
	"testing"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func TestSaveContentFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := NewRepositoryContent(db)

	content := entity.Content{}
	content.Data = postgres.Jsonb{
		RawMessage: json.RawMessage(`[{"image":"/api/v1/images/123123131/bab83187fa23441b9e401dc74f235c13-img_avatar-2022-04-21.png", "description":"this is a description 5"}, {"image":"/api/v1/images/123123131/bab83187fa23441b9e401dc74f235c13-img_avatar-2022-04-21.png", "description":"this is a description 5"}]`),
	}
	content.IsPublished = 1
	content.ContentTypeId = 1
	content.CreatedByName = "Donald"
	content.CreatedById = 123331

	err = contentRepository.db.Create(&content).Error
	
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil , content.Data)
	assert.Equal(t, 1 , content.IsPublished)
	assert.Equal(t, 1 , content.ContentTypeId)
	assert.Equal(t, "Donald" , content.CreatedByName)
	assert.Equal(t, 123331 , content.CreatedById)
	assert.NotEqual(t, 0 , content.ID)
}

func TestFindContentByIdFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := NewRepositoryContent(db)

	id := 1

	var content entity.Content
	err = contentRepository.db.Where("id = ?", id).Find(&content).Error
	
	assert.Equal(t, nil, err)
	assert.Equal(t, id, content.ID)
}	

func TestUpdateContentFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := NewRepositoryContent(db)
	content := entity.Content{}
	content.ID = 1
	content.Data = postgres.Jsonb{
		RawMessage: json.RawMessage(`[{"image":"/api/v1/images/123123131/bab83187fa23441b9e401dc74f235c13-img_avatar-2022-04-21.png", "description":"this is a description 5 updated"}, {"image":"/api/v1/images/123123131/bab83187fa23441b9e401dc74f235c13-img_avatar-2022-04-21.png", "description":"this is a description 5 updated"}]`),
	}
	content.IsPublished = 0
	content.ContentTypeId = 1
	content.UpdatedByName = "Thomas"
	content.UpdatedById = 123332
	err = contentRepository.db.Save(&content).Error
	assert.Equal(t, nil, err)
}

func TestFindAllContentFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()	
	if err != nil {
		log.Println("error postgresql connection: ", err)	
	}
	defer Sql.Close()

	contentRepository := NewRepositoryContentType(db)

	var contents []entity.Content
	err = contentRepository.db.Find(&contents).Error
	
	assert.Equal(t, nil, err)
}
func TestFindAllContentWithLimitFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()	
	if err != nil {
		log.Println("error postgresql connection: ", err)	
	}
	defer Sql.Close()

	limit := 10
	contentRepository := NewRepositoryContentType(db)

	var contents []entity.Content
	err = contentRepository.db.Limit(limit).Find(&contents).Error
	
	assert.Equal(t, nil, err)
}
func TestFindAllContentWithLimitAndOffsetFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()	
	if err != nil {
		log.Println("error postgresql connection: ", err)	
	}
	defer Sql.Close()

	limit := 5
	contentRepository := NewRepositoryContentType(db)

	var contents []entity.Content
	err = contentRepository.db.Limit(limit).Find(&contents).Error
	
	assert.Equal(t, nil, err)
}
