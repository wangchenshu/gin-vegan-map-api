package models

import (
	"gin-vegan-map-api/server/db"

	"github.com/jinzhu/gorm"
)

// Pics -
type Pics struct {
	Model

	Name     string `json:"name"`
	Regional string `json:"regional"`
	URL      string `json:"url"`
	Type     string `json:"type"`
}

// GetPics - get all pics
func GetPics() ([]Pics, error) {
	var (
		pics []Pics
		err  error
	)

	err = db.Db.Find(&pics).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return pics, nil
}

// GetPicsByTypeAndRegional -
func GetPicsByTypeAndRegional(typeName string, regional string) ([]Pics, error) {
	var (
		pics []Pics
		err  error
	)

	err = db.Db.Where("type = ? AND regional = ?", typeName, regional).Find(&pics).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return pics, nil
}

// GetPicsMessageByTypeAndRegional -
func GetPicsMessageByTypeAndRegional(typeName string, regional string) (Messages, error) {
	pics, err := GetPicsByTypeAndRegional(typeName, regional)
	if err != nil {
		return Messages{}, err
	}

	return getGalleryMessages(pics)
}
