package models

import (
	"gin-vegan-map-api/server/db"
	e "gin-vegan-map-api/server/enum"

	"github.com/jinzhu/gorm"
)

// Restaurant -
type Restaurant struct {
	Model

	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	BusinessHour string `json:"business_hour"`
	Comment      string `json:"comment"`
}

// ExistRestaurantByName checks if there is a restaurant with the same name
func ExistRestaurantByName(name string) (bool, error) {
	var restaurant Restaurant
	err := db.Db.Select("id").Where("name = ?", name).First(&restaurant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if restaurant.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddRestaurant Add a Restaurant
func AddRestaurant(data interface{}) error {

	if err := db.Db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// GetRestaurants - get all restaurant
func GetRestaurants() ([]Restaurant, error) {
	var (
		restaurants []Restaurant
		err         error
	)

	err = db.Db.Find(&restaurants).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return restaurants, nil
}

// GetRestaurantsByTypeAndRegional - get restaurant by type and regional
func GetRestaurantsByTypeAndRegional(typeName string, regional string) ([]Restaurant, error) {
	var (
		restaurants []Restaurant
		err         error
	)

	err = db.Db.Where("type = ? AND regional = ?", typeName, regional).Find(&restaurants).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return restaurants, nil
}

// GetRestaurantsByFriedRegional - get restaurant by regional
func GetRestaurantsByFriedRegional(regional string) ([]Restaurant, error) {
	var (
		restaurants []Restaurant
		err         error
	)

	err = db.Db.Where("regional = ? AND type = ?", regional, e.FriedCHTName).Find(&restaurants).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return restaurants, nil
}

// GetRestaurant - get restaurant
func GetRestaurant(id int) (Restaurant, error) {
	var (
		restaurant Restaurant
		err        error
	)

	err = db.Db.Where("id = ?", id).First(&restaurant).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return restaurant, err
	}

	return restaurant, nil
}

// GetRestaurantTotal -
func GetRestaurantTotal() (int, error) {
	var count int
	if err := db.Db.Model(&Restaurant{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistRestaurantByID -
func ExistRestaurantByID(id int) (bool, error) {
	var restaurant Restaurant
	err := db.Db.Select("id").Where("id = ?", id).First(&restaurant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if restaurant.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteRestaurant delete a restaurant
func DeleteRestaurant(id int) error {
	if err := db.Db.Where("id = ?", id).Delete(&Restaurant{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateRestaurant -
func UpdateRestaurant(id int, data interface{}) error {
	if err := db.Db.Model(&Restaurant{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
