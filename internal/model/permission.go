package model

type Permission struct {
	BaseModel
	Group  string  `gorm:"size:32" json:"group"`
	Path   string  `gorm:"size:128" json:"path"`
	Method string  `gorm:"size:8" json:"method"`
	Desc   string  `gorm:"size:255" json:"desc"`
	Index  int     `json:"index"`
	Roles  []*Role `gorm:"many2many:role_permission" json:"-"`
}
