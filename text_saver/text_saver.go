package textsaver

import (
	"io/ioutil"

	"github.com/tucnak/telebot"
)

func ShowNews(myBot *telebot.Bot, m *telebot.Message) error {
	news, err := ioutil.ReadFile("resources//news.txt")
	if err != nil {
		return err
	}
	myBot.Send(m.Chat, string(news[:]))
	return nil
}

func NewNews(myBot *telebot.Bot, m *telebot.Message) error {
	err := ioutil.WriteFile("resources//news.txt", []byte(m.Text[9:]), 0777)
	if err != nil {
		return err
	}

	news, err := ioutil.ReadFile("resources//news.txt")
	if err != nil {
		return err
	}
	myBot.Send(m.Chat, string(news[:]))

	return nil
}
