package models

type Permission struct {
	BaseModel
	Path   string  `gorm:"size:128" json:"path"`
	Method string  `gorm:"size:8" json:"method"`
	Desc   string  `gorm:"size:255" json:"desc"`
	Roles  []*Role `gorm:"many2many:role_permission" json:"-"`
}
