package content_usecase

import (
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/repository/psql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentByIdFromUsecase(t *testing.T) {
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
	assert.Equal(t, id, content.ID)
}