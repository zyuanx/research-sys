package form

type RecordCreateForm struct {
	Title       string                 `json:"title" binding:"required"`
	ResearchID  string                 `json:"researchID" binding:"required"`
	FieldsValue map[string]interface{} `json:"fieldsValue" binding:"required"`
}

type RecordUpdateForm struct {
	Status int `json:"status" binding:"required"`
}
