package domain

import "time"

type Message struct {
	Created  time.Time `json:"created"`
	AuthorId string    `json:"authorid"`
	Text     string    `json:"text"`
}

type Messages interface {
	List() []string
}
