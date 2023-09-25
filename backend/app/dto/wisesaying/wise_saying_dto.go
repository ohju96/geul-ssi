package wisesaying

type AddWiseSayingRequest struct { // 명언 추가 요청
	WiseSaying string `json:"wiseSaying" validate:"required"`
	Nickname   string `json:"nickname" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
