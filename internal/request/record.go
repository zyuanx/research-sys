package request

type RecordListQuery struct {
	Pagination
	ResearchID string `form:"researchID"`
}

type RecordCreatePayload struct {
	ResearchID int    `json:"researchID" binding:"required"`
	Values     string `json:"values" binding:"required"`
}

type OpenRecordListQuery struct {
	Pagination
	ResearchID int `form:"researchID" binding:"required"`
}

type OpenRecordCreatePayload struct {
	ResearchID int    `json:"researchID" binding:"required"`
	Values     string `json:"values" binding:"required"`
}
