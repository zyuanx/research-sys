package models

type Permission struct {
	BaseModel
	Path   string `gorm:"size:128"`
	Method string `gorm:"size:8"`
	Desc   string `gorm:"size:255"`
	Roles  []*Role `gorm:"many2many:role_permission"`
}
