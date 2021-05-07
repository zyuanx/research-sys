package req

type CreateRoleValidate struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
}

type RoleUpdateReq struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Permissions []int  `json:"permissions"`
}
