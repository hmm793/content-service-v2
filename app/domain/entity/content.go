package entity

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Content struct {
	ID            int
	Name          string
	Data          postgres.Jsonb `gorm:"type:jsonb;column:data"`
	IsPublished   int
	CreatedByName string
	DeletedByName string
	UpdatedByName string
	CreatedById   int
	UpdatedById   int
	DeletedById   int
	ContentTypeId int
	CreatedAt 	  time.Time
	UpdatedAt 	  time.Time
	DeletedAt 	  time.Time
	ContentType   Content_Type
}

type ScanedContent struct {
	ID            int
	Name          string
	Data          json.RawMessage
	IsPublished   int
	CreatedByName string
	DeletedByName string
	UpdatedByName string
	CreatedById   int
	UpdatedById   int
	DeletedById   int
	ContentTypeId int
	CreatedAt 	  time.Time
	UpdatedAt 	  time.Time
	DeletedAt 	  time.Time
	ContentType   Content_Type
}