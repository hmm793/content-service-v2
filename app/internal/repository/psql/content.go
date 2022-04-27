package psql

import (
	"content-service-v2/app/domain/entity"
	"time"

	"gorm.io/gorm"
)

type repositoryContent struct {
	db *gorm.DB
}

func NewRepositoryContent(db *gorm.DB) *repositoryContent {
	return &repositoryContent{db}
}
//
func (r *repositoryContent) SaveContent(content entity.Content) (entity.Content, error) {
	err := r.db.Create(&content).Error
	if err != nil {
		return content, err
	}
	return content, nil
}
//
//
func (r *repositoryContent) UpdateContent(content entity.Content) (entity.Content, error) {
	err := r.db.Save(&content).Error
	if err != nil {
		return content, err
	}
	return content, nil
}//
//
func (r *repositoryContent) DeleteContentById(content entity.Content) (int64, error) {
	result := r.db.Save(&content)
	err := result.Error
	rowsAffected := result.RowsAffected
	if err != nil {
		return  rowsAffected, err
	}
	return rowsAffected, nil
}//

//
func (r *repositoryContent) FindContentById(id int) (entity.ScanedContent, error) {
	var scannedContent entity.ScanedContent
	var content entity.Content
	var t time.Time
	err := r.db.Where("id = ?", id).Where("deleted_at = ?", t).Preload("ContentType").Find(&content).Scan(&scannedContent).Error
	if err != nil {
		return scannedContent, err
	}
	return scannedContent, nil
}
//

//
func (r *repositoryContent) FindAll() ([]entity.ScanedContent, error) {
	var scannedContent []entity.ScanedContent
	var contents []entity.Content
	var t time.Time
	err := r.db.Where("deleted_at = ?", t).Preload("ContentType").Find(&contents).Scan(&scannedContent).Error

	if err != nil {
		return scannedContent, err
	}
	return scannedContent, nil
}
//

//
func (r *repositoryContent) FindAllWithLimit(limit int) ([]entity.ScanedContent, error) {
	var scannedContent []entity.ScanedContent
	var contents []entity.Content
	var t time.Time
	err := r.db.Limit(limit).Where("deleted_at = ?", t).Preload("ContentType").Find(&contents).Scan(&scannedContent).Error
	if err != nil {
		return scannedContent, err
	}
	return scannedContent, nil
}
//