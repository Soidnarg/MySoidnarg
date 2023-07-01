package main

import (
	"log"
	"time"

	"github.com/Grandioso99/adminBot/openai_html"
	"github.com/Grandioso99/adminBot/openai_telebot"
	"github.com/Grandioso99/adminBot/utils"
)

func main() {
	openai_telebot.StartTelebot(
		utils.GetToken("telegram"),
	)
	openai_html.StartHttp()

	for {
		time.Sleep(15 * time.Minute)
		log.Printf("checking hour: %v", time.Now())
	}
}
