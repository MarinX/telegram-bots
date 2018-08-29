package main

import (
	"github.com/yanzay/tbot"
)

const apperInURL = "https://appear.in/"

// Handler
type Handler struct{}

// NewHandler creates new handler for telegram
func NewHandler() *Handler {
	return &Handler{}
}

// CreateRoom creates appear.in link
func (h *Handler) CreateRoom(message *tbot.Message) {
	room := message.Vars["name"]
	if len(room) <= 0 {
		//we dont have a room name, lets create generic one
		room = GenerateRandomString()
	}

	tmpl := "%s has created a room on appearIn.\nJoin here %s"

	message.Replyf(tmpl, message.From.FirstName, apperInURL+room)
}
