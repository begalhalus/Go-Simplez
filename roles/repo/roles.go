package repo

import (
	"fmt"
	"go-simple/entity"
	"go-simple/helper"
	"go-simple/roles"

	"gorm.io/gorm"
)

type RolesRepoImpl struct {
	DB *gorm.DB
}

func CreateRolesRepo(DB *gorm.DB) roles.RolesRepo {
	return &RolesRepoImpl{DB}
}

func (e *RolesRepoImpl) Create(roles *entity.Roles) (*entity.Roles, error) {
	// validate
	var count int64
	err := e.DB.Model(&entity.Roles{}).Where("role_nm = ?", roles.RoleNm).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf(helper.Existing(*roles.RoleNm))
	}

	err = e.DB.Save(roles).Error
	if err != nil {
		fmt.Printf("[RolesRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.FailedToCreateData("Role"))
	}
	return roles, nil
}

func (e *RolesRepoImpl) List(page, limit int, search, sort, sortField string) (*[]entity.Roles, int64, error) {
	var roles []entity.Roles
	offset := (page - 1) * limit
	query := e.DB.Offset(offset).Limit(limit)

	if search != "" {
		query = query.Where("role_nm LIKE ?", "%"+search+"%")
	}

	if sort == "asc" {
		query = query.Order(sortField + " ASC")
	} else if sort == "desc" {
		query = query.Order(sortField + " DESC")
	}

	var totalData int64
	if err := query.Model(&entity.Roles{}).Count(&totalData).Error; err != nil {
		fmt.Printf("[RolesRepoImpl.List] error counting data %v \n", err)
		return nil, 0, fmt.Errorf(helper.FailedToGetData("Role"))
	}

	err := query.Find(&roles).Error
	if err != nil {
		fmt.Printf("[RolesRepoImpl.List] error execute query %v \n", err)
		return nil, 0, fmt.Errorf(helper.FailedToGetData("Role"))
	}
	return &roles, totalData, nil
}

func (e *RolesRepoImpl) Detail(id int) (*entity.Roles, error) {
	var roles = entity.Roles{}
	err := e.DB.Table("com_role").Where("role_id = ?", id).First(&roles).Error
	if err != nil {
		fmt.Printf("[RolesRepoImpl.Detail] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.DataNotFound("Role"))
	}
	return &roles, nil
}

func (e *RolesRepoImpl) Update(id int, roles map[string]interface{}) (*entity.Roles, error) {
	var newName string
	if val, ok := roles["role_nm"]; ok {
		newName = val.(string)
	}

	// validate name if change
	if newName != "" {
		var count int64
		err := e.DB.Model(&entity.Roles{}).Where("role_id != ? AND role_nm = ?", id, newName).Count(&count).Error
		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf(helper.Existing(newName))
		}
	}

	var upRoles entity.Roles
	err := e.DB.Table("com_role").Where("role_id = ?", id).Updates(roles).Error
	if err != nil {
		fmt.Printf("[RolesRepoImpl.Update] error execute query %v \n", err)
		return nil, fmt.Errorf(helper.FailedToUpdateData("Role"))
	}

	return &upRoles, nil
}

func (e *RolesRepoImpl) Delete(id int) error {
	var roles = entity.Roles{}
	err := e.DB.Table("com_role").Where("role_id = ?", id).First(&roles).Delete(&roles).Error
	if err != nil {
		fmt.Printf("[RolesRepoImpl.Delete] error execute query %v \n", err)
		return fmt.Errorf(helper.DataNotFound("Role"))
	}
	return nil
}
