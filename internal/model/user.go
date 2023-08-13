package model

type User struct {
	BaseModel
	Username  string  `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Gender    int     `gorm:"size:8" json:"gender"`
	Nickname  string  `gorm:"size:64" json:"nickname"`
	Password  string  `gorm:"size:255;not null" json:"-"`
	Telephone string  `gorm:"size:11" json:"telephone"`
	Email     string  `gorm:"size:255" json:"email"`
	Roles     []*Role `gorm:"many2many:user_role" json:"roles"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
