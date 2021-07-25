package form

type Pagination struct {
	Size int `form:"size,default=10"`
	Page int `form:"page,default=1"`
}
