package main

import (
	"fmt"

	"github.com/xorcare/blockchain"
	"github.com/yanzay/tbot"
)

// Handler holds blockchain client
type Handler struct {
	client *blockchain.Client
}

// NewHandler creates new handler for telegram
func NewHandler() *Handler {
	return &Handler{
		client: blockchain.New(),
	}
}

// GetWalletInfo responds with basic wallet - total recv, sent and final balance
func (h *Handler) GetWalletInfo(message *tbot.Message) {
	wallet := message.Vars["wallet"]
	if len(wallet) <= 0 {
		message.Reply("Missing wallet address.")
		return
	}

	addr, err := h.client.GetAddress(wallet)
	if err != nil {
		message.Replyf("Error obtaining address info %s", err.Error())
		return
	}

	tmpl := "Total received %s BTC\nTotal sent %s BTC\nFinal balance %s BTC"
	message.Replyf(tmpl,
		FormatBTC(addr.TotalReceived),
		FormatBTC(addr.TotalSent),
		FormatBTC(addr.FinalBalance),
	)

}

// GenerateQR generates QR code url
func (h *Handler) GenerateQR(message *tbot.Message) {
	wallet := message.Vars["wallet"]
	if len(wallet) <= 0 {
		message.Reply("Missing wallet address")
		return
	}
	message.Reply(fmt.Sprintf("https://blockchain.info/qr?data=%s&size=200", wallet))
}

// GetLatestBlock receive the latest block of the main chain
func (h *Handler) GetLatestBlock(message *tbot.Message) {
	block, err := h.client.GetLatestBlock()
	if err != nil {
		message.Replyf("Error getting latest block %s", err.Error())
		return
	}
	message.Replyf("Latest block hash %s", block.Hash)
}
