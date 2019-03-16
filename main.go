package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	http.HandleFunc("/linemadi", linemadi)
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}

func linemadi(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		fmt.Println(err.Error())
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}
