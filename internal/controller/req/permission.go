package req

type PermissionUpdateReq struct {
	Group  string `json:"group" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Index  int    `json:"index" binding:"required"`
}
