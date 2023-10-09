package entity

type Users struct {
	UserID   *int    `json:"user_id" gorm:"primaryKey"`
	RoleID   *int    `json:"role_id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password []byte  `json:"password"`
	File     string  `json:"file"`
	Token    *string `json:"token"`
}

type UsersRequest struct {
	RoleID   int    `json:"role_id" form:"role_id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (e *Users) TableName() string {
	return "com_user"
}
