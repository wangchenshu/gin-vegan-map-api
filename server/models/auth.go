package models

import (
	"gin-vegan-map-api/server/common"
	"gin-vegan-map-api/server/db"

	"github.com/jinzhu/gorm"
)

// CheckAuth checks if authentication information exists
func CheckAuth(name, password string) (bool, error) {
	var auth User

	err := db.Db.Where("name = ?", name).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		match := common.CheckPasswordHash(password, auth.Password)

		if match == true {
			return true, nil
		}
	}

	return false, nil
}
