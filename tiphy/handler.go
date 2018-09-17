package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/yanzay/tbot"
)

const giphySearch = "https://api.giphy.com/v1/gifs/search?api_key=3eFQvabDx69SMoOemSPiYfh9FY0nzO9x&q=%s&offset=0&limit=10"

// Handler
type Handler struct{}

// NewHandler creates new handler for telegram
func NewHandler() *Handler {
	rand.Seed(time.Now().Unix())
	return &Handler{}
}

// Gif replays with gif from giphy
func (h *Handler) Gif(message *tbot.Message) {
	query := message.Vars["query"]
	if len(query) <= 0 {
		message.Reply("Please provide query for searching")
		return
	}
	message.Replyf(h.getGif(query))
}

func (h *Handler) getGif(search string) string {
	resp, err := http.Get(fmt.Sprintf(giphySearch, search))
	if err != nil {
		return fmt.Sprintf("Error getting gif %s", err.Error())
	}
	var gr GiphyResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		return fmt.Sprintf("Error decoding gifs %s", err.Error())
	}

	if len(gr.Data) <= 0 {
		return fmt.Sprintf("No gifs matched by query :(")
	}

	n := rand.Int() % len(gr.Data)
	return gr.Data[n].Url
}
