package dto

type CreateContentInput struct {
	Name          string      `json:"name" binding:"required"`
	CreatedByName string      `json:"createdByName" binding:"required"`
	CreatedById   int         `json:"createdById" binding:"required"`
	IsPublished   int         `json:"isPublished" binding:"required,number"`
	Data          interface{} `json:"data" binding:"required"`
	ContentTypeId int         `json:"contentTypeId" binding:"required"`
}

type GetContentIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateContentInput struct {
	Name          string      `json:"name" binding:"required"`
	UpdatedByName string      `json:"updatedByName" binding:"required"`
	UpdatedById   int         `json:"updatedById" binding:"required"`
	IsPublished   int         `json:"isPublished" binding:"required,number"`
	Data          interface{} `json:"data" binding:"required"`
	ContentTypeId int         `json:"contentTypeId" binding:"required"`
}
