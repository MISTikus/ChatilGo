package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MISTikus/ChatilGo/domain"
	"github.com/julienschmidt/httprouter"
)

type users struct {
	Users []domain.User
}

func UserHandler() *users {
	return &users{}
}

type Users interface {
	List(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (h *users) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	b, err := json.Marshal(h.Users)
	if err != nil {
		log.Println(err)
	}
	writer.Write(b)
	writer.WriteHeader(http.StatusOK)
}
