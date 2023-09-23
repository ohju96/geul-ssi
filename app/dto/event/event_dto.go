package event

import "time"

func Dummy() {

}

type Event struct {
	User     string    `json:"user"`
	Password string    `json:"password"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

type GetEventResponse struct {
	User    string `json:"user"`
	Content string `json:"content"`
}
