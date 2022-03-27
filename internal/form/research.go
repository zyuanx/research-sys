package form

type ResearchListForm struct {
	Pagination
	Access string `json:"access" binding:"required"`
}
type ResearchCreateForm struct {
	Title       string                   `json:"title" binding:"required"`
	Desc        string                   `json:"desc" binding:"required"`
	Access      string                   `json:"access" binding:"required"`
	Once        int                      `json:"once"`
	FieldsValue map[string]interface{}   `json:"fieldsValue" binding:"required"`
	Detail      []map[string]interface{} `json:"detail" binding:"required"`
	Rules       map[string]interface{}   `json:"rules" binding:"required"`
}

type ResearchUpdateForm struct {
	Title  string `json:"title" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Once   int    `json:"once"`
	Status int    `json:"status"`
}
