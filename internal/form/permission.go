package form

type PermissionCreateForm struct {
	Group  string `json:"group" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Index  int    `json:"index" binding:"required"`
}
type PermissionUpdateForm struct {
	Group  string `json:"group" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Index  int    `json:"index" binding:"required"`
}
