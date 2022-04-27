package psql

import (
	"content-service-v2/app/domain/entity"
	"content-service-v2/app/internal/config"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllContentTypeFromRepository(t *testing.T){
	db, Sql, err := config.ConnectPostgresTest()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentTypeRepository := NewRepositoryContentType(db)

	contentTypeDatas := []entity.Content_Type{}
	err = contentTypeRepository.db.Find(&contentTypeDatas).Error
	
	assert.Equal(t, nil, err)
}

