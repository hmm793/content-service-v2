package mapper

import (
	"content-service-v2/app/domain/entity"
	"time"
)

func DeleteContentMapper(deleteByName string, deleteById int,  contentData entity.Content) entity.Content{
	contentData.DeletedAt = time.Now()
	contentData.DeletedByName = deleteByName
	contentData.DeletedById = deleteById
	return contentData
}