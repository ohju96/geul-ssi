package event

import "geulSsi/app/dto/custom"

func Dummy() {

}

type Event struct {
	User     string `json:"user"`
	Content  string `json:"content"`
	Password string `json:"password"`
	EventId  int    `json:"eventId"`
}

type GetEventResponse struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

type CreateEventPayload struct {
	EventID  int    `json:"eventID" example:"1"`
	NickName string `json:"nickName" example:"nickname"`
}

type CreateEventResponse struct {
	Payload CreateEventPayload `json:"payload"`
	custom.Response
}

type GetEventListPayload struct {
	EventID       int    `json:"eventID" example:"1"`
	EventContents string `json:"eventContents" example:"eventContents"`
	NickName      string `json:"nickName" example:"nickname"`
	HeartCount    int    `json:"heartCount" example:"1"`
}

type GetEventListResponse struct {
	GetEventListPayload []GetEventListPayload `json:"payload"`
	custom.Response
}

type GetEventReq struct {
	NickName string `json:"nickName" example:"nickname"`
	Password string `json:"passWord" example:"password"`
}

type CreateHeartReq struct {
	EventId  int    `json:"eventId" example:"1"`
	NickName string `json:"nickName" example:"nickname"`
	Password string `json:"passWord" example:"password"`
}

type CreateHeartResponse struct {
	custom.Response
}
