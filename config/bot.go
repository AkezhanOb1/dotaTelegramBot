package config

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Bot - represents telegram bot
var Bot *tb.Bot

func init() {
	var err error
	Bot, err = tb.NewBot(tb.Settings{
		Token: "1012701036:AAF6-nFSSBnbWb6FyBuVbPDDa_Y9Vtoxy2Q", //test bot
		//Token:  "923261617:AAGIM0wr4rrUZxFL7xCAn1i62eQK5h7pCUE",
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println("error occured while creating bot", err)
	}

	log.Println(" bot buccessfully launched")
}
