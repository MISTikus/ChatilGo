package domain

type Room struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	OwnerId  string    `json:"ownerid"`
	Members  []string  `json:"members"`
	Messages []Message `json:"messages"`
}

func (r *Room) AddMember(memberId string) {
	r.Members = append(r.Members, memberId)
	// ToDo: inform all about adding new member
}

func (r *Room) AddMessage(message Message) {
	r.Messages = append(r.Messages, message)
	// ToDo: inform all about adding new message
}
