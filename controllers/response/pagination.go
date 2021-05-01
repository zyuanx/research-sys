package response

type Pagination struct {
	Size    uint        `json:"size"`
	Page    uint        `json:"page"`
	Results interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total   uint        `json:"total"`
}
