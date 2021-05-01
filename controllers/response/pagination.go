package response

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Pagination struct {
	Size    uint        `json:"size"`
	Page    uint        `json:"page"`
	Results interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total   uint        `json:"total"`
}


func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}