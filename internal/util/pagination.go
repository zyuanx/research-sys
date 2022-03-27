package util

import (
	"gorm.io/gorm"
)

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 999
		case size <= 0:
			size = 20
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
