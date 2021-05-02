package req

type PaginationQuery struct {
	Size int `form:"size"`
	Page int `form:"page"`
}
