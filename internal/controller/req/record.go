package req

type RecordCreateReq struct {
	Title       string                 `json:"title" binding:"required"`
	ResearchID  string                 `json:"researchID" binding:"required"`
	FieldsValue map[string]interface{} `json:"fieldsValue" binding:"required"`
}

type RecordUpdateReq struct {
	Title  string `json:"title" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Once   int    `json:"once"`
	Status int    `json:"status"`
}
