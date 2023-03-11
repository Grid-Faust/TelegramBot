package main

import (
	"flag"
	"log"
	"telegrambot/clients/telegram"
)

const (
	thBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(token())



}

func token() string {
	// bot -tg-bot-token 'my token'
	
	token := flag.String(
		"token " ,
		 "", 
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}