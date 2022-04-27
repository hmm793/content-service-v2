package content_type_usecase

import (
	"content-service-v2/app/internal/config"
	"content-service-v2/app/internal/repository/psql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllContentTypeFromUsecase(t *testing.T) {
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentTypeRepository := psql.NewRepositoryContentType(db)
	contentTypeService := NewServiceContentType(contentTypeRepository)

	_, err = contentTypeService.repository.FindAllContentType()
	assert.Equal(t, nil, err)
}
