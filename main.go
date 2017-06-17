package main

import (
	"os"

	"log"

	"net/http"

	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot: ", bot, ", err: ", err)

	http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	log.Println("Port: ", port)
	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			log.Println(message.Text)
		}
	}

}
