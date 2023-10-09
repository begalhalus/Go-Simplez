package roles

import "go-simple/entity"

type RolesRepo interface {
	Create(roles *entity.Roles) (*entity.Roles, error)
	List(page, limit int, search, sort, sortField string) (*[]entity.Roles, int64, error)
	Detail(id int) (*entity.Roles, error)
	Update(id int, roles map[string]interface{}) (*entity.Roles, error)
	Delete(id int) error
}
