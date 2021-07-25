package form

type RoleCreateForm struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Permissions []int  `json:"permissions"`
}

type RoleUpdateForm struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Permissions []int  `json:"permissions"`
}
