package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

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
				if connectDB() {
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Connect to database successfully")).Do()
					if err != nil {
						log.Print(err)
					}
				}
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func connectDB() bool {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Successfully connected!")
	return true
}

// 1st sprint
// [] add player
// [] remove player
// [] list player
// [] check bill
//
