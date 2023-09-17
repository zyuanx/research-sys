package model

type Role struct {
	BaseModel
	Title       string  `gorm:"size:32;not null;uniqueIndex;" json:"title"`
	Description string  `gorm:"size:255;" json:"description"`
	Users       []*User `gorm:"many2many:user_role;" json:"-"`
}

type RoleUpdateReq struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type RoleRes struct {
}
