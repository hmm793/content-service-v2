package content_usecase

import (
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/repository/psql"
	"content-service-v2/app/internal/usecase/content_usecase/mapper"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteContentFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := psql.NewRepositoryContent(db)
	contentService := NewServiceContent(contentRepository)

	id := 1
	content, err := contentService.repository.FindContentById(id)

	assert.Equal(t, nil, err)
	fixedContent := mapper.ScannedContentToContent(content)

	deleteByName := "andi"
	deleteById := 20283782743839
	mapperData := mapper.DeleteContentMapper(deleteByName, deleteById, fixedContent)

	rowAffected, err := contentService.repository.DeleteContentById(mapperData)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), rowAffected)
}