package openai_telebot

import (
	"github.com/Grandioso99/adminBot/openai_local"
	"github.com/tucnak/telebot"
)

func ChattoGPT(myBot *telebot.Bot, m *telebot.Message) error {
	testo := m.Text[10:]
	n, client := openai_local.NewOpenAI(testo)

	myBot.Send(m.Chat, testo[n:])
	resp, err := openai_local.AskToChatGPT(client, testo[n:])
	if err != nil {
		return err
	}

	myBot.Send(m.Chat, resp)

	return nil
}

func ChattoDaVinci(myBot *telebot.Bot, m *telebot.Message) error {
	testo := m.Text[9:]
	n, client := openai_local.NewOpenAI(testo)

	myBot.Send(m.Chat, testo[n:])
	resp, err := openai_local.AskToDaVinci(client, testo[n:])
	if err != nil {
		return err
	}

	myBot.Send(m.Chat, resp)

	return nil
}

func ChattoAda(myBot *telebot.Bot, m *telebot.Message) error {
	testo := m.Text[5:]
	n, client := openai_local.NewOpenAI(testo)

	myBot.Send(m.Chat, testo[n:])
	resp, err := openai_local.AskToDumbAda(client, testo[n:])
	if err != nil {
		myBot.Send(m.Chat, err)
		return err
	}

	myBot.Send(m.Chat, resp)

	return nil
}

func ChattoGPT4(myBot *telebot.Bot, m *telebot.Message) error {
	testo := m.Text[6:]
	n, client := openai_local.NewOpenAI(testo)

	myBot.Send(m.Chat, testo[n:])
	resp, err := openai_local.AskToGPT4(client, testo[n:])
	if err != nil {
		myBot.Send(m.Chat, err)
		return err
	}

	myBot.Send(m.Chat, resp)
	return nil
}

func ChattoImageCreator(myBot *telebot.Bot, m *telebot.Message) error {
	testo := m.Text[10:]
	n, client := openai_local.NewOpenAI(testo)

	myBot.Send(m.Chat, testo[n:])
	resp, err := openai_local.AskToImageGPT(client, testo[n:])
	if err != nil {
		myBot.Send(m.Chat, err)
		return err
	}

	myBot.Send(m.Chat, resp)

	return nil
}
