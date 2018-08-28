package main

import (
	"log"
	"os"

	"github.com/yanzay/tbot"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) <= 0 {
		log.Println("Telegram token is missing from env")
		return
	}

	if err := initServer(token); err != nil {
		log.Println(err)
	}
}

func initServer(token string) error {
	bot, err := tbot.NewServer(token)
	if err != nil {
		return err
	}

	handler := NewHandler()

	bot.HandleFunc("/wallet {wallet}", handler.GetWalletInfo)
	bot.HandleFunc("/payment {wallet}", handler.GenerateQR)
	bot.HandleFunc("/latest_block", handler.GetLatestBlock)

	return bot.ListenAndServe()
}
