package openai_telebot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	chathelper "github.com/Grandioso99/adminBot/chat_helper"
	"github.com/Grandioso99/adminBot/concorsi"
	"github.com/Grandioso99/adminBot/openai_local"
	textsaver "github.com/Grandioso99/adminBot/text_saver"
	"github.com/tucnak/telebot"
)

func StartTelebot(idToken string) {
	myBot, err := telebot.NewBot(telebot.Settings{
		Token:  idToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	myBot.Handle(telebot.OnAddedToGroup, func(m *telebot.Message) {
		firstTime, err := os.ReadFile("resources//imhere.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		myBot.Send(m.Chat, string(firstTime[:]))
	})

	myBot.Handle(telebot.OnUserJoined, func(m *telebot.Message) {
		welcome, err := os.ReadFile("resources//welcome.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		myBot.Send(m.Chat, m.UserJoined.FirstName+string(welcome[:]))
	})

	myBot.Handle(telebot.OnUserLeft, func(m *telebot.Message) {
		myBot.Send(m.Chat, m.UserLeft.FirstName+" se n'è andato, ci mancherà forse")
	})

	myBot.Handle("/start", func(m *telebot.Message) {
		myBot.Delete(m)
		myBot.Send(m.Chat, "MySoidnargBot attivo, piacere di conoscerti "+m.Sender.FirstName)
		myBot.Send(m.Chat, "Usa il commando /commands per avere una lista dettagliata dei comandi")
		myBot.Send(m.Chat, "In caso di problemi rivolgiti al mio creatore @Grandios99 maestro di gilda")
	})

	myBot.Handle("/commands", func(m *telebot.Message) {
		commands, err := os.ReadFile("resources//commands.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		myBot.Send(m.Chat, string(commands[:]))
	})

	myBot.Handle("/quest", func(m *telebot.Message) {
		myBot.Delete(m)
		myBot.Send(m.Chat, "Comando quest in manutenzione finché non mi andrà di sistemarlo")
		/*
			var user string
			err := ioutil.WriteFile("resources//questUsers.txt", []byte(user), 0777)
			if err != nil {
				log.Fatal(err)
				return
			}
		*/
	})

	myBot.Handle("/rules", func(m *telebot.Message) {
		myBot.Delete(m)
		rules, err := os.ReadFile("resources//rules.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		myBot.Send(m.Chat, string(rules[:]))
	})

	myBot.Handle("/fban", func(m *telebot.Message) {
		myBot.Delete(m)
		option := telebot.SendOptions{ReplyTo: m.ReplyTo, DisableWebPagePreview: false, DisableNotification: false, ParseMode: ""}
		file := telebot.File{
			FileLocal: "resources//FAKEUntitled.mp4",
		}
		video := &telebot.Video{File: file, Width: 1920, Height: 1080}
		_, err := video.Send(myBot, m.Chat, &option)
		if err != nil {
			myBot.Reply(m.ReplyTo, "Nope "+err.Error())
		}
	})

	myBot.Handle("/news", func(m *telebot.Message) {
		myBot.Delete(m)
		err := textsaver.ShowNews(myBot, m)
		if err != nil {
			log.Println(err)
			return
		}
	})

	myBot.Handle("/newNews", func(m *telebot.Message) {
		myBot.Delete(m)

		err := textsaver.NewNews(myBot, m)
		if err != nil {
			log.Println(err)
			return
		}
	})

	var speakToChat *telebot.Chat
	myBot.Handle("/getChat", func(m *telebot.Message) {
		myBot.Delete(m)

		speakToChat, err = chathelper.GetChat(m)
		if err != nil {
			log.Println(err)
		}
	})

	myBot.Handle("/speak", func(m *telebot.Message) {
		myBot.Delete(m)

		err = chathelper.SpeakChat(myBot, m, speakToChat)
		if err != nil {
			log.Println(err)
		}
	})

	myBot.Handle("/ita", func(m *telebot.Message) {
		tradotto := make([]string, 1024)

		if !m.IsReply() {
			myBot.Send(m.Chat, "Digitare /ita in risposta ad un messaggio per rispondere")
			return
		}
		myBot.Send(m.Chat, m.ReplyTo.Text)
		url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=it&dt=t&text=" + m.ReplyTo.Text
		myBot.Send(m.Chat, url)
		resp, err := http.Get(url)
		if err != nil {
			myBot.Send(m.Chat, "No page "+err.Error())
			return
		}

		read, err := io.ReadAll(resp.Body)
		if err != nil {
			myBot.Send(m.Chat, "Traduzione fallita "+err.Error())
			return
		}
		json.Unmarshal(read, &tradotto)
		_, err = myBot.Send(m.Chat, tradotto[:])
		if err != nil {
			myBot.Send(m.Chat, "Traduzione fallita "+err.Error())
			return
		}
	})

	// Cosa private Ruben
	myBot.Handle("/concorsi", func(m *telebot.Message) {
		myBot.Delete(m)
		concorsi.GetConc(myBot, m)
	})

	myBot.Handle("/chiedimi", func(m *telebot.Message) {
		myBot.Delete(m)
		ChattoGPT(myBot, m)
	})

	myBot.Handle("/davinci", func(m *telebot.Message) {
		myBot.Delete(m)
		ChattoDaVinci(myBot, m)
	})

	myBot.Handle("/ada", func(m *telebot.Message) {
		myBot.Delete(m)
		ChattoAda(myBot, m)
	})

	myBot.Handle("/gpt4", func(m *telebot.Message) {
		myBot.Delete(m)
		ChattoGPT4(myBot, m)
	})

	myBot.Handle("/immagina", func(m *telebot.Message) {
		myBot.Delete(m)
		ChattoImageCreator(myBot, m)
	})

	myBot.Handle("/fixKey", func(m *telebot.Message) {
		myBot.Delete(m)
		openai_local.TempoKey = m.Text[8:]
	})

	myBot.Handle("/deleteKey", func(m *telebot.Message) {
		myBot.Delete(m)
		openai_local.TempoKey = ""
	})

	go myBot.Start()
}
