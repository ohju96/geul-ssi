package user

import "geulSsi/app/dto/custom"

type CreateUserRequest struct {
	Nickname string `json:"nickname" validate:"required" example:"nickname"`
	Password string `json:"password" validate:"required" sample:"password"`
}

type CreateUserResponse struct {
	custom.Response
	Payload CreateUserPayload `json:"payload"`
}

type CreateUserPayload struct {
	UserID     int    `json:"userID" example:"1"`
	SecretCode string `json:"secretCode" example:"uuid"`
}
