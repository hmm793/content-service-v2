package formatter

import (
	"encoding/json"
	"time"
)

//
type GetContentByIdResponseFormatter struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	CreatedByName string    `json:"createdByName"`
	CreatedById   int       `json:"createdById"`
	IsPublished   int       `json:"isPublished"`
	Data          json.RawMessage    `json:"data"`
	ContentTypeId int       `json:"contentTypeId"`
	UpdatedByName string    `json:"updatedByName"`
	UpdatedById   int       `json:"updatedById"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
//