package usecase

import (
	"go-simple/entity"
	"go-simple/roles"
)

type RolesUsecaseImpl struct {
	rolesRepo roles.RolesRepo
}

func CreateRolesUsecase(rolesRepo roles.RolesRepo) roles.RolesUsecase {
	return &RolesUsecaseImpl{rolesRepo}
}

func (e *RolesUsecaseImpl) Create(roles *entity.Roles) (*entity.Roles, error) {
	return e.rolesRepo.Create(roles)
}

func (e *RolesUsecaseImpl) List(page, limit int, search, sort, sortField string) (*[]entity.Roles, int64, error) {
	return e.rolesRepo.List(page, limit, search, sort, sortField)
}

func (e *RolesUsecaseImpl) Detail(id int) (*entity.Roles, error) {
	return e.rolesRepo.Detail(id)
}

func (e *RolesUsecaseImpl) Update(id int, roles map[string]interface{}) (*entity.Roles, error) {
	return e.rolesRepo.Update(id, roles)
}

func (e *RolesUsecaseImpl) Delete(id int) error {
	return e.rolesRepo.Delete(id)
}
