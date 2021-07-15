package res

import (
	"gin-research-sys/internal/model"
)

type UserInfoResponse struct {
	BaseData
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
}

func InfoSerializer(user model.User) UserInfoResponse {
	return UserInfoResponse{
		Username:  user.Username,
		Telephone: user.Telephone,
		BaseData: BaseData{
			Id:        user.ID,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},

	}
}
