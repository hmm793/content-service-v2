package formatter

import "time"

type GetAllContentResponseFormatter struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	CreatedByName   string      `json:"createdByName"`
	CreatedById     int         `json:"createdById"`
	IsPublished     int         `json:"isPublished"`
	Data            interface{} `json:"data"`
	ContentTypeId   int         `json:"contentTypeId"`
	ContentTypeName string      `json:"contentTypeName"`
	UpdatedByName   string      `json:"updatedByName"`
	UpdatedById     int         `json:"updatedById"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
}
