package formatter

type UpdateContentByIdResponseFormatter struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	UpdatedByName string      `json:"updatedByName"`
	UpdatedById   int         `json:"updatedById"`
	IsPublished   int         `json:"isPublished"`
	Data          interface{} `json:"data"`
	ContentTypeId int         `json:"contentTypeId"`
}
