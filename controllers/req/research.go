package req

type ResearchCreateReq struct {
	Title       string                   `json:"title" binding:"required"`
	Desc        string                   `json:"desc" binding:"required"`
	Once        int                      `json:"once"`
	FieldsValue map[string]interface{}   `json:"fieldsValue" binding:"required"`
	Detail      []map[string]interface{} `json:"detail" binding:"required"`
	Rules       map[string]interface{}   `json:"rules" binding:"required"`
}
