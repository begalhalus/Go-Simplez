package usecase

import (
	"go-simple/entity"
	"go-simple/users"
)

type UsersUsecaseImpl struct {
	usersRepo users.UsersRepo
}

func CreateUsersUsecase(usersRepo users.UsersRepo) users.UsersUsecase {
	return &UsersUsecaseImpl{usersRepo}
}

func (e *UsersUsecaseImpl) Login(email string, password string, users *entity.Users) (*entity.Users, error) {
	return e.usersRepo.Login(email, password, users)
}

func (e *UsersUsecaseImpl) Create(users *entity.Users) (*entity.Users, error) {
	return e.usersRepo.Create(users)
}

func (e *UsersUsecaseImpl) List(page, limit int, search, sort, sortField string) (*[]entity.Users, int64, error) {
	return e.usersRepo.List(page, limit, search, sort, sortField)
}

func (e *UsersUsecaseImpl) Detail(id int) (*entity.Users, error) {
	return e.usersRepo.Detail(id)
}

func (e *UsersUsecaseImpl) Update(id int, users map[string]interface{}) (*entity.Users, error) {
	return e.usersRepo.Update(id, users)
}

func (e *UsersUsecaseImpl) Delete(id int) error {
	return e.usersRepo.Delete(id)
}
