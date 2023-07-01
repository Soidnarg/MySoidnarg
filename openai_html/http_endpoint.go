package openai_html

import (
	"log"
	"net/http"
	"os"
)

func StartHttp() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/chiedimi", func(w http.ResponseWriter, r *http.Request) {
		err = RequestToGPT(w, r)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("richiesta chatgpt")
		}
	})

	http.HandleFunc("/davinci", func(w http.ResponseWriter, r *http.Request) {
		err = RequestToDaVinci(w, r)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("richiesta davinci")
		}
	})

	http.HandleFunc("/ada", func(w http.ResponseWriter, r *http.Request) {
		err = RequestToAda(w, r)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("richiesta ada")
		}
	})

	http.HandleFunc("/gpt4", func(w http.ResponseWriter, r *http.Request) {
		err = RequestToGPT4(w, r)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("richiesta gpt4")
		}
	})

	http.HandleFunc("/immagina", func(w http.ResponseWriter, r *http.Request) {
		err = RequestToImageCreator(w, r)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("richiesta immagine")
		}
	})

	go http.ListenAndServe(":"+port, nil)
}
