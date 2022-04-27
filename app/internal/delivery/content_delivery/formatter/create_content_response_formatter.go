package formatter

type CreateContentResponseFormatter struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	CreatedByName string      `json:"createdByName"`
	CreatedById   int         `json:"createdById"`
	IsPublished   int         `json:"isPublished"`
	Data          interface{} `json:"data"`
	ContentTypeId int         `json:"contentTypeId"`
}
