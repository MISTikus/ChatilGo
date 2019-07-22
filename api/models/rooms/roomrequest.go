package models

type NewRoomRequest struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
	OwnerId string   `json:"ownerid"`
}
