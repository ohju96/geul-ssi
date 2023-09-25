package user

import "geulSsi/app/dto/custom"

type ChangePasswordRequest struct { // 비밀번호 변경
	Nickname   string `json:"nickname" validate:"required" example:"nickname"`
	Password   string `json:"password" validate:"required" example:"password"`
	SecretCode string `json:"secretCode" validate:"required" example:"uuid"`
}

type CreateUserRequest struct { // 회원가입
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
