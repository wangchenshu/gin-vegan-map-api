package models

import (
	"gin-vegan-map-api/server/db"
	e "gin-vegan-map-api/server/enum"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// User -
type User struct {
	Model

	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Roles    string `json:"roles"`
	Password string `json:"-"`
}

// TableName -
func (User) TableName() string {
	return "users_go"
}

// ExistUserByName checks if there is a user with the same name
func ExistUserByName(name string) (bool, error) {
	var user User

	err := db.Db.Select("id").Where("name = ?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddUser Add a User
func AddUser(data User) error {
	if err := db.Db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// GetUsers - get all user
func GetUsers() ([]User, error) {
	var (
		users []User
		err   error
	)

	err = db.Db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// GetUsersByName - get user by name
func GetUsersByName(name string) ([]User, error) {
	var (
		users []User
		err   error
	)

	err = db.Db.Where("name = ?", name).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// GetUser - get user
func GetUser(id int) (User, error) {
	var (
		user User
		err  error
	)

	err = db.Db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

// GetUserByName - get user by name
func GetUserByName(name string) (User, error) {
	var (
		user User
		err  error
	)

	err = db.Db.Where("name = ?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

// GetUserTotal -
func GetUserTotal() (int, error) {
	var count int

	if err := db.Db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistUserByID -
func ExistUserByID(id int) (bool, error) {
	var user User

	err := db.Db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteUser delete a user
func DeleteUser(id int) error {
	if err := db.Db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser -
func UpdateUser(id int, data User) error {
	if err := db.Db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func containsRoles(rolesStr string, roles e.RolesEnum) bool {
	if strings.Contains(rolesStr, e.RolesEnum.String(e.RoleAdmin)) {
		// 用戶是最高權限, 都放行
		return true
	} else if strings.Contains(rolesStr, e.RolesEnum.String(e.RoleManagement)) {
		if roles == e.RoleManagement {
			return true
		}
	}
	return false
}

// CheckUserRoles -
func CheckUserRoles(userName string, roles e.RolesEnum) bool {
	if roles == e.RoleAdmin || roles == e.RoleManagement {
		user, err := GetUserByName(userName)
		if err != nil {
			log.Println(err)
			return false
		}

		if containsRoles(user.Roles, roles) {
			return true
		}
	}

	return false
}
