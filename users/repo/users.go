package repo

import (
	"fmt"
	"go-simple/entity"
	"go-simple/helper"
	"go-simple/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersRepoImpl struct {
	DB *gorm.DB
}

func CreateUsersRepo(DB *gorm.DB) users.UsersRepo {
	return &UsersRepoImpl{DB}
}

func (e *UsersRepoImpl) Login(email string, password string, users *entity.Users) (*entity.Users, error) {
	var upUsers = entity.Users{}

	err := e.DB.Table("com_user").Where("email = ?", email).First(&upUsers).Error
	if err != nil {
		return nil, fmt.Errorf("Email atau password salah !")
	}

	err = bcrypt.CompareHashAndPassword([]byte(upUsers.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Email atau password salah !")
	}

	return &upUsers, nil
}

func (e *UsersRepoImpl) Create(users *entity.Users) (*entity.Users, error) {
	// validate
	var count int64
	err := e.DB.Model(&entity.Users{}).Where("email = ?", users.Email).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf(helper.Existing(*users.Email))
	}

	err = e.DB.Save(users).Error
	if err != nil {
		fmt.Printf("[UsersRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.FailedToCreateData("Users"))
	}

	return users, nil
}

func (e *UsersRepoImpl) List(page, limit int, search, sort, sortField string) (*[]entity.Users, int64, error) {
	var users []entity.Users
	offset := (page - 1) * limit
	query := e.DB.Offset(offset).Limit(limit)

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if sort == "asc" {
		query = query.Order(sortField + " ASC")
	} else if sort == "desc" {
		query = query.Order(sortField + " DESC")
	}

	var totalData int64
	if err := query.Model(&entity.Users{}).Count(&totalData).Error; err != nil {
		fmt.Printf("[UsersRepoImpl.List] error counting data %v \n", err)
		return nil, 0, fmt.Errorf(helper.FailedToGetData("Users"))
	}

	err := query.Find(&users).Error
	if err != nil {
		fmt.Printf("[UsersRepoImpl.List] error execute query %v \n", err)
		return nil, 0, fmt.Errorf(helper.FailedToGetData("Users"))
	}
	return &users, totalData, nil
}

func (e *UsersRepoImpl) Detail(id int) (*entity.Users, error) {
	var users = entity.Users{}
	err := e.DB.Table("com_user").Where("user_id = ?", id).First(&users).Error
	if err != nil {
		fmt.Printf("[UsersRepoImpl.Detail] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.DataNotFound("Users"))
	}
	return &users, nil
}

func (e *UsersRepoImpl) Update(id int, users map[string]interface{}) (*entity.Users, error) {
	var newEmail string
	if val, ok := users["email"]; ok {
		newEmail = val.(string)
	}

	// validate mail if change
	if newEmail != "" {
		var count int64
		err := e.DB.Model(&entity.Users{}).Where("user_id != ? AND email = ?", id, newEmail).Count(&count).Error
		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf(helper.Existing(newEmail))
		}
	}

	var upUsers entity.Users
	err := e.DB.Table("com_user").Where("user_id = ?", id).Updates(users).Error
	if err != nil {
		fmt.Printf("[UsersRepoImpl.Update] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.FailedToUpdateData("Users"))
	}

	return &upUsers, nil
}

func (e *UsersRepoImpl) Delete(id int) error {
	var users = entity.Users{}
	err := e.DB.Table("com_user").Where("user_id = ?", id).First(&users).Delete(&users).Error
	if err != nil {
		fmt.Printf("[UsersRepoImpl.Delete] error execute query %v \n", err)
		return fmt.Errorf(helper.DataNotFound("Users"))
	}
	return nil
}
