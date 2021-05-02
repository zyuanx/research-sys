package res

type RoleListResponse struct {
	BaseData
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

//func RoleListSerializer(role []models.Role) []RoleListResponse {
//	var rlr []RoleListResponse
//	for
//	return RoleListResponse{}
//}
