package req

type RoleUpdateReq struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Permissions []int  `json:"permissions"`
}
