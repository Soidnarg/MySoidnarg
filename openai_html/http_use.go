package openai_html

import (
	"encoding/json"
	"net/http"

	"github.com/Grandioso99/adminBot/openai_local"
)

type MessageJSON struct {
	Text string `json:"text"`
	Key  string `json:"key,omitempty"`
}

type ResponseJSON struct {
	Text string `json:"text"`
}

/*
	func (buffJSON MessageJSON) requestToJSON(r io.Reader) (err error) {
		var buffAppener []byte
		buff := make([]byte, 50)

		for n, err := r.Read(buff); n != 0; {
			if err != nil {
				log.Println(err)
			}

			buffAppener = append(buffAppener, buff...)
		}
		err = json.Unmarshal(buffAppener, buffJSON)

		return
	}
*/
func RequestToGPT(w http.ResponseWriter, r *http.Request) (err error) {
	var m MessageJSON
	var resp ResponseJSON

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return err
	}

	_, client := openai_local.NewOpenAI(m.Key)

	resp.Text, err = openai_local.AskToChatGPT(client, m.Text)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}

func RequestToDaVinci(w http.ResponseWriter, r *http.Request) (err error) {
	var m MessageJSON
	var resp ResponseJSON

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return err
	}

	_, client := openai_local.NewOpenAI(m.Key)

	resp.Text, err = openai_local.AskToDaVinci(client, m.Text)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}

func RequestToAda(w http.ResponseWriter, r *http.Request) (err error) {
	var m MessageJSON
	var resp ResponseJSON

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return err
	}

	_, client := openai_local.NewOpenAI(m.Key)

	resp.Text, err = openai_local.AskToDumbAda(client, m.Text)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}

func RequestToGPT4(w http.ResponseWriter, r *http.Request) (err error) {
	var m MessageJSON
	var resp ResponseJSON

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return err
	}

	_, client := openai_local.NewOpenAI(m.Key)

	resp.Text, err = openai_local.AskToGPT4(client, m.Text)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}

func RequestToImageCreator(w http.ResponseWriter, r *http.Request) (err error) {
	var m MessageJSON
	var resp ResponseJSON

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return err
	}

	_, client := openai_local.NewOpenAI(m.Key)

	resp.Text, err = openai_local.AskToImageGPT(client, m.Text)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}
