package req

type CreateRoleValidate struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
}


type RoleUpdateValidate struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
