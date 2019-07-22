package api

import (
	"github.com/MISTikus/ChatilGo/api/handlers"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	Rooms handlers.Rooms
	Users handlers.Users
}

func NewServer() server {
	return server{
		Rooms: handlers.RoomHandler(),
		Users: handlers.UserHandler(),
	}
}

func (s server) Start() {
	router := httprouter.New()
	router.GET("/contacts", s.Users.List)
	router.GET("/rooms", s.Rooms.List)
	router.POST("/rooms", s.Rooms.New)
}
