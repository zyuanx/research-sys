package request

type OpenRecordListQuery struct {
	Pagination
	ResearchID int `form:"researchID" binding:"required"`
}

type OpenRecordCreatePayload struct {
	ResearchID int    `json:"researchID" binding:"required"`
	Values     string `json:"values" binding:"required"`
}
