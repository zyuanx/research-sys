package model

type User struct {
	BaseModel
	Username string  `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Password string  `gorm:"size:255;not null" json:"-"`
	Email    string  `gorm:"size:255" json:"email"`
	Roles    []*Role `gorm:"many2many:user_role" json:"roles"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserChangePasswordReq struct {
	Password  string `json:"password" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

type UserListReq struct {
	Pagination
	Username string `form:"username"`
	Name     string `form:"name"`
}

type UserCreateReq struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
	Email     string `json:"email"`
	Roles     []int  `json:"roles"`
}

type UserUpdateReq struct {
	Email string `json:"email"`
	Roles []int  `json:"roles"`
}

type UserRes struct {
	ID       uint     `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}
