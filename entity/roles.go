package entity

type Roles struct {
	RoleID   *int    `json:"role_id" gorm:"primaryKey"`
	RoleNm   *string `json:"role_nm"`
	RoleDesc *string `json:"role_desc"`
}

type RolesRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

func (e *Roles) TableName() string {
	return "com_role"
}
