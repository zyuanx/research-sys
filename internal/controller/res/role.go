package res

type RoleResponse struct {
	BaseData
	Title string `json:"title"`
	Desc  string `json:"desc"`
}