package event

import "geulSsi/app/dto/custom"

func Dummy() {

}

type Event struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Content  string `json:"content"`
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
