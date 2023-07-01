package chathelper

import (
	"fmt"

	"github.com/tucnak/telebot"
)

func GetChat(m *telebot.Message) (*telebot.Chat, error) {
	if m == nil {
		return nil, fmt.Errorf("no chat gettable")
	}
	return m.Chat, nil
}

func SpeakChat(myBot *telebot.Bot, m *telebot.Message, speakToChat *telebot.Chat) (err error) {
	_, err = myBot.Send(speakToChat, m.Text[7:])
	return
}
