package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	models "github.com/MISTikus/ChatilGo/api/models/rooms"
	"github.com/MISTikus/ChatilGo/domain"
	"github.com/MISTikus/ChatilGo/helpers"
	"github.com/julienschmidt/httprouter"
)

type rooms struct {
	Rooms []domain.Room
}

func RoomHandler() *rooms {
	return &rooms{}
}

type Rooms interface {
	List(http.ResponseWriter, *http.Request, httprouter.Params)
	New(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (h *rooms) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	b, err := json.Marshal(h.Rooms)
	if err != nil {
		log.Println(err)
	}
	writer.Write(b)
	writer.WriteHeader(http.StatusOK)
}

func (h *rooms) New(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	r := models.NewRoomRequest{}
	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Print(err)
		return
	}
	id, err := helpers.NewId()
	h.Rooms = append(h.Rooms, domain.Room{
		Id:      id,
		Name:    r.Name,
		OwnerId: r.OwnerId,
		Members: r.Members,
	})
	writer.WriteHeader(http.StatusCreated)
}
