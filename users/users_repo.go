package users

import "go-simple/entity"

type UsersRepo interface {
	Login(email string, password string, users *entity.Users) (*entity.Users, error)
	Create(users *entity.Users) (*entity.Users, error)
	List(page, limit int, search, sort, sortField string) (*[]entity.Users, int64, error)
	Detail(id int) (*entity.Users, error)
	Update(id int, users map[string]interface{}) (*entity.Users, error)
	Delete(id int) error
}
