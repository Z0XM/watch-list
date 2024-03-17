package requests

import "github.com/projects/watch-list/server/core/user/service/requests"

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req *LoginRequest) ToSvcRequest() requests.SvcLoginRequest {
	return requests.SvcLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}
