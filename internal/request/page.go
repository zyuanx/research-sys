package request

type Pagination struct {
	Size int `form:"size,default=20"`
	Page int `form:"page,default=1"`
}
