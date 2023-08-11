package request

type UserListQuery struct {
	Pagination
	Username string `form:"username"`
	Name     string `form:"name"`
}
