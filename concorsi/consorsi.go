package concorsi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tucnak/telebot"
)

func GetConc(myBot *telebot.Bot, m *telebot.Message) error {
	var s []string
	var concatena string
	var inizio, fine int
	url := "https://www.regione.vda.it/amministrazione/concorsi/default_i.asp"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	for _, i := range html {
		if string(i) == " " {
			concatena += string(i)
			s = append(s, concatena)
			concatena = ""
			continue
		}
		concatena += string(i)
	}
	for number, i := range s {
		if strings.Contains(i, "<caption>") {
			inizio = number
			break
		}
	}
	for number, i := range s {
		if strings.Contains(i, "</table>") {
			fine = number
			break
		}
	}
	var conc bytes.Buffer

	for i := inizio + 1; i < fine; i++ {
		if strings.HasPrefix(s[i], "<") && !strings.Contains(s[i], "href") {
			continue
		}

		s[i] = strings.ReplaceAll(s[i], "<td><a", "")
		s[i] = strings.ReplaceAll(s[i], "</caption>", "")
		s[i] = strings.ReplaceAll(s[i], "</th>", "")
		s[i] = strings.ReplaceAll(s[i], "</a></td>", "")
		if strings.Contains(s[i], "&nbsp") {
			s[i] = ""
		}
		s[i] = strings.ReplaceAll(s[i], "href=\"", "https://www.regione.vda.it")
		s[i] = strings.ReplaceAll(s[i], "\">", "\n")
		s[i] = strings.ReplaceAll(s[i], "pubblicazione", "")
		s[i] = strings.ReplaceAll(s[i], "presso", "")
		s[i] = strings.ReplaceAll(s[i], "l'amministrazione", "")
		s[i] = strings.ReplaceAll(s[i], "regionale", "")
		s[i] = strings.ReplaceAll(s[i], "di", "")
		s[i] = strings.ReplaceAll(s[i], "scadenza", "")

		_, err := conc.WriteString(s[i])
		if err != nil {
			return err
		}

	}
	myBot.Send(m.Chat, conc.String())
	return nil
}
