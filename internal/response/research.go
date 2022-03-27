package response

type ResearchResponse struct {
	BaseModel
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Access     string `json:"access"`
	Once       int    `json:"once"`
	ResearchID string `json:"researchID"`
}
