package content_usecase

import (
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/repository/psql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllContentFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := psql.NewRepositoryContent(db)
	contentService := NewServiceContent(contentRepository)

	_, err = contentService.repository.FindAll()
	assert.Equal(t, nil, err)
}

func TestFindAllContentWithLimitFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentRepository := psql.NewRepositoryContent(db)
	contentService := NewServiceContent(contentRepository)

	limit := 10
	_, err = contentService.repository.FindAllWithLimit(limit)
	assert.Equal(t, nil, err)
}

