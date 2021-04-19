package response

type InfoSerializer struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
}
